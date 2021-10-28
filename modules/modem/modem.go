package modem

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/minish144/go-sms-api/gen/pb"
	"github.com/spf13/viper"
	"github.com/tarm/serial"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

type Modem struct {
	comport  string
	boudrate int
	instance *serial.Port
}

// TODO add checking if device is connected for windows
func New(comport string, boudrate int) (*Modem, error) {
	if _, err := os.Stat(comport); runtime.GOOS != "windows" && os.IsNotExist(err) {
		return nil, errors.New("comport was not found, try checking config.yaml or reconnecting your device")
	}

	m := &Modem{comport: comport, boudrate: boudrate}
	c := &serial.Config{Name: comport, Baud: boudrate, ReadTimeout: time.Second}
	var err error
	m.instance, err = serial.OpenPort(c)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Modem) Send(number string, message string) error {
	c := make(chan error)

	go func() {
		m.sendCommand("AT+CMGF=1\r", false)
		m.sendCommand("AT+CMGS=\""+number+"\"\r", false)
		_, err := m.sendCommand(message+string(rune(26)), true) // 26 = CTRL+Z
		if err != nil {
			c <- err
		}
		c <- nil
	}()

	select {
	case err := <-c:
		return err
	case <-time.After(viper.GetDuration("modem.sms_timeout") * time.Second):
		return errors.New("failed to send sms: timeout exceed")
	}
}

func (m *Modem) ReadAll() ([]*pb.Message, error) {
	m.sendCommand("AT+CMGF=1\r", false)
	x, err := m.sendCommand("AT+CMGL=\"ALL\"\r", true)
	if err != nil {
		log.Println(err.Error())
	}
	return parseMessage(x)
}

func parseMessage(text string) ([]*pb.Message, error) {
	var list []*pb.Message
	listLines := strings.Split(text, "\r\n")

	firstIndex := 0
	if len(listLines) > 0 && len(listLines[0]) < 3 {
		firstIndex = 1
	}

	for i := firstIndex; i < len(listLines)-3; i = i + 2 {
		tmp := strings.Split(listLines[i], ",")
		fmt.Println(tmp)
		if len(tmp) < 3 {
			continue
		}
		tmp[2] = strings.Replace(tmp[2], `"`, ``, -1)
		id := tmp[0][7:]
		phone := tmp[2]
		msg := listLines[i+1]

		bs, err := hex.DecodeString(msg)
		if err != nil {
			return nil, err
		}
		e := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
		es, _, err := transform.Bytes(e.NewDecoder(), bs)
		if err != nil {
			return nil, err
		}
		fmt.Println(string(es))
		list = append(list, &pb.Message{Id: id, Phone: string(phone), Date: tmp[4], Message: string(msg)})
	}
	return list, nil
}

/*
func (m *Modem) Delete(id string) {
	m.sendCommand("AT+CMGF=1\r", false)
	x := m.sendCommand("AT+CMGD="+id+"\r", true)
	log.Println("MESSAGE ", x)
}
*/

func (m *Modem) sendCommand(message string, wait bool) (string, error) {
	m.instance.Flush()
	_, err := m.instance.Write([]byte(message))
	if err != nil {
		return "", err
	}
	buf := make([]byte, 1024)
	var loop int = 1
	if wait {
		loop = 10
	}
	var msg string
	var status string
	for i := 0; i < loop; i++ {
		n, _ := m.instance.Read(buf)
		if n > 0 {
			status = string(buf[:n])
			msg += status
			if strings.HasSuffix(status, "OK\r\n") || strings.HasSuffix(status, "ERROR\r\n") {
				break
			}
		}
	}

	return msg, nil
}
