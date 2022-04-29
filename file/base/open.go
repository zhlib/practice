package base

import (
	"fmt"
	"syscall"
)

func OpenFile() {
	fd, err := syscall.Open("1.txt",
		syscall.O_CREAT|
		syscall.O_RDWR|
		syscall.O_CLOEXEC, 0666)
	if err != nil {
		fmt.Printf("syscall.Open error(%v)", err)
	}
	fmt.Printf("fd is %d",fd)

	n,err := syscall.Write(fd,[]byte("first byte"))
	fmt.Println(n,err)
}




func ReadExistAndNotExistFile() {
	var fd int
	fd, err := syscall.Open("./1.txt", syscall.O_RDONLY, 0)
 	if err != nil {
		fmt.Printf("%d %s\n", fd,err) // 3
	}else {
		fmt.Println("fd is :", fd)
	}
}


func O_CLOEXEC() {

	//fd ,err := syscall.Open("/etc/shadow",syscall.O_RDONLY,0)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//syscall.ForkExec()
}
