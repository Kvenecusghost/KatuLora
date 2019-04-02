package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server\n============================")
	r, err := http.Get("https://ground-sensors-app.data.thethingsnetwork.org/&key=ttn-account-v2.Bf5PmOZpgKj5vzAHdRMgO_PRdRsl-9KPKfP0baapZiU")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Response Header", r.Header)
	fmt.Println("Response Bpdu", r.Body)
	//json.Unmarshal(r)
}
