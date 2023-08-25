package util

import "reflect"

func IsEmpty(value interface{}) bool {
	if value == nil {
		return true
	}

	switch v := value.(type) {
	case string:
		return v == "" // jika ""
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return v == 0 // jika nilainya hanya 0
	case float32, float64:
		return v == 0.0 // jika nilainya hanya 0
	case map[string]interface{}:
		return len(v) == 0 // panjang string value bertipe interface dari key tidak boleh kosong
	case map[string]string:
		return len(v) == 0 // panjang string value bertipe string dari key tidak boleh kosong
	case map[string]int:
		return len(v) == 0 // panjang string value bertipe int dari key tidak boleh kosong
	default:
		return reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
	} // memeriksa apakah value adalah sama dengan nilai nol sesuai dengan tipe datanya. Jika value adalah nol atau nilai-nilai nol seperti 0, false, nil, atau nilai-nilai lain yang sesuai dengan tipe data, maka pernyataan ini akan mengembalikan true, menandakan bahwa nilai tersebut dianggap kosong.
}
