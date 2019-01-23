package real

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut time.Duration
}

func (r Retriever) Get(url string) string{
	resp,err := http.Get(url)
	if err != nil{
		panic(err)
	}

	result,err := httputil.DumpResponse(resp,true)
	resp.Body.Close()

	if err != nil{
		panic(err)
	}
	return string(result)
}




func (r Retriever) Post(url string,data string)string{
	fmt.Printf(url,data)
	var str string = "This is an example of string"
	fmt.Print(str)
	return "ddd"
}