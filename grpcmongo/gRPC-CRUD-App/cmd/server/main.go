package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	pb "grpcmongo/gRPC-CRUD-App/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Change the host here only
const (
	address = "127.0.0.1:50052"
)

func main() {
	// Connect to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := pb.NewCRUDClient(conn)

	// This is a menu-based application

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// Menu-based program allowing the user to choose from CRUD
	fmt.Println("\nWelcome to a simple gRPC/MongoDB based app that performs CRUD",
		" operations!")
	fmt.Println("Enter the one of the folliwing choices below:")
	fmt.Print("1 to create an item; 2 to read; 3 to update and 4 to remove: ")

	choice := bufio.NewReader(os.Stdin)
	text, _ := choice.ReadString('\n')

	fmt.Print(text)
	switch text = strings.TrimSpace(text); text{
	// switch text {
	// switch strings.TrimSpace(text) {
	case "1":
		// CreateItem operation
		// Read the name
		fmt.Print("\nEnter the name: ")
		name := bufio.NewReader(os.Stdin)
		n, _ := name.ReadString('\n')
		n = strings.Trim(n, "\n")

		// Read the ID
		fmt.Print("Enter the ID: ")
		id := bufio.NewReader(os.Stdin)
		i, _ := id.ReadString('\n')
		i = strings.Trim(i, "\n")

		// Read the category
		fmt.Print("Enter the Category: ")
		category := bufio.NewReader(os.Stdin)
		cat, _ := category.ReadString('\n')
		cat = strings.Trim(cat, "\n")
		fmt.Println(cat)
		// text = strings.TrimSpace(text)
		catInt, _ := strconv.Atoi(cat)

		// Populate the Employee struct
		// item, err := c.CreateItem(ctx, &pb.Employee{Name: n, Id: i,
		// 	Category: int32(catInt)})
			item, err := c.CreateItem(ctx, &pb.Employee{Name: n, Id: i,
				Category: int32(catInt)})
			
			fmt.Println(reflect.TypeOf(catInt))
		if err != nil {
			log.Fatalf("Could not create a new item: %v", err)
		}
		// fmt.Println("\nInserted", n, "with the ID", item.Id, "and category", catInt)
		fmt.Println("\nInserted", n, "with the ID", item.Id, "and category", cat)


	case "2":
		// ReadItem operation
		fmt.Print("\nEnter the ID: ")
		id := bufio.NewReader(os.Stdin)
		i, _ := id.ReadString('\n')
		i = strings.Trim(i, "\n")

		read, err := c.ReadItem(ctx, &pb.ID{Id: i})
		if err != nil {
			log.Fatalf("Error reading the item: %v", err)
		}
		fmt.Println("\nItem found!")
		fmt.Println("Name:", read.Name)
		fmt.Println("ID:", read.Id)
		fmt.Println("Category:", read.Category)

	case "3":
		// UpdateItem operation
		// Read the ID
		fmt.Print("\nEnter the existing ID: ")
		id := bufio.NewReader(os.Stdin)
		i, _ := id.ReadString('\n')
		i = strings.Trim(i, "\n")

		// Read the name
		fmt.Print("Enter the new name: ")
		name := bufio.NewReader(os.Stdin)
		n, _ := name.ReadString('\n')
		n = strings.Trim(n, "\n")

		// Read the category
		fmt.Print("Enter the new category: ")
		category := bufio.NewReader(os.Stdin)
		cat, _ := category.ReadString('\n')
		cat = strings.Trim(cat, "\n")
		catInt, err := strconv.Atoi(cat)

		up, err := c.UpdateItem(ctx, &pb.Employee{Name: n, Id: i, Category: int32(catInt)})
		if err != nil {
			log.Fatalf("Error updating the item: %v", err)
		}
		log.Printf("\nItem updated with the ID: %s", up.Id)

	case "4":
		// DeleteItem operation
		// Ignoring the error - should always be successful regardless of the
		// implicit find result
		// Read the ID
		fmt.Print("\nEnter the existing ID: ")
		id := bufio.NewReader(os.Stdin)
		i, _ := id.ReadString('\n')
		i = strings.Trim(i, "\n")
		del, _ := c.DeleteItem(ctx, &pb.ID{Id: i})
		if del != nil {
			log.Printf("\nItem with the ID %s deleted", del.Id)
		} else {
			log.Printf("Successful delete (even though ID didn't exist)")
		}

	default:
		fmt.Println("\nWrong option!")
	}
}
