package user

import (
	"net/http"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var TWILIO_ACCOUNT_SID string = "AC0db6ef3790d9106217135913857f82bc"
var TWILIO_AUTH_TOKEN string = "5d5724cb76150062eab6e1f26524c619"
var VERIFY_SERVICE_SID string = "VAd23ee4381a65d7929a175a199de806f4"
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
