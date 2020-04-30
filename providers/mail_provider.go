package providers

import (
	"fmt"
	"net/smtp"
	"volutarios_api/config"
	"volutarios_api/domain/apierrors"
)

var auth smtp.Auth

func init() {
	mail, pass := config.GetMailCredentials()
	auth = smtp.PlainAuth("", mail, pass, config.SmtpHost)
}

func SendMail(emailAddress string, password string) apierrors.ApiError {
	err := smtp.SendMail(
		config.SmtpAddress,
		auth,
		"noreply@voluntariadoing.org",
		[]string{emailAddress},
		[]byte(fmt.Sprintf("Hola, fuiste aceptado en el voluntariado de UCC Ingenieria, tu clave de acceso es %s", password)),
	)
	if err != nil {
		return apierrors.NewInternalServerApiError(fmt.Sprintf("Error while trying to mail %s", emailAddress), err)
	}
	return nil
}
