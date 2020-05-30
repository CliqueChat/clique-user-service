package services

import (
	"bytes"
	"github.com/CliqueChat/clique-user-service/helpers"
	"github.com/CliqueChat/clique-user-service/resources"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"path/filepath"
)

var prop, _ = resources.GetApplicationProfile()

type stmpServer struct {
	host     string // host of the email service
	port     string // port of the email service
	fromMail string // email used by the server to send emails
	auth     smtp.Auth
}

func (s *stmpServer) Address() string {
	return s.host + ":" + s.port
}

// Stores the instance of the stmp server
var emailServer stmpServer

// Stores the templates
var userVerificationTemplate *template.Template
var userPasswordResetTemplate *template.Template

func ConnectEmailServer() {

	emailSimulation := prop.GetBool(helpers.SimulateEmailProcess, true)

	if !emailSimulation {
		validateEmailServerInfoAndPopulateEmailServer()
		log.Println("CONNECTING TO EMAI SERVER. ADDRESS: " + emailServer.host + ":" + emailServer.port)
	} else {
		log.Println("SIMULATED, EMAIL SERVER CONNECTION IGNORED")
	}

	wd, _ := os.Getwd()

	// Load Email templates
	path := filepath.Join([]string{wd, "resources", "assets", "user_verification_email_template.html"}...)

	userVerificationTemplate, _ = template.ParseFiles(path)

	userPasswordResetTemplate, _ = template.New("user_password_reset").
		ParseFiles(filepath.
			Join([]string{wd, "resources", "assets", "user_password_reset_email_templete.html"}...))
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

func SendEmail(to []string, subject string, body string) {

	msg := "Subject:" + subject + "\n" + "\n" + body

	// Message in bytes
	message := []byte(msg)

	err := smtp.SendMail(emailServer.Address(), emailServer.auth, emailServer.fromMail, to, message)

	if err != nil {
		panic(err)
	}

}

func SendEmailWithHTMLContent(to []string, subject string, body string) {

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject = "Subject:" + subject + "\n"

	message := []byte(subject + mime + body)

	err := smtp.SendMail(emailServer.Address(), emailServer.auth, emailServer.fromMail, to, message)

	if err != nil {
		panic(err)
	}

}

type UserVerificationEmailContent struct {
	Name       string
	OtpCode    string
	ExpiryTime string // in minutes
	AppName    string
}

func SendUserVerificationEmail(to []string, content UserVerificationEmailContent) {

	content.AppName = "Clique Chat"
	var tpl bytes.Buffer
	err := userVerificationTemplate.Execute(&tpl, content)

	if err != nil || (UserVerificationEmailContent{}) == content {
		panic("Error")
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: [VERIFY] - Please Verify your email to continue with Clique Chat" + "\n"

	message := []byte(subject + mime + tpl.String())

	err = smtp.SendMail(emailServer.Address(), emailServer.auth, emailServer.fromMail, to, message)

	if err != nil {
		panic(err)
	}
	// TODO - return statements
}

type UserPassWordResetContent struct {
	AppName string
}

func SendUserPasswordResetEmail(to []string, content UserPassWordResetContent) {

	content.AppName = "Clique Chat"
	var tpl bytes.Buffer

	err := userPasswordResetTemplate.Execute(&tpl, content)

	if err != nil || (UserPassWordResetContent{}) == content {
		panic("Error")
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: [VERIFY] - Please Reset your Clique Chat account password" + "\n"

	message := []byte(subject + mime + tpl.String())

	err = smtp.SendMail(emailServer.Address(), emailServer.auth, emailServer.fromMail, to, message)

	if err != nil {
		panic(err)
	}
	// TODO - return statement
}
