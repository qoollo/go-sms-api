package main

import (
	"fmt"

	"github.com/minish144/go-sms-api/modules/modem"
	"github.com/minish144/go-sms-api/modules/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// viper init
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatalln()
	}
	logrus.Info("viper config initialized successfully")

	comport := viper.GetString("modem.comport")
	baudrate := viper.GetInt("modem.baudrate")
	newModem, err := modem.New(comport, baudrate)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{
				"error":    err.Error(),
				"comport":  comport,
				"baudrate": baudrate,
			},
		).Errorln("failed to initialize a modem")
	}

	if resp, err := newModem.USSD("AA18AC3602"); err != nil {
		fmt.Println("resp:", resp, "err:", err.Error())
	}

	// server init
	server.Run()
}
