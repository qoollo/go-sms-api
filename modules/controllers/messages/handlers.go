package messages

import (
	"context"

	"github.com/argandas/sim900"
	"github.com/minish144/go-sms-api/gen/pb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Send(ctx context.Context, in *pb.Messages_SendRequest) (*pb.Messages_SendResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	comport := viper.GetString("modem.comport")
	baudrate := viper.GetInt("modem.baudrate")
	newModem := sim900.New()
	if err := newModem.Connect(comport, baudrate); err != nil {
		logrus.WithFields(
			logrus.Fields{
				"error":    err.Error(),
				"comport":  comport,
				"baudrate": baudrate,
			},
		).Errorln("failed to initialize a modem")
		return nil, err
	}

	if err := newModem.SendSMS(in.Message.Phone, in.Message.Message); err != nil {
		logrus.WithFields(
			logrus.Fields{
				"error": err.Error(),
			},
		).Errorln("failed to send sms")
		return nil, err
	}

	logrus.WithFields(
		logrus.Fields{
			"phone":   in.Message.Phone,
			"message": in.Message.Message,
		},
	).Infoln("sms was sent successfully")

	return &pb.Messages_SendResponse{}, nil
}

// func List(ctx context.Context, in *pb.Messages_ListRequest) (*pb.Messages_ListResponse, error) {
// 	comport := viper.GetString("modem.comport")
// 	baudrate := viper.GetInt("modem.baudrate")
// 	newModem, err := modem.New(comport, baudrate)
// 	if err != nil {
// 		logrus.WithFields(
// 			logrus.Fields{
// 				"error":    err.Error(),
// 				"comport":  comport,
// 				"baudrate": baudrate,
// 			},
// 		).Errorln("failed to initialize a modem")
// 		return nil, err
// 	}

// 	messages, err := newModem.ReadAll()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.Messages_ListResponse{Messages: messages}, nil
// }
