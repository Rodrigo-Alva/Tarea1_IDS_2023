package main

import (
	//"bufio"
	"fmt"
	"log"
	"os"
)

type node struct {
    tagName string;
    children []*node;
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
    for i := 0; i < 4; i++ {
        fmt.Println(getLine(*file));
    }
}

func getLine(file os.File) (string, error) {
    var str string;
    tempbuff := make([]byte, 1);
    for {
        file.Read(tempbuff);
        if string(tempbuff) == "\n" { break; }
        str += string(tempbuff);
    }
    return str, nil
}
