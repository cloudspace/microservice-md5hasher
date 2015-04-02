package main // import "github.com/cloudspace/microservice-md5hasher"

import  (
	"fmt"
	"os"
	"crypto/md5"
	"encoding/json"
)

func main() {
	unEncryptedInput := []byte(os.Args[1])
	result := md5.Sum(unEncryptedInput)
	var strResult = fmt.Sprintf("%x", result)
	var wrappedResults = map[string]string{
		"result": strResult,
	}

	encodedResults, _ := json.Marshal(wrappedResults)
	fmt.Println(string(encodedResults))
}