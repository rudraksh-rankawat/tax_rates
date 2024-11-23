package conversion

import "strconv"

func ConvertStrToFloatMap(str_map []string) ([]float64) {

	float_map := make([]float64, len(str_map))
	for priceIndex, price := range str_map {
		// price_data_float = append(price_data_float, strconv.ParseFloat(price, 64))
		float_map[priceIndex], _ = strconv.ParseFloat(price, 64)
	}
	return float_map
}