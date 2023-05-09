package main

import (
	"bytes"
	"log"
	"os/exec"

	"github.com/mattn/go-tty"
)

func SwitchWallpaper() (msg string, err error) {
	cmd := exec.Command("plasma-apply-wallpaperimage", "/home/princewillo/Pictures/wallpaper/samurai_strike.jpg")

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		return
	}
	return out.String(), nil
}

func main() {
	var key_combo []string = make([]string, 5)
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		key_combo = append(key_combo, string(r))
		if len(key_combo) >= 4 {
			if key_combo[0] == "[" && key_combo[1] == "6" && key_combo[2] == "~" && key_combo[3] == "a" {
				SwitchWallpaper()
			}
			key_combo = nil
		}
	}
}
