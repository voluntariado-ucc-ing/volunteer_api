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
	msg := []byte(fmt.Sprintf("To: %s\r\n" +
		"Subject: ¡Fuiste aceptado en Voluntariado Ingeniería UCC!\r\n" +
		"\r\n" +
		"¡Enhorabuena voluntario!\n\n\nTu acceso al sistema del Voluntariado UCC fue aprobado, tu clave es %s.\nApenas ingreses, podrás cambiar tu contraseña por la que desees.\nMuchas gracias.\r\n", emailAddress, password))
	err := smtp.SendMail(
		config.SmtpAddress,
		auth,
		"voluntariadoing.noreply@ucc.edu.ar",
		[]string{emailAddress},
		msg,
	)

	if err != nil {
		e := apierrors.NewInternalServerApiError(fmt.Sprintf("Error while trying to mail %s and password %s", emailAddress, password), err)
		fmt.Println(e)
		return
	}

	fmt.Println("Successfully posted mail to user ", emailAddress)
}