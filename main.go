package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"time"
)

var wallpapers []string = []string{}

var usr, _ = user.Current()
var dir string = usr.HomeDir

var folder string = dir + "/Pictures/wallpaper/"

func switchWallpaper(wallpaper int64) (msg string, err error) {
	var wallpaper_path string = folder + wallpapers[wallpaper]
	var command string
	var args []string = []string{}

	var de string = os.Getenv("XDG_SESSION_DESKTOP")
	de = strings.ToLower(de)

	switch de {
	case "kde":
		command = "plasma-apply-wallpaperimage"
		args = append(args, wallpaper_path)
	case "gnome":
		command = "gsettings"
		args = append(args, "set", "org.gnome.desktop.background", "picture-uri", wallpaper_path)
	case "mate":
		command = "gsettings"
		args = append(args, "set", "org.mate.desktop.background", "picture-uri", wallpaper_path)
	case "cinnamon":
		command = "gsettings"
		args = append(args, "set", "org.cinnamon.desktop.background", "picture-uri", wallpaper_path)
	case "xfce":
		command = "xfconf-query"
		args = append(args, "-c", "xfce4-desktop", "-p", "/backdrop/screen0/monitor1/workspace0/last-image", "-s", wallpaper_path)
	}
	cmd := exec.Command(command, args...)

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
	fmt.Println(switchWallpaper(current_wallpaper))
}
