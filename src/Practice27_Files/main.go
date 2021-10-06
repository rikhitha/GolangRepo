package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func getnumber() int {
	var num int
	fmt.Printf(" \n choose \n 1:ioutil \n 2: absolute file path \n 3: command line argument \n 4: creating / Reading / writing (log.faltalf,writeString,Name,Len,ioutil.Readfile) \n 5: Reading files in chunks using bufio package")
	fmt.Scanln(&num)

	return num
}

func CreateFile() {

	fmt.Println("Creating and writing into file in go lang")

	filecontent, error := os.Create("file1.txt")

	if error != nil {
		log.Fatalf("Failed to create file :%s ", error)
	}

	defer filecontent.Close()

	filecontentlen, errorcon := filecontent.WriteString("Hello hai Contents in the filecontent file ")

	if errorcon != nil {
		log.Fatalf("\n Failed to write into file : %s ", errorcon)
	}

	fmt.Printf(" \n length of contents in file : %d ", filecontentlen)

	fmt.Printf("\n \n Name of the file : %s  ", filecontent.Name())

}

func ReadFile() {

	fmt.Println("\n Reading contents in the file")

	fileName := "file1.txt"

	data, error2 := ioutil.ReadFile("file1.txt")

	if error2 != nil {
		log.Panicf("failed reading data from file: %s", error2)
	}
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)

}

func ReadBufio() {

	filecontent, error := os.Create("file1.txt")

	if error != nil {
		log.Fatalf("Failed to create file :%s ", error)
	}

	filecontentlen, errorcon := filecontent.WriteString("Hello hai Contents in the filecontent file ")

	if errorcon != nil {
		log.Fatalf("\n Failed to write into file : %s ", errorcon)
	}

	f1 := filecontent.Name()
	f, err := os.Open(f1)
	if err != nil {
		log.Fatal(err)
	}
	// defer func() {
	// 	if err = f.Close(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()
	r := bufio.NewReader(f)

	b := make([]byte, 3)
	for {
		n, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b[0:n]))
	}
	fmt.Printf("length of file1: %d", filecontentlen)

}

func main() {

	fmt.Println("Reading files")

	switch n := getnumber(); {
	case n == 1:

		fmt.Println("Simple Read File")

		content, Errors := ioutil.ReadFile("C:/Users/user/Desktop/links.txt")

		if Errors != nil {
			fmt.Println("Errorsl:==", Errors)
		}

		fmt.Println("Contents in File", string(content))

	case n == 2:
		fmt.Println("absolute file path")

		content, Errors := ioutil.ReadFile("C:/Users/user/Desktop/links.txt")

		if Errors != nil {
			fmt.Println("Errorsl:==", Errors)
		}

		fmt.Println("Contents in File", string(content))

	case n == 3:
		fmt.Println("Passing the file path as a command line flag")

		fptr := flag.String("fpath", "links.txt", "file path to read from")
		flag.Parse()
		fmt.Println("value of fpath is", *fptr)

	case n == 4:

		fmt.Println("Creating , Reading , Writing")

		go CreateFile()

		time.Sleep(200 * time.Millisecond)

		go ReadFile()
		time.Sleep(200 * time.Millisecond)

		fmt.Println("Done ")

	case n == 5:
		fmt.Println("Reading Files in chunks using bufio package")

		ReadBufio()

	}
}
