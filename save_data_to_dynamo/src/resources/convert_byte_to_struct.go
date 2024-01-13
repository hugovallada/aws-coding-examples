package resources

import "encoding/json"

func ConvertByteToStruct[T any](data []byte) (obj T, err error) {
	err = json.Unmarshal(data, &obj)
	return obj, err
}
