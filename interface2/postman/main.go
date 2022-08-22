package main

import "GOLANG/interface2/fedex"

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

type Sender interface {
	Send(parcel string)
}

func main() {
	var sender Sender = &fedex.FedexSender{}
	// var sender Sender = &koreapost.PostSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 고르바", sender)
}
