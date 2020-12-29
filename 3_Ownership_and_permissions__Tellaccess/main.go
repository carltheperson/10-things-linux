package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

const (
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
	colorReset = "\033[0m"
)

type filePermissions struct {
	owner string
	group string
	other string
}

func getStat(filename string) *syscall.Stat_t {
	info, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}
	stat := info.Sys().(*syscall.Stat_t)
	return stat
}

func getMode(filename string) os.FileMode {
	info, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}

	mode := info.Mode()
	return mode
}

func getUserName(filename string) string {
	stat := getStat(filename)
	uid := stat.Uid
	u := strconv.FormatUint(uint64(uid), 10)
	usr, _ := user.LookupId(u)
	return strings.ToLower(usr.Name)
}

func getGroupNameFromID(id string) string {
	group, _ := user.LookupGroupId(id)
	return strings.ToLower(group.Name)
}

func getGroupName(filename string) string {
	stat := getStat(filename)
	gid := stat.Gid
	groupID := strconv.FormatUint(uint64(gid), 10)
	return getGroupNameFromID(groupID)
}

func generateHumanReadablePermissions(symbolicPermissions string) string {
	permissions := ""
	if symbolicPermissions[0] == 'r' {
		permissions += "read"
	}
	if symbolicPermissions[1] == 'w' {
		if len(permissions) != 0 {
			permissions += ", write"
		} else {
			permissions += "write"
		}
	}
	if symbolicPermissions[2] == 'x' {
		if len(permissions) != 0 {
			permissions += ", execute"
		} else {
			permissions += "execute"
		}
	}
	return colorRed + permissions + colorReset
}

func checkIfInGroup(groupName string) bool {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	userGroupIds, _ := usr.GroupIds()
	for _, id := range userGroupIds {
		if getGroupNameFromID(id) == groupName {
			return true
		}
	}
	return false
}

func getFilePermissions(filename string) filePermissions {
	mode := getMode(filename)

	permissions := filePermissions{owner: "", group: "", other: ""}

	for i := 1; i < 4; i++ {
		permissions.owner += string(mode.String()[i])
	}

	for i := 4; i < 7; i++ {
		permissions.group += string(mode.String()[i])
	}

	for i := 7; i < 10; i++ {
		permissions.other += string(mode.String()[i])
	}

	return permissions
}

func printFilePermissions(filename string) {

	permissions := getFilePermissions(filename)

	fmt.Printf("Owner (%s): %s", getUserName(filename), generateHumanReadablePermissions(permissions.owner))

	fmt.Printf("\nGroup (%s): %s", getGroupName(filename), generateHumanReadablePermissions(permissions.group))

	fmt.Printf("\nOther: %s\n", generateHumanReadablePermissions(permissions.other))
}

func printUserPermissions(filename string) {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	userName := strings.ToLower(usr.Name)
	permissions := getFilePermissions(filename)

	if userName == getUserName(filename) {
		fmt.Printf("You as user %s can: %s\n", userName, generateHumanReadablePermissions(permissions.owner))
	} else if checkIfInGroup(getGroupName(filename)) {
		fmt.Printf("You being apart of the group %s can: %s\n", getGroupName(filename), generateHumanReadablePermissions(permissions.group))
	} else {
		fmt.Printf("You not being the owner or apart of the group can: %s\n", generateHumanReadablePermissions(permissions.other))
	}

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You need to supply ONE argument: The filename.")
		return
	}
	filename := os.Args[1]

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Printf("Not a file or directory: %s\n", filename)
		return
	}
	if err != nil {
		panic(err)
	}

	fmt.Printf("Permissions for %s:\n\n", filename)
	printFilePermissions(filename)
	fmt.Println()
	printUserPermissions(filename)
}
