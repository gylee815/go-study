package koreaPost

import (
	"fmt"
)

type KoreaPostSender struct {
}

func (p *KoreaPostSender) Send(parcel string) {
	fmt.Printf("KoreaPost send %s parcel\n", parcel)
}
