package main
import "fmt"

func main() {
  fruits := [3]string{"apple", "orange", "banana"}
  for idx, val := range fruits {
     fmt.Println(idx, val)
  }
}