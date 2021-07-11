package main

import (
	"fmt"
	"time"

	"github.com/shortcuts/keys"
)

const (
	delayKeyfetchMS = 100
)

func main() {

	for {
		key := keys.GetKey()
		
		if !key.Empty {
			if  key.Modifiers.CTRL && key.Modifiers.ALT{
				fmt.Print("CTRL + ALT ")
				for _, k := range key.Keys {
					fmt.Print(string(k.Rune), " ")
				}
				fmt.Print("\n")
			}
		}

		time.Sleep(delayKeyfetchMS * time.Millisecond)
	}
}