package main

import (
	//"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type HTMLnode struct {
    tagName string;
    id string;
    className string;
    children []*HTMLnode;
}

func main() {
    fmt.Println("Opening file...");

    file, fileErr := os.Open("./test.md");
    if (fileErr != nil) {
        log.Fatal(fileErr);
    }

    defer func() {
        closingErr := file.Close();
        if (closingErr != nil) {
            log.Fatal(closingErr);
        }
    }();

    //out := make([]byte, 1024);
    //fileReader := bufio.NewReader(file);
    //fileReader.Read(out);
    //fmt.Println(string(out));
    str, err := getLine(*file)

    for ; err == nil; {
        fmt.Println(str);
        str, err = getLine(*file)
    }
}

func getLine(file os.File) (string, error) {
    var str string;
    tempbuff := make([]byte, 1);
    for {
        _, err := file.Read(tempbuff);
        if err == io.EOF {
            return "", errors.New("reached end of file");
        }
        if string(tempbuff) == "\n" { break; }
        str += string(tempbuff);
    }
    return str, nil
}
