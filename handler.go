package webhook

import (
	"encoding/json"
	"fmt"
	"github.com/aiteung/atapi"
	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
	"github.com/whatsauth/wa"
	"net/http"
	"os"
)

func PostBalasan(w http.ResponseWriter, r *http.Request) {
	var msg model.IteungMessage
	var resp atmessage.Response
	json.NewDecoder(r.Body).Decode(&msg)
	if r.Header.Get("Secret") == os.Getenv("SECRET") {
		dt := &wa.TextMessage{
			To:       msg.Phone_number,
			IsGroup:  false,
			Messages: "Hai Hai Haiii kamuuuui " + msg.Alias_name + "rofinya lagi gaadaa \n aku giseuubott salam kenall yaaaa \n",
		}
		resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/send/message/text")
	} else {
		resp.Response = "Secret Salah"
	}
	fmt.Fprintf(w, resp.Response)
}
