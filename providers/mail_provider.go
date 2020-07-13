package providers

import (
	"errors"
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/voluntariado-ucc-ing/volunteer_api/domain/apierrors"
	"github.com/voluntariado-ucc-ing/volunteer_api/domain/auth"
	"time"
)


var rb = rest.RequestBuilder{
	Timeout:             500 * time.Millisecond,
	BaseURL:             "172.17.0.3:3000",
	ContentType:         rest.JSON,
}

func PostMail(r auth.MailRequest) apierrors.ApiError {
	response := rb.Post("/send/volunteers", r)
	if response == nil || response.Response == nil {
		return apierrors.NewInternalServerApiError("Error restclient posting mail to mail api", errors.New("error restclient"))
	}
	if response.StatusCode > 399 {
		error := errors.New(response.String())
		return apierrors.NewInternalServerApiError("Error posting mail", error)
	}
	fmt.Println("successfully sent mail for users")
	return nil
}
