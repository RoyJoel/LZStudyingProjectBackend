package test

import (
	"fmt"
	"testing"

	"github.com/RoyJoel/LZStudyingProject/package/utils"
)

func TestS(t *testing.T) {
	str := utils.NowTimeStamp()
	fmt.Println(str)

	timeStr := utils.TimeStamp2NowTimeStr(str)
	fmt.Println(timeStr)

	stamp := utils.NowTimeStr2TimeStamp(timeStr)
	fmt.Println(stamp)
}
