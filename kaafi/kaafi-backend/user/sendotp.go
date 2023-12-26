package user

import (
	"net/http"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var TWILIO_ACCOUNT_SID string = "<enter your twilio account>"
var TWILIO_AUTH_TOKEN string = "<your twilio aouthtoken>"
var VERIFY_SERVICE_SID string = "<your serviece sid>"
var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: TWILIO_ACCOUNT_SID,
	Password: TWILIO_AUTH_TOKEN,
})

func Sendotp(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("there is proplem for parsing the proplem"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	channel := r.Form.Get("channel")
	phonenumber := r.Form.Get("phonenumber")

	if channel == "" || phonenumber == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("there is proplem for getting the channel or the phonenumber"))
		return
	}
	params := &openapi.CreateVerificationParams{}
	params.SetTo(phonenumber)
	params.SetChannel(channel)

	resp, err := client.VerifyV2.CreateVerification(VERIFY_SERVICE_SID, params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("there is proplem from pack end"))
		w.Write([]byte(*resp.Status))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("we send the virifacation"))

}
