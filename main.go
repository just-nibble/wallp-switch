package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"os/user"
	"time"
)

var wallpapers []string = []string{}

var usr, _ = user.Current()
var dir string = usr.HomeDir

var folder string = dir + "/Pictures/wallpaper/"

func SwitchWallpaper(wallpaper int64) (msg string, err error) {
	var wallpaper_path string = folder + wallpapers[wallpaper]
	cmd := exec.Command("plasma-apply-wallpaperimage", wallpaper_path)

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		return
	}
	return out.String(), nil
}

func get_wallpapers() {
	file, err := os.Open(folder)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	names, err := file.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}

	wallpapers = names
}

func main() {
	rand.Seed(time.Now().UnixNano())
	get_wallpapers()
	var current_wallpaper int64 = rand.Int63n(int64(len(wallpapers)) - 1)
	fmt.Println(SwitchWallpaper(current_wallpaper))
}
