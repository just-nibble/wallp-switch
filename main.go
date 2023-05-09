package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

var wallpapers []string = []string{}

var usr, _ = user.Current()
var dir string = usr.HomeDir

var folder string = dir + "/Pictures/wallpaper/"
var config_file string = dir + "/current_wallpaper.txt"

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

func set_wallpaper(data int64) {
	f, err := os.Create(config_file)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(strconv.FormatInt(data, 10))

	if err2 != nil {
		log.Fatal(err2)
	}

}

func get_current_wallpaper() int64 {
	content, err := os.ReadFile(config_file)
	var current string = strings.TrimSuffix(string(content), "\n")
	if err != nil {
		log.Fatal(err)
	}
	current_wallpaper, err := strconv.ParseInt(current, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return current_wallpaper
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
	get_wallpapers()
	var current_wallpaper int64 = get_current_wallpaper()
	fmt.Println(SwitchWallpaper(current_wallpaper))
	if current_wallpaper <= int64(len(wallpapers)-1) {
		current_wallpaper++
	} else {
		current_wallpaper = 0
	}
	set_wallpaper(current_wallpaper)
}
