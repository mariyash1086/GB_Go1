package Configuration

import (
	"fmt"
	"github.com/namsral/flag"
	"net/url"
)

type Configuration struct {
	myUrl  string
	myPort string
}

var (
	MyUrlflag  = flag.String("myUrl", "", "URL of configuration")
	MyPortflag = flag.String("myPort", "", "Port of configuration")
)

func Load() (*Configuration, error) {

	var result Configuration
	var err error

	flag.Parse()
	_, err = url.ParseRequestURI(*MyUrlflag)

	result.myUrl = *MyUrlflag
	result.myPort = *MyPortflag

	if err != nil {
		fmt.Println("it is not URL ")
	}
	return &result, err
}
