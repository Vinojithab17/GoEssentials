package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/demo_go/note"
	"example.com/demo_go/todo"
)

type saver interface {
	Save() error
}

type outPutable interface {
	saver
	Display()
}

func basic() {
	hashMap := map[string]string{
		"Name": "jahgshjdgahj",
	}
	fmt.Println(hashMap)
	// hashMap = append()
	hashMap["lastnake"] = "AB17"
	fmt.Println(hashMap)

	printter(1)
	printter(1.34)
	printter("aklsjlkajsd")
	title, content := getData()
	todoText := getInput("Enter The todo Text")

	userTodo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userTodo)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := summer(1, 2)
	fmt.Println(sum)
}

func outputData(data outPutable) error {
	data.Display()
	return saveData(data)
}
func getInput(promt string) string {
	fmt.Print(promt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func getData() (string, string) {
	title := getInput("Enter Title : ")
	content := getInput("Enter content : ")
	return title, content
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving failed")
	}
	fmt.Println("Succefully saved the data")
	return nil
}

func printter(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("The Integer Value is ", value)

	case float64:
		fmt.Println("The float Value is ", value)

	case string:
		fmt.Println("The string Value is ", value)

	}

}

// gerneris in golang
func summer[T int | float64 | string](a, b T) T {
	return a + b

}
