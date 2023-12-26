package user

import (
	"encoding/json"
	"net/http"

	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

type Responseotp struct {
	Virify  bool `json:"virify"`
	Account bool `json:"account"`
	AccSid  string  `json:"accsid"`
}

func Checkotp(w http.ResponseWriter, r *http.Request) {
	var responseotp Responseotp
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("there is proplem for parsing the proplem"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	code := r.Form.Get("code")
	phonenumber := r.Form.Get("phonenumber")

	if code == "" || phonenumber == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("there is proplem for getting the code or the phonenumber"))
		return
	}
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(phonenumber)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(VERIFY_SERVICE_SID, params)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if *resp.Status == "approved" {
		responseotp.Virify = true
		var err error
		responseotp.AccSid, err= searchacc(phonenumber)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ther is eror in the datebase side"))
			return
		}
		if responseotp.AccSid =="" {
			responseotp.Account = false
		}else {
			responseotp.Account = false
		}
	} else {
		responseotp.Virify = false
		
	}
	result, err := json.Marshal(responseotp)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
