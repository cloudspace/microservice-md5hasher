package main // import "github.com/cloudspace/microservice-md5hasher"

import  (
	"fmt"
	"os"
	"crypto/md5"
)

func main() {
	unEncryptedInput := []byte(os.Args[1])
	result := md5.Sum(unEncryptedInput)
	fmt.Printf("%x", result)
}