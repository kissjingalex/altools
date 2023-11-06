package interact

import (
	"bufio"
	"fmt"
	"os"
)

type InteractService struct {
}

var service = &InteractService{}

func NewInteractService() *InteractService {
	return service
}

func (svc *InteractService) Start() {
	var str string
	//使用os.Stdin开启输入流
	in := bufio.NewScanner(os.Stdin)
	for {
		in.Scan()
		str = in.Text()
		if str == "exit" {
			break
		}

		fmt.Printf("input=%s\n", str)

	}
}
