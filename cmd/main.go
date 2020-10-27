package main

import (
        "context"
        "log"

        "github.com/KhurshedUlugov/search/pkg/search"
)

//"regexp"
func main() {

        files := []string{"data/text.txt", "data/text1.txt"}

        root := context.Background()
        //      ctx, _ := context.WithCancel(root)

        aa := search.All(root, "Khurik", files)

        log.Print(aa)

}

