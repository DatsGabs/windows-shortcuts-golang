package keys

import (
	"syscall"
	"unicode/utf8"
	"unsafe"

	"github.com/gonutz/w32"
)


type Keys struct {
	Empty   bool
	Keys []Key
	Modifiers Modifiers
}

type Key struct {
	Rune rune
	Keycode int
}

type Modifiers struct {
	CTRL bool
	ALT bool
	SHIFT bool
}

var (
	moduser32 = syscall.NewLazyDLL("user32.dll")

	procGetKeyboardLayout     = moduser32.NewProc("GetKeyboardLayout")
	procToUnicodeEx           = moduser32.NewProc("ToUnicodeEx")
)

// getKey gets the current entered key by the user, if there is any
func GetKey() Keys {
	var activeKeys []int
	var keyState uint16

	for i := 0; i < 256; i++ {

		keyState = w32.GetAsyncKeyState(i)

		isPressed := keyState&(1<<15) != 0
		isChar := !(i < 0x2F) && (i < 160 || i > 165) && (i < 91 || i > 93)

		if isPressed && isChar {
			activeKeys = append(activeKeys, i)
		}
	}

	if len(activeKeys) > 0 {
		keys := Keys{Empty: false}

		if w32.GetAsyncKeyState(w32.VK_SHIFT)&(1<<15) != 0 {
			keys.Modifiers.SHIFT = true
		}

		if w32.GetAsyncKeyState(w32.VK_CONTROL)&(1<<15) != 0 {
			keys.Modifiers.CTRL = true
		}

		if w32.GetAsyncKeyState(w32.VK_MENU)&(1<<15) != 0 {
			keys.Modifiers.ALT = true
		}

		for _, keyCode := range activeKeys {
			keys.Keys = append(keys.Keys, ParseKeycode(keyCode))
		}
		return keys 
	}

	return Keys{Empty: true}
}



// ParseKeycode returns the correct Key struct for a key taking in account the current keyboard settings
// That struct contains the Rune for the key

func ParseKeycode(keyCode int) Key {
	key := Key{Keycode: keyCode}

	// Only one rune has to fit in
	outBuf := make([]uint16, 1)

	// Get keyboard layout for this process (0)
	kbLayout, _, _ := procGetKeyboardLayout.Call(uintptr(0))

	// empty keyboard state
	keyboardState := make([]uint8, 256)

	_, _, _ = procToUnicodeEx.Call(
		uintptr(keyCode),
		uintptr(0),
		uintptr(unsafe.Pointer(&keyboardState[0])),
		uintptr(unsafe.Pointer(&outBuf[0])),
		uintptr(1),
		uintptr(1),
		uintptr(kbLayout))
	
	key.Rune, _ = utf8.DecodeRuneInString(syscall.UTF16ToString(outBuf))

	return key
}