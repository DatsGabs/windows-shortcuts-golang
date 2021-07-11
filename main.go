package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/shortcuts/keys"
)

const (
	delayKeyfetchMS = 100
)


func main() {
	// modify this to write commands to run on cmd
	// order ctrl > shift > alt
	commands := map[string]string {
		"ctrl+alt+t": "wt",
		"ctrl+alt+s": "wsl --shutdown",
	}

	
	for {
		key := keys.GetKey()
		
		if !key.Empty {
			mapString := ""

			if key.Modifiers.CTRL {
				mapString += "ctrl+"
			}
			if key.Modifiers.SHIFT {
				mapString += "shift+"
			}
			if key.Modifiers.ALT {
				mapString += "alt+"
			}

			for index, k := range key.Keys {
				if index < len(key.Keys) - 1 {
					mapString += ""
				}
				mapString += string(k.Rune)
			}
			
			command, exists := commands[mapString]
			if exists {
				cmd := exec.Command("cmd", "/c", command)
				err := cmd.Start()
				if err != nil {
					fmt.Println(err)
				}
			}
		}

		time.Sleep(delayKeyfetchMS * time.Millisecond)
	}

}
