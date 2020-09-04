package providers

import (
	"fmt"
	"github.com/voluntariado-ucc-ing/volunteer_api/config"
	"github.com/voluntariado-ucc-ing/volunteer_api/domain/apierrors"
	"net/smtp"
)

var auth smtp.Auth

func init() {
	mail, pass := config.GetMailCredentials()
	auth = smtp.PlainAuth("", mail, pass, config.SmtpHost)
}

func SendMail(emailAddress string, password string) {
	err := smtp.SendMail(
		config.SmtpAddress,
		auth,
		"voluntariadoing.noreply@ucc.edu.ar",
		[]string{emailAddress},
		[]byte(fmt.Sprintf("Hola, fuiste aceptado en el voluntariado de UCC Ingenieria, tu clave de acceso es %s", password)),
	)
	if err != nil {
		e := apierrors.NewInternalServerApiError(fmt.Sprintf("Error while trying to mail %s and password %s", emailAddress, password), err)
		fmt.Println(e)
		return
	}

	fmt.Println("Successfully posted mail to user ", emailAddress)
}