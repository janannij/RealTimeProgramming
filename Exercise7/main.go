package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

var counter int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func spawnBackup() {
	//backupCmd :=
	//(exec.Command("osascript", "-e", `tell app "Terminal" to do script "go run Documents/TTK4145sanntid/exercise-7-janannij/main.go"`).Start())
	//backupCmd.Start()
}

func main() {
	filename, err := os.Create("phnxfile.txt")
	check(err)
	defer filename.Close() // Defer closing the file
	primary := false

	//Backup setup
	if !(primary) {
		_, err := ioutil.ReadFile("phnxfile.txt")
		fmt.Println(err)
		if err == nil {
			primary = true
			fmt.Println("Primary")
			spawnBackup()
		} else {
			//s1 := strconv.FormatInt(int64(counter), 10)
			//insertData := []byte(s1)
			//ioutil.WriteFile("phnxfile.txt", insertData, 0644)
		}

	}

	//Primary setup
	for i := 0; i < 5; i++ {
		counter++
		s1 := strconv.FormatInt(int64(counter), 10)
		insertData := []byte(s1)
		ioutil.WriteFile("phnxfile.txt", insertData, 0644)
		fmt.Println(s1)
		time.Sleep(time.Second)
	}
}
