package main

import (
	"github.com/ew1l/wb-l2/develop/dev08/shell"
)

func main() {
	s := shell.NewShell()

	s.Run()
}

// emil@DESKTOP-NJQTAVE:/mnt/c/Users/emil/Documents/WB/L2/wb-l2/develop/dev08$ ./cmd/shell
// $ pwd
// /mnt/c/Users/emil/Documents/WB/L2/wb-l2/develop/dev08
// $
// $ echo Hello
// Hello
// $
// $
// $
// $ ls
// cmd  go.mod  main.go  shell
// $ cd shell
// $
// $ ls -lah
// total 4.0K
// drwxrwxrwx 1 emil emil 4.0K Mar 13 15:53 .
// drwxrwxrwx 1 emil emil 4.0K Mar 13 15:55 ..
// -rwxrwxrwx 1 emil emil 1.3K Mar 13 15:45 shell.go
// $
// $ cat ../go.mod
// module github.com/ew1l/wb-l2/develop/dev08
// 
// go 1.17
// $
// $ cd ..
// $ ls
// cmd  go.mod  main.go  shell
// $
// $ ps
//   PID TTY          TIME CMD
//  2230 pts/0    00:00:00 bash
//  4396 pts/0    00:00:00 shell
//  4407 pts/0    00:00:00 ps
// $
// $ kill 4396
// Terminated
