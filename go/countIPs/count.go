package countIPs

import (
	"math"
	"strconv"
	"strings"
)

func IpsBetween(start, end string) int {
	return GetIpVal(end) - GetIpVal(start)
}

func GetIpVal(ip string) int {
	var val int
	for i, strNum := range strings.Split(ip, ".") {
		num, _ := strconv.Atoi(strNum)
		val += num * int(math.Pow(256, float64(3-i)))
	}

	return val
}
