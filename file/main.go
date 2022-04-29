package main

//
//import (
//	"fmt"
//	"log"
//	"os"
//	"syscall"
//
//	"github.com/docker/docker/pkg/reexec"
//)
//
//var fd int
//var err error
//
//func init() {
//
//	fd, err = syscall.Open("./1.txt", syscall.O_RDONLY, 0)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Println("my open fd ", fd)
//
//	log.Printf("init start, os.Args = %+v\n", os.Args)
//	reexec.Register("childProcess", childProcess)
//	if reexec.Init() {
//		os.Exit(0)
//	}
//
//}
//func childProcess() {
//
//	err := syscall.Setuid(500)
//	fmt.Println(err, "uid")
//	err = syscall.Setgid(500)
//	fmt.Println(err, "gid")
//
//	var buf []byte
//	buf = make([]byte, 20, 100)
//	n, err := syscall.Read(fd, buf)
//	fmt.Println(n, string(buf), err)
//	log.Println("childProcess", syscall.Getpid(), fd)
//}
//
//func main() {
//
//	log.Printf("main start, os.Args = %+v\n", os.Args)
//	cmd := reexec.Command("childProcess")
//	cmd.Stdin = os.Stdin
//	cmd.Stdout = os.Stdout
//	cmd.Stderr = os.Stderr
//	if err := cmd.Start(); err != nil {
//		log.Panicf("failed to run command: %s", err)
//	}
//	if err := cmd.Wait(); err != nil {
//		log.Panicf("failed to wait command: %s", err)
//	}
//	log.Println("main exit", syscall.Getpid())
//
//	var buf []byte
//	buf = make([]byte, 20, 100)
//	n, err := syscall.Read(fd, buf)
//	fmt.Println("main", n, string(buf), err, fd)
//	//base.OpenFile()
//	//base.ReadExistAndNotExistFile()
//}
