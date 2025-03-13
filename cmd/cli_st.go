package cmd

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var defaultPlan = fmt.Sprintf("%s/.plan.txt", os.Getenv("HOME"))

func CreatePlan() (loc string) {
	_, err := os.Create(fmt.Sprintf("%s/.plan.txt", os.Getenv("HOME")))

	if err != nil {
		fmt.Println(err)
	}

	return "~/.plan.txt"

}

func AddItem(item string) {
	fmt.Printf("Adding item %s", item)
	// f, err := os.Open()
	// err := os.WriteFile(fmt.Sprintf("%s/.plan.txt", os.Getenv("HOME")), []byte(fmt.Sprintf("[ ] %s", item)), 0644)
	// if err != nil {
	// 	panic(err)
	// }
	file, err := os.OpenFile(defaultPlan, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	check(err)

	_, errer := file.WriteString(fmt.Sprintf("\n[ ] %s", item))
	check(errer)

	file.Sync()
	defer file.Close()

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("hello")
	// bWriter := bufio.NewWriter(f)
	// _, error := bWriter.WriteString(item)
	// f.Sync()
	// if error != nil {
	// 	fmt.Println(error)
	// }
	// defer f.Close()
}

func CheckOffItem(item string) {
	fileContent := []string{}

	arg, err := strconv.Atoi(item)
	if err != nil {
		check(err)
	} else {
		//fmt.Println(arg)
		file, err := os.OpenFile(defaultPlan, os.O_CREATE|os.O_RDWR, os.ModeAppend)
		check(err)

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			fileContent = append(fileContent, scanner.Text())
			//fmt.Println(scanner.Text())
		}
		file.Close()
		if arg > len(fileContent) {
			return
		}
		fileContent[arg] = chkoff(fileContent[arg])
		fmt.Println(fileContent[arg])
		//fmt.Println(fileContent[arg+1])
		f, err := os.Create(defaultPlan)
		if err != nil {
			check(err)
		}
		//f.WriteString("\n")
		for _, str := range fileContent {
			f.WriteString(str + "\n")
		}

		f.Close()
	}
}

func DeleteItem(item string) {
	fileContent := []string{}
	arg, err := strconv.Atoi(item)
	if err != nil {
		check(err)
	} else {
		//fmt.Println(arg)
		file, err := os.OpenFile(defaultPlan, os.O_CREATE|os.O_RDWR, os.ModeAppend)
		check(err)

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			fileContent = append(fileContent, scanner.Text())
			//fmt.Println(scanner.Text())
		}
		file.Close()
		if arg > len(fileContent) {
			return
		}
		fmt.Println(fileContent)
		fileContent = slices.Delete(fileContent, arg, arg+1)
		fmt.Println(fileContent)
		f, err := os.Create(defaultPlan)
		if err != nil {
			check(err)
		}
		//f.WriteString("\n")
		for _, str := range fileContent {
			f.WriteString(str + "\n")
		}

		f.Close()
	}
}

func PrintAtOpen() {
	fmt.Println("add 'cat .plan.txt' into your .bashrc file. that should work.\n the ablilty to programatically do that is coming soon!")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func chkoff(item string) (ret string) {
	// arg: "[ ] something, something"
	// want to return "[x] something, something"
	return strings.Replace(item, "[ ]", "[x]", 1)
}

// func del(slice []string) (retSlice []string) {
// 	slices.Delete()
// }
