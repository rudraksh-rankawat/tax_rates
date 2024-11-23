package file_handling

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"example.com/practice_app/conversion"
)


type FileHandler struct {
	FileName string
}

func NewFileHandler(fileName string) (*FileHandler) {
	return &FileHandler {
		FileName: fileName,
	}
}

func (f FileHandler) LoadData() ([]float64) {
	file, err := os.Open(f.FileName)

	if err != nil {
		fmt.Println("Cannot open the file")
		fmt.Println(err)
		return nil
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	price_data_str := []string{}

	for scanner.Scan() {
		price_data_str = append(price_data_str, scanner.Text())
	}
	file.Close()
	
	result := conversion.ConvertStrToFloatMap(price_data_str)

	return result
}


func (f FileHandler) WriteJson(data interface{}, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return errors.New(err.Error())
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("failed to write JSON data: %w", err)
	}

	fmt.Printf("HEHEHEHE WRITTEN SUCCESSFULLY")
	return nil


}

