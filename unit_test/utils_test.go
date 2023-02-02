package unit

import (
	"LibraryManagementSys/utils"
	"fmt"
	"testing"
	"time"
)

func TestParseTimeDur(t *testing.T) {
	dur, err := utils.ParseTimeDur("1d")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		now := time.Now()
		fmt.Println(now)
		fmt.Println(now.Add(dur))
	}

	fmt.Println("------------------")

	dur, err = utils.ParseTimeDur("1h")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		now := time.Now()
		fmt.Println(now)
		fmt.Println(now.Add(dur))
	}

	fmt.Println("------------------")

	dur, err = utils.ParseTimeDur("1m")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		now := time.Now()
		fmt.Println(now)
		fmt.Println(now.Add(dur))
	}
}
