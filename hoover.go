package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

type Details struct {
	file      string
	directory string
	hash      string
}

func hashMD5File(filePath string) string {

	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String
}

func Header() {
	fmt.Println(`
		  ||
		  ||
		  ||
		  ||
		  ||
		  ||
		  ||     Here you go, sweep
		  ||     that up..............
		 /||\
		/||||\
		======         __|__
		||||||        / ~@~ \
		||||||       |-------|
		||||||       |_______|
`)
}

func GetCommandLineArgs() (string, string, string, string) {
	silencePtr := flag.String("s", "", "Flag to silence the binary")
	filePtr := flag.String("f", "", "example bad image file (example: 404_image.jpeg)")
	dirPtr := flag.String("d", "./", "directory of files to check (default: cwd)")
	md5Ptr := flag.String("h", "", "example md5sum of bad image file")
	flag.Parse()
	return *silencePtr, *filePtr, *dirPtr, *md5Ptr
}

func AddSlash(path string) string {
	if path[len(path)-1:] != "/" {
		path = path + "/"
	}
	return path
}

func DeleteWithHash(current Details) {
	if o, _ := regexp.MatchString("([a-fA-F0-9]{32})", current.hash); o != true {
		fmt.Println("Not valid md5 hash")
		os.Exit(3)
	}
	files, err := ioutil.ReadDir(current.directory)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if hashMD5File(current.directory+f.Name()) == current.hash {
			fmt.Println("Deleting file " + current.directory + f.Name())
			os.Remove(current.directory + f.Name())
		}
	}
}

func DeleteWithFile(current Details) {
	if _, err := os.Stat(current.file); os.IsNotExist(err) {
		fmt.Println("Error! " + current.file + " does not exist or is unaccessible!")
		os.Exit(3)
	}
	current.hash = hashMD5File(current.file)
	DeleteWithHash(current)
}

func Organization(current Details) {
	current.directory = AddSlash(current.directory)
	if current.file != "" && current.hash == "" {
		_ = ""
	} else if current.file == "" && current.hash != "" {
		_ = ""
	} else if current.file != "" && current.hash != "" {
		fmt.Println("Please exclusively use the md5 or the file, not both.")
		os.Exit(3)
	} else {
		fmt.Println("Either file or hash is required.")
		os.Exit(3)
	}
	if current.hash != "" {
		DeleteWithHash(current)
	} else if current.file != "" {
		DeleteWithFile(current)
	}
}

func main() {
	var current Details
	var silence string
	silence, current.file, current.directory, current.hash = GetCommandLineArgs()
	if silence == "" {
		Header()
	}
	Organization(current)
}
