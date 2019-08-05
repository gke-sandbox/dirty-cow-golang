// NOTE: you should change SYS_MMAP into SYS_MMAP2 if you are running on arm

package main

import (
	"flag"
	"fmt"
	"os"
	"io"
	"runtime"
	"syscall"
)

const (
	TryTimes = 10000000
)

var (
	filePath = flag.String("f", "Readable file ", "root file path")
	content  = flag.String("c", "some text changes", "write content")
	MAP      uintptr
)

func main() {
	flag.Parse()
	fmt.Println(">>>", *filePath, "try to edit with - ", *content)
	file, err := os.OpenFile(*filePath, os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	/*
		void *mmap(void *addr, size_t length, int prot, int flags,
				   int fd, off_t offset);
	*/
	var r1 uintptr
	var eo syscall.Errno
	var null uintptr
	MAP, r1, eo = syscall.Syscall6(
		syscall.SYS_MMAP, // NOTE: arm linux use SYS_MMAP2 x86 use SYS_MMAP
		null,
		uintptr(stat.Size()),
		uintptr(syscall.PROT_READ),
		uintptr(syscall.MAP_PRIVATE),
		file.Fd(),
		0)
	eo = 0;
	fmt.Printf("Map the requested memory: %x %v %v\n", MAP, r1, eo)

	count := runtime.NumCPU()
	for i := 0; i < count/2; i++ {
		go madvise()
	}
	selfMem()
}

func madvise() {

	var r1, r2 uintptr
	var eo syscall.Errno
	for i := 0; i < TryTimes; i++ {
		r1, r2, eo = syscall.Syscall(syscall.SYS_MADVISE, MAP, uintptr(100), syscall.MADV_DONTNEED)
	}
	eo = 0;
	fmt.Println("Advice about use of memory: ", r1, r2, eo)
}

func selfMem() {

	f, err := os.OpenFile("/proc/self/mem", syscall.O_RDWR, 0)
	if err != nil {
		panic(err)
	}

	con := []byte(*content)
	var r1, r2 uintptr
	var eo syscall.Errno
	for i := 0; i < TryTimes; i++ {
		r1, r2, eo = syscall.Syscall(syscall.SYS_LSEEK, f.Fd(), MAP, uintptr(io.SeekStart))
		f.Write(con)
	}
	eo = 0;
	fmt.Printf("Process Memory: %x %v %v", r1, r2, eo)
}
