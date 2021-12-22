package core

import (
	"fmt"
	"os/exec"
)

/*
 * Execute "go get -u <package_path>" command
 * @params (package path)
 * @return (void)
 */
func ExecCommand(packagePath string) {
	fmt.Println("\nInstalling " + packagePath + "...")

	cmd := exec.Command("go", "get", "-u", packagePath)
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	fmt.Println("\nSuccessfully ^.^")
}
