package main

import (
	fedex "goprojects/ex20.2/fedex"
	kp "goprojects/ex20.2/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func FdxSendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func KpSendBook(name string, sender *kp.KoreaPostSender) {
	sender.Send(name)
}

func main() {
	sender := &fedex.FedexSender{} // sender라는 구조체 선언

	// Sender interface에 sender FedexSender{} struct를 대입. 즉, FedexSender{} struct가 Sender interface를 구현한다.
	SendBook("little prince", sender)

	sender2 := &kp.KoreaPostSender{}
	SendBook("One piece", sender2)
}
