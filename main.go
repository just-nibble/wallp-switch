package main

import (
	"bytes"
	"fmt"
	"os/exec"
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
	fmt.Println(SwitchWallpaper())
}
