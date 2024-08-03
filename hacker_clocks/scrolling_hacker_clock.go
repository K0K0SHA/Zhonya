// filename: scrolling_hacker_clock.go
// A primitive clock that takes a terminal Window, where it scrolls infinitely, hogging a lot of resources
// Runs on any shell
// It's recommended to run this in a terminal which doesn't have logging or a long history
package main

import "time"
import "fmt"

var version = 0.6
var time_now = time.Now()

func main(){
        beavis := 0
        for beavis < 69{
                fmt.Println(time.Now())
                beavis = 0
        }
}

/*
TODO: 
[ ] implement better formatting  
[ ] implement coroutines (goroutines)  
[ ] implement interrupt handle  
*/
