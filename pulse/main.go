package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"

	"github.com/evantbyrne/pulse"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app        = kingpin.New("pulse", "Check the status of a web page.")
	url        = app.Arg("url", "URL to check.").Required().String()
	configFile = app.Flag("config", "SMTP configuration file.").Short('c').String()
)

type smtpConfig struct {
	From         string
	SmtpHost     string
	SmtpPassword string
	SmtpPort     int
	SmtpUsername string
	To           string
}

func main() {
	var config smtpConfig

	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *configFile != "" {
		data, configErr := ioutil.ReadFile(*configFile)
		if configErr != nil {
			fmt.Printf("Could not read file '%s'\n", *configFile)
			os.Exit(1)
		}
		if configErr = json.Unmarshal(data, &config); configErr != nil {
			fmt.Println(configErr)
			os.Exit(1)
		}
	}

	if err := pulse.Check(*url); err != nil {
		fmt.Println(err)
		if *configFile != "" {
			auth := smtp.PlainAuth(
				"",
				config.SmtpUsername,
				config.SmtpPassword,
				config.SmtpHost,
			)
			smtpErr := smtp.SendMail(
				fmt.Sprintf("%s:%d", config.SmtpHost, config.SmtpPort),
				auth,
				config.From,
				[]string{config.To},
				[]byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: Pulse Failure @ %s\r\n\r\nError: %s\r\n", config.From, config.To, *url, err)),
			)
			if smtpErr != nil {
				fmt.Println(smtpErr)
			}
		}
		os.Exit(1)
	}
}
