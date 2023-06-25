package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Чтение текста и конвертация в float64
func GetFloat() (float64, error) {
	fmt.Println("Hi")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	fmt.Println("Hi")
	return number, nil
}
