package main

import (
	//"database/sql/driver"
	"fmt"
)

// define struct
type GolangLTDError struct {
	GolangLTDee int
	GolangLTDer int
}

// error interface
func (this *GolangLTDError) Error() string{
	strMsg := "www.Golang.LTD error!"
      return fmt.Sprintf(strMsg,this.GolangLTDee)
}
// int divsion
func Golangltd(varGolangee int, varGolanger int) (result int, errorMsg string) {

	if varGolanger == 0 {
		dData := GolangLTDError{
			varGolangee,
			varGolanger,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varGolangee / varGolanger, ""
	}
}

func main() {
	  // normally
	  if result,errorMsg := Golangltd(100,  10); errorMsg == ""{
	  	fmt.Println( "100/10 = " ,result)
	  }
    // error
	if _,errorMsg := Golangltd( 100,  0); errorMsg != ""{
		fmt.Println( "errorMsg is : " ,errorMsg)
	}
}