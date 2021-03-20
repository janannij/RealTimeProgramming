package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func spawnBackup() {
	backupCmd := exec.Command("osascript", "-e", `tell app "Terminal" to do script "go run Documents/TTK4145sanntid/exercise-7-janannij/main.go"`)
	err := backupCmd.Start()
	check(err)
}

//Primary setup
func primary(counter int) {
	for {
		counter++
		s1 := strconv.FormatInt(int64(counter), 10)
		insertData := []byte(s1)
		ioutil.WriteFile("phnxfile.txt", insertData, 0644)
		fmt.Println(s1)
		time.Sleep(time.Second)
	}
}

func CheckWriteToFile() int {
	for {
		d1, err1 := ioutil.ReadFile("phnxfile.txt")
		check(err1)
		data1, error1 := strconv.Atoi(string(d1))
		check(error1)

		time.Sleep(time.Second)

		d2, err2 := ioutil.ReadFile("phnxfile.txt")
		check(err2)
		data2, error2 := strconv.Atoi(string(d2))
		check(error2)

		if data1 == data2 {
			return data1
		}
	}
}

func main() {
	if _, err := os.Stat("phnxfile.txt"); os.IsNotExist(err) {
		filename, err := os.Create("phnxfile.txt")
		check(err)
		defer filename.Close() //defer closing the file
		_, err1 := filename.WriteString(fmt.Sprintf("%d", 0))
		check(err1)
	} else {
		filename, err := os.Open("phnxfile.txt")
		check(err)
		defer filename.Close()
	}

	count := CheckWriteToFile()

	spawnBackup()

	for {
		primary(count)
	}

}
