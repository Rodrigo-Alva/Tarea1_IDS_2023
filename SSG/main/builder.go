package main

import (
	"bufio"
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

    out := make([]byte, 1024);

    fileReader := bufio.NewReader(file);

    fileReader.Read(out);

    fmt.Println(string(out));
}
