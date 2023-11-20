package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	uid := os.Getuid()

	usr, _ := user.LookupId(fmt.Sprint(uid))
	fmt.Println(usr.Username)
}
