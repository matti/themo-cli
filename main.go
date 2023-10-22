package main

import (
	"flag"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/matti/themo-cli/themo"
)

func main() {
	var username string
	var password string
	flag.StringVar(&username, "username", "", "")
	flag.StringVar(&password, "password", "", "")

	var low string
	var high string
	var lowhigh string
	flag.StringVar(&low, "low", "", "")
	flag.StringVar(&high, "high", "", "")
	flag.StringVar(&lowhigh, "lowhigh", "", "")

	flag.Parse()

	client := resty.New()
	client.EnableTrace()
	client.SetBaseURL("https://app.themo.io")

	t := themo.New(client)

	if _, err := t.Login(username, password); err != nil {
		log.Fatalln(err)
	}
	var me *themo.Me
	if m, err := t.Me(); err != nil {
		log.Fatalln(err)
	} else {
		me = m
	}

	var devices []*themo.Device
	if d, err := t.Devices(me.Id); err != nil {
		log.Fatalln(err)
	} else {
		devices = d
	}

	device := devices[0]

	var colorLow *themo.Color
	var colorHigh *themo.Color

	if lowhigh == "" {
		if c, err := themo.ParseColor(low); err != nil {
			log.Fatalln(err)
		} else {
			colorLow = c
		}
		if c, err := themo.ParseColor(high); err != nil {
			log.Fatalln(err)
		} else {
			colorHigh = c
		}
	} else {
		if c, err := themo.ParseColor(lowhigh); err != nil {
			log.Fatalln(err)
		} else {
			colorLow = c
			colorHigh = c
		}
	}

	if err := device.SetColor(colorLow, colorHigh); err != nil {
		log.Fatalln(err)
	}

	// if c, err := device.Colors(); err != nil {
	// 	log.Fatalln(err)
	// } else {
	// 	log.Println(c[0], c[1])
	// }
}
