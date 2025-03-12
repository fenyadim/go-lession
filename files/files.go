package files

import (
	"fmt"
	"os"
)

func ReadFile() {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(file))
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File created successfully")
}
