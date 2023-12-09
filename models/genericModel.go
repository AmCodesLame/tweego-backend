package models

import (
	"fmt"
)

type Msg struct {
	Response string
}

func ReturnMsg() Msg {
	var test_var Msg
	test_var.Response = "Working"
	return test_var
}

func ReturnMsgParam(param string) Msg {
	var test_var Msg
	test_var.Response = fmt.Sprintf("Hello %v", param)
	return test_var
}
