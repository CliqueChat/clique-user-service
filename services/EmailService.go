package services

import (
	"github.com/CliqueChat/clique-user-service/helpers"
	"github.com/CliqueChat/clique-user-service/resources"
	"log"
	"net/smtp"
)

var prop = resources.GetApplicationProfile()

type stmpServer struct {
	host     string // host of the email service
	port     string // port of the email service
	fromMail string // email used by the server to send emails
	auth     smtp.Auth
}

func (s *stmpServer) Address() string {
	return s.host + ":" + s.port
}

var emailServer stmpServer

func TestSendMail(to []string, message string) {

	// Message.
	bytes := []byte(message)

	err := smtp.SendMail(emailServer.Address(), emailServer.auth, emailServer.fromMail, to, bytes)

	if err != nil {
		panic(err)
	}
}

func ConnectEmailServer() {

	emailSimulation := prop.GetBool(helpers.SimulateEmailProcess, true)

	if !emailSimulation {
		validateEmailServerInfoAndPopulateEmailServer()
		log.Println("CONNECTING TO EMAI SERVER. ADDRESS: " + emailServer.host + ":" + emailServer.port)
	} else {
		log.Println("SIMULATED, EMAIL SERVER CONNECTION IGNORED")
	}

}

func validateEmailServerInfoAndPopulateEmailServer() {

	if host, valid := prop.Get(helpers.EmailServerHost); valid {
		emailServer.host = host
	} else {
		panic("email host not provided in property file")
	}

	if port, valid := prop.Get(helpers.EmailServerPort); valid {
		emailServer.port = port
	} else {
		panic("email port not provided in property file")
	}

	if mail, valid := prop.Get(helpers.EmailServerSenderMail); valid {
		emailServer.fromMail = mail
	} else {
		panic("mail not provided in property file to start email server")
	}

	if pwd, valid := prop.Get(helpers.EmailServerSenderMailPassword); valid {
		emailServer.auth = smtp.PlainAuth("", emailServer.fromMail, pwd, emailServer.host)
	} else {
		panic("password not provided in property file to start email server")
	}
}
