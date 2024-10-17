package reader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type inputData struct {
	Numbers []int `json:"numbers"`
}
// Read numbers from JSON file. UNmarshalled JSON data to "inputData" struct.
func ReadJSONFile(filename string) ([]int, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var jsonNums inputData
	if err := json.Unmarshal(data, &jsonNums); err != nil {
		return nil, err
	}

	return jsonNums.Numbers, nil
}
// Read numbers from STDIN
func ReadStdin() ([]int, error) {
	fmt.Println("Enter numbers separated by spaces.\nPress \"Enter\" to countinue: ")
	reader := bufio.NewReader(os.Stdin)
    data, err := reader.ReadString('\n')
    if err != nil {
        return nil, err
    }

	numberStrings := strings.Fields(string(data))
	var numbers []int
	for _, numStr := range numberStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, fmt.Errorf("invalid num format: %s", numStr)
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}