package main

import (
	"flag"
	"fmt"
)

func main() {

	versionFlag := flag.Bool("version", false, "Version")
	flag.Parse()

	if *versionFlag {
		fmt.Println("Git Commit:", GitCommit)
		fmt.Println("Version:", Version)
		if VersionPrerelease != "" {
			fmt.Println("Version PreRelease:", VersionPrerelease)
		}
		return
	}

	s, m, err := render("Hello {world} at {position}!", "Earth", struct {
		Lat  float32
		Long float32
	}{
		Lat:  24.7,
		Long: 132.2,
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(m)
}