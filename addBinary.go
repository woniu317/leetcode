package main

import "fmt"

func main() {
	fmt.Println(addBinary("11", "1"))
}
func addBinary(a string, b string) string {
	result := []byte{}
	c := 0
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || c != 0; i, j = i-1, j-1 {
		if i >= 0 {
			c += int(a[i] - '0')
		}
		if j >= 0 {
			c += int(b[j] - '0')
		}
		result = append(result, byte(c)%2+'0')
		c /= 2
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
