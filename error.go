// error.go
package main

import (
	"fmt"
)

type error interface {
	Error() string
}

//define a DivideError struct
type DivideError struct {
	dividee int
	divider int
}

//implement error interface
func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed,the divider is zero.
	dividee: %d
	divider: 0
	`
	return fmt.Sprintf(strFormat, de.dividee)
}

//define `int` type divide operation
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func main() {
	//normal case
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	//when the divider is zero,return error msg
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

}
