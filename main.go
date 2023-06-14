package main

import (
	"fmt"
	"github.com/copo888/copo_otp/helper/otpx"
	"log"
)

func main() {

	account := "yangoo"

	r, err := otpx.GenOtpKey("copo", account)
	if err != nil {
		panic(err)
	}

	fmt.Println(r)

	var passcode string
	fmt.Println("input one time password:")
	fmt.Scanf("%s", &passcode)
	valid := otpx.Validate(passcode, r.Code)

	log.Println(valid)
}
