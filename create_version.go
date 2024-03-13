package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const template string = `
package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command(%s, os.Args[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
`

func main() {
	fmt.Printf("Print new version (For example: 3.8): ")

	var newVersion string
	_, err := fmt.Scanf("%s", &newVersion)
	if err != nil {
		log.Fatal(err)
	}

	versionDirname, err := os.MkdirTemp("D:/Pythons/", ".*")
	if err != nil {
		log.Fatal(err)
	}

	versionFile, err := os.CreateTemp(versionDirname, fmt.Sprintf("python%s_*.go", newVersion))
	if err != nil {
		log.Fatal(err)
	}
	pipFile, err := os.CreateTemp(versionDirname, fmt.Sprintf("pip%s_*.go", newVersion))
	if err != nil {
		log.Fatal(err)
	}

	_, err = versionFile.WriteString(fmt.Sprintf(template, fmt.Sprintf("\"D:/Pythons/Python-%s/python.exe\"", newVersion)))
	if err != nil {
		log.Fatal(err)
	}
	_, err = pipFile.WriteString(fmt.Sprintf(template, fmt.Sprintf("\"D:/Pythons/Python-%s/Scripts/pip.exe\"", newVersion)))
	if err != nil {
		log.Fatal(err)
	}

	err = versionFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = pipFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	os.Chdir(fmt.Sprintf("D:/Pythons/%s", versionDirname))

	cmd := exec.Command(
		"go", "build", "-o",
		fmt.Sprintf("D:/Pythons/python%s.exe", newVersion),
		versionFile.Name(),
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command(
		"go", "build", "-o",
		fmt.Sprintf("D:/Pythons/pip%s.exe", newVersion),
		pipFile.Name(),
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	os.Chdir("..")

	err = os.RemoveAll(versionDirname)
	if err != nil {
		log.Fatal(err)
	}
}
