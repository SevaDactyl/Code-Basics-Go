/*
Написать интересный код самостоятельно на текущем уровне будет затруднительно, поэтому скопируйте код ниже.

wg := sync.WaitGroup{}

for i := 0; i < 3; i++ {
    wg.Add(1)
    go func() {
        fmt.Println("Go!")
        wg.Done()
    }()
}

wg.Wait()
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	// BEGIN (write your solution here)
	wg := sync.WaitGroup{}

for i := 0; i < 3; i++ {
    wg.Add(1)
    go func() {
        fmt.Println("Go!")
        wg.Done()
    }()
}

wg.Wait()
	// END
}
