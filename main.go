package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetNextVersion(currentVersion string, commitMessage string) string {
	// splitting version to variables
	semVerSplit := strings.Split(currentVersion, ".")
	major, err := strconv.Atoi(semVerSplit[0])
	check(err)
	minor, err := strconv.Atoi(semVerSplit[1])
	check(err)
	patch, err := strconv.Atoi(semVerSplit[2])
	check(err)

	if strings.Contains(commitMessage, "semver") {
		if strings.Contains(commitMessage, "patch") {
			return fmt.Sprintf("%d.%d.%d", major, minor, patch+1)
		}
		if strings.Contains(commitMessage, "minor") {
			return fmt.Sprintf("%d.%d.%d", major, minor+1, 0)
		}
		if strings.Contains(commitMessage, "major") {
			return fmt.Sprintf("%d.%d.%d", major+1, 0, 0)
		}
	}

	return "0.0.0"
}

func main() {
	// parsing command flags
	commitMessage := flag.String("commit", "", "Message for the commit")
	currentVersion := flag.String("version", "", "Version to update to")
	flag.Parse()

	newSemVar := GetNextVersion(*currentVersion, *commitMessage)

	d1 := []byte(newSemVar)
	err := os.WriteFile("./version.txt", d1, 0644)
	check(err)
}
