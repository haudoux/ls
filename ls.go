package main
//Readdir
//Readdirname
//isdir
import (
	"os"
	"io/ioutil"
	"fmt"
)

func main(){
	path := ""
	if len(os.Args) == 1 {
		path = "./"
	} else {
		path = os.Args[1]
	}
	readDirectory(path)
}
func readDirectory(path string){
	files, err := ioutil.ReadDir(path)
	checkError(err)
	//fmt.Println(files)
    for _, file := range files {
		if file.IsDir() {
			readDirectory(path + file.Name() + "/")
		} else {
			fmt.Println(path + file.Name())
		}
	}
}
func checkError(err error){
	if err != nil {
		fmt.Println("Error :")
		fmt.Println(err.Error())
		os.Exit(0)
	}
}