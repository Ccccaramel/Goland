package main

import (
	"fmt"
)

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero.
	dividee: %d
	divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

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
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100 / 10 = ", result)
	}
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMagis: ", errorMsg)
	}
	fmt.Println()
}

/*
错误处理:
	error 类型是一个接口
		type error interface{
			Error() string
		}
	使用 error.New 可返回一个错误信息
		func Sqrt(f float64) (float64, error){
			if f < 0 {
				return 0, errors.New("Math: square root of negative number.")
			}
		}
*/
