package main

import (
	"log"
	"proxypay/lib"
	"time"
)

var times = time.Tick(1 * time.Minute)

func processPayment() {

	for range times {
		pays, err := lib.GetPayments(20)

		if err != nil {
			log.Panic(err)
		}

		for _, pay := range *pays {
			go lib.ConfirmPayment(pay.ID)
		}
	}
}
