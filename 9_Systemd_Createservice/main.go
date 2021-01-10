package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getServiceName(filename string) string {
	if strings.Contains(filename, "/") {
		dirs := strings.Split(filename, "/")
		return dirs[len(dirs)-1] + ".service"
	}
	return filename + ".service"
}

func createServiceFile(filename string, extraArgs []string) {
	extraArgsJoined := strings.Join(extraArgs, " ")
	// This is mostly meant for web services
	file := fmt.Sprintf(`
[Unit]
Description = Service automatically created using createservice.
After = network-online.target

[Service]
ExecStart = %s %s

[Install]
WantedBy = multi-user.target`, filename, extraArgsJoined)

	f, err := os.Create(filename + ".service")
	defer f.Close()

	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(file)
	if err != nil {
		panic(err)
	}
}

func enableService(filename string) {
	out, err := exec.Command("systemctl", "enable", filename+".service").Output()
	if err != nil {
		fmt.Println(out)
		panic(err)
	}
}

func startService(serviceName string) {
	out, err := exec.Command("systemctl", "start", serviceName).Output()
	if err != nil {
		fmt.Println(out)
		panic(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You need to supply at least argument: The executable.")
		return
	}
	filename := os.Args[1]
	extraArgs := os.Args[2:]

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Printf("Cannot find: %s\n", filename)
		return
	}
	if err != nil {
		panic(err)
	}

	fmt.Println("Creating service file")
	createServiceFile(filename, extraArgs)

	fmt.Println("Enableing service")
	enableService(filename)

	fmt.Println("Starting service")
	startService(getServiceName(filename))

	fmt.Println("Created " + getServiceName(filename))
}
