package hotkeys

import (
	"bytes"
	"fmt"
	"syscall"
	"unsafe"
)

const (
	ModAlt = 1 << iota
	ModCtrl
	ModShift
	ModWin
)

type Hotkey struct {
	Id        int // Unique id
	Modifiers int // Mask of modifiers
	KeyCode   int // Key code, e.g. 'A'
}

type MSG struct {
	HWND   uintptr
	UINT   uintptr
	WPARAM int16
	LPARAM int64
	DWORD  int32
	POINT  struct{ X, Y int64 }
}

// String returns a human-friendly display name of the hotkey
// such as "Hotkey[Id: 1, Alt+Ctrl+O]"
func (h *Hotkey) String() string {
	mod := &bytes.Buffer{}
	if h.Modifiers&ModAlt != 0 {
		mod.WriteString("Alt+")
	}
	if h.Modifiers&ModCtrl != 0 {
		mod.WriteString("Ctrl+")
	}
	if h.Modifiers&ModShift != 0 {
		mod.WriteString("Shift+")
	}
	if h.Modifiers&ModWin != 0 {
		mod.WriteString("Win+")
	}
	return fmt.Sprintf("Hotkey[Id: %d, %s%c]", h.Id, mod, h.KeyCode)
}

func registerHotkeys(user32 *syscall.DLL, keys map[int16]*Hotkey) {

	reghotkey := user32.MustFindProc("RegisterHotKey")
	// Hotkeys to listen to:

	// Register hotkeys:
	for _, v := range keys {
		r1, _, err := reghotkey.Call(
			0, uintptr(v.Id), uintptr(v.Modifiers), uintptr(v.KeyCode))
		if r1 == 1 {
			fmt.Println("Registered", v)
		} else {
			fmt.Println("Failed to register", v, ", error:", err)
		}
	}

}

func listenToHotkeys(user32 *syscall.DLL, keys map[int16]*Hotkey) {

	getmsg := user32.MustFindProc("GetMessageW")

	for {
		var msg = &MSG{}
		getmsg.Call(uintptr(unsafe.Pointer(msg)), 0, 0, 0, 1)
		// Registered id is in the WPARAM field:
		fmt.Printf("%v", msg.WPARAM)
		if id := msg.WPARAM; id != 0 {
			fmt.Println("Hotkey pressed:", keys[id])
			if id == 3 { // CTRL+ALT+X = Exit
				fmt.Println("CTRL+ALT+X pressed, goodbye...")
				return
			}
		}
	}
}

func Init(keys map[int16]*Hotkey) {
	go func() {
		user32 := syscall.MustLoadDLL("user32")
		defer user32.Release()
		registerHotkeys(user32, keys)
		fmt.Println("waiting for hotkeys")
		listenToHotkeys(user32, keys)
	}()
}
