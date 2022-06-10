package fedex

import (
	"fmt"
)

type FedexSender struct {
}

func (fs *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex sends %v parcel\n", parcel)
}
