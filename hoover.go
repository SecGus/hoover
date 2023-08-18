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
        "os/exec"
        "regexp"
        "strings"
)

type Details struct {
        file      string
        directory string
        hash      string
        command   string
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
                  ||     Here you go, clean
                  ||     that up..............
                 /||\
                /||||\
                ======         __|__
                ||||||        / ~@~ \
                ||||||       |-------|
                ||||||       |_______|
`)
}

func GetCommandLineArgs() (string, string, string, string, string) {
        silencePtr := flag.String("s", "", "Flag to silence the binary")
        filePtr := flag.String("f", "", "example bad image file (example: 404_image.jpeg)")
        dirPtr := flag.String("d", "./", "directory of files to check (default: cwd)")
        md5Ptr := flag.String("h", "", "example md5sum of bad image file")
        commandPtr := flag.String("c", "", "command to run on each file matching md5")
        flag.Parse()
        return *silencePtr, *filePtr, *dirPtr, *md5Ptr, *commandPtr
}

func AddSlash(path string) string {
        if path[len(path)-1:] != "/" {
                path = path + "/"
        }
        return path
}

func ExecuteCommandWithHash(current Details) {
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
                        commandWithFile := strings.Replace(current.command, "{{file}}", current.directory+f.Name(), -1)
                        cmd := exec.Command("/bin/sh", "-c", commandWithFile)
                        err := cmd.Run()
                        if err != nil {
                                log.Fatalf("Failed to execute command: %s", err)
                        }
                }
        }
}

func ExecuteCommandWithFile(current Details) {
        if _, err := os.Stat(current.file); os.IsNotExist(err) {
                fmt.Println("Error! " + current.file + " does not exist or is unaccessible!")
                os.Exit(3)
        }
        current.hash = hashMD5File(current.file)
        ExecuteCommandWithHash(current)
}

func Organization(current Details) {
        // (Same as above, not shown here for brevity)

        if current.hash != "" {
                ExecuteCommandWithHash(current)
        } else if current.file != "" {
                ExecuteCommandWithFile(current)
        }
}

func main() {
        var current Details
        var silence string
        silence, current.file, current.directory, current.hash, current.command = GetCommandLineArgs()
        if silence == "" {
                Header()
        }
        Organization(current)
}
