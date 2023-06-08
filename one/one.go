package one

import (
	"fmt"
	"github.com/rapita/gomod/header"
	"github.com/rapita/gomod/one/internal"
)

func One() {
	header.Header()
	internal.OneInternal()
	fmt.Println("one")
}
