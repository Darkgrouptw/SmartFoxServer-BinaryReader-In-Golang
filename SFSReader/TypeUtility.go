package SFSReader

import "encoding/hex"
import "strconv"

import "../Common"

// Index: 1
// 讀取 Byte (1 byte)
func ReadingBool(HexData []byte, Index int) (interface{}, int) {
	HexString := hex.EncodeToString(HexData[Index : Index+1])
	BoolData, Error := strconv.ParseBool(HexString)
	Common.Check(Error)

	return BoolData, Index + 1
}

// Index: 2
// 讀取 Byte (1 byte)
func ReadingByte(HexData []byte, Index int) (interface{}, int) {
	HexString := hex.EncodeToString(HexData[Index : Index+1])
	UIntData, Error := strconv.ParseUint(HexString, 16, 8)
	Common.Check(Error)

	return byte(UIntData), Index + 1
}

// Index: 3
// 讀取 Short (2 bytes)
func ReadingShort(HexData []byte, Index int) (interface{}, int) {
	HexString := hex.EncodeToString(HexData[Index : Index+2])
	UIntData, Error := strconv.ParseUint(HexString, 16, 16)
	Common.Check(Error)

	return int16(UIntData), Index + 2
}

// 讀取 UShort (2 byte)
func ReadingUShort(HexData []byte, Index int) (interface{}, int) {
	HexString := hex.EncodeToString(HexData[Index : Index+2])
	UIntData, Error := strconv.ParseUint(HexString, 16, 16)
	Common.Check(Error)

	return uint16(UIntData), Index + 2
}

// Index: 4
// 讀取 Int (4 bytes)
func ReadingInt(HexData []byte, Index int) (interface{}, int) {
	HexString := hex.EncodeToString(HexData[Index : Index+4])
	IntData, Error := strconv.ParseUint(HexString, 16, 32)
	Common.Check(Error)

	return int32(IntData), Index + 4
}

// Index: 5
// 讀取 Long (8 byte)
func ReadingLong(HexData []byte, Index int) (interface{}, int) {
	HexString := hex.EncodeToString(HexData[Index : Index+8])
	LongData, Error := strconv.ParseUint(HexString, 16, 64)
	Common.Check(Error)

	return int64(LongData), Index + 8
}

// Index: 6
// 讀取 Float (4 byte)
func ReadingFloat(HexData []byte, Index int) (interface{}, int) {
	HexString := hex.EncodeToString(HexData[Index : Index+4])
	FloatData, Error := strconv.ParseFloat(HexString, 32)
	Common.Check(Error)

	return float32(FloatData), Index + 4
}

// Index: 7
// 讀取 Double (4 byte)
func ReadingDouble(HexData []byte, Index int) (interface{}, int) {
	HexString := hex.EncodeToString(HexData[Index : Index+8])
	DoubleData, Error := strconv.ParseFloat(HexString, 32)
	Common.Check(Error)

	return float64(DoubleData), Index + 8
}

// Index: 8
// 讀取 UTF_STRING
func ReadingUTFString(HexData []byte, Index int) (string, int) {
	StringCount, Index := ReadingUShort(HexData, Index)

	// 抓取文字
	ResultText := string(HexData[Index : Index+int(StringCount.(uint16))])
	Index += int(StringCount.(uint16))

	return ResultText, Index
}

// Index: 9
// 讀取 Bool Array
func ReadingBoolArray(HexData []byte, Index int) (interface{}, int) {
	var ArraySize interface{}
	ArraySize, Index = ReadingShort(HexData, Index)
	BoolDataArray := make([]bool, ArraySize.(int16))
	for i := 0; i < int(ArraySize.(int16)); i++ {
		var Temp interface{}
		Temp, Index = ReadingByte(HexData, Index)
		BoolDataArray[i] = Temp.(bool)
	}
	return BoolDataArray, Index
}

// Index: 10
// 讀取 Byte Array
func ReadingByteArray(HexData []byte, Index int) (interface{}, int) {
	var ArraySize interface{}
	ArraySize, Index = ReadingInt(HexData, Index)
	ByteDataArray := make([]byte, ArraySize.(int32))
	for i := 0; i < int(ArraySize.(int32)); i++ {
		var Temp interface{}
		Temp, Index = ReadingByte(HexData, Index)
		ByteDataArray[i] = Temp.(byte)
	}
	return ByteDataArray, Index
}

// Index: 11
// 讀取 Short Array
func ReadingShortArray(HexData []byte, Index int) (interface{}, int) {
	var ArraySize interface{}
	ArraySize, Index = ReadingInt(HexData, Index)
	ShortDataArray := make([]int16, ArraySize.(int16))
	for i := 0; i < int(ArraySize.(int16)); i++ {
		var Temp interface{}
		Temp, Index = ReadingShort(HexData, Index)
		ShortDataArray[i] = Temp.(int16)
	}
	return ShortDataArray, Index
}

// Index: 12
// 讀取 Int Array
func ReadingIntArray(HexData []byte, Index int) (interface{}, int) {
	var ArraySize interface{}
	ArraySize, Index = ReadingShort(HexData, Index)
	IntDataArray := make([]int32, ArraySize.(int16))
	for i := 0; i < int(ArraySize.(int16)); i++ {
		var Temp interface{}
		Temp, Index = ReadingInt(HexData, Index)
		IntDataArray[i] = Temp.(int32)
	}
	return IntDataArray, Index
}

// Index: 13
// 讀取 Long Array
func ReadingLongArray(HexData []byte, Index int) (interface{}, int) {
	var ArraySize interface{}
	ArraySize, Index = ReadingShort(HexData, Index)
	LongDataArray := make([]int64, ArraySize.(int16))
	for i := 0; i < int(ArraySize.(int16)); i++ {
		var Temp interface{}
		Temp, Index = ReadingInt(HexData, Index)
		LongDataArray[i] = Temp.(int64)
	}
	return LongDataArray, Index
}

// Index: 14
// 讀取 Float Array
func ReadingFloatArray(HexData []byte, Index int) (interface{}, int) {
	var ArraySize interface{}
	ArraySize, Index = ReadingShort(HexData, Index)
	FloatDataArray := make([]float32, ArraySize.(int16))
	for i := 0; i < int(ArraySize.(int16)); i++ {
		var Temp interface{}
		Temp, Index = ReadingFloat(HexData, Index)
		FloatDataArray[i] = Temp.(float32)
	}
	return FloatDataArray, Index
}

// Index: 15
// 讀取 Double Array
func ReadingDoubleArray(HexData []byte, Index int) (interface{}, int) {
	var ArraySize interface{}
	ArraySize, Index = ReadingShort(HexData, Index)
	DoubleDataArray := make([]float64, ArraySize.(int16))
	for i := 0; i < int(ArraySize.(int16)); i++ {
		var Temp interface{}
		Temp, Index = ReadingFloat(HexData, Index)
		DoubleDataArray[i] = Temp.(float64)
	}
	return DoubleDataArray, Index
}

// Index: 17
// 讀取 Array
func ReadingSFSDataArray(HexData []byte, Index int) ([]DataWrapper, int) {
	var wrapper = []DataWrapper{}

	ArrayCount, Index := ReadingShort(HexData, Index)
	for i := 0; i < int(ArrayCount.(int16)); i++ {
		var TempWrapper DataWrapper
		TempWrapper, Index = DecodeObject(HexData, Index)
		wrapper = append(wrapper, TempWrapper)
	}
	return wrapper, Index
}

// Index: 18
// 讀取 Class(Object)
func ReadingSFSObject(HexData []byte, Index int) (DataWrapper, int) {
	if len(HexData) < 3+Index {
		// 長度太短
		panic("Smartfox Data Object is invalid because of length !!")
	}

	// 一開始先看 Header
	HeaderInt, Index := ReadingByte(HexData, Index)
	if HeaderInt.(byte) != 18 {
		panic("The Header (" + strconv.FormatInt(int64(HeaderInt.(byte)), 10) + ") is invalid")
	}

	// 接著判斷他的 Size 然後重覆去讀取 (Short)
	HeaderInt, Index = ReadingShort(HexData, Index)
	if HeaderInt.(int16) < 0 {
		panic("SFSObejct Size Error because of negative value(" + strconv.FormatInt(int64(HeaderInt.(int16)), 10) + ")!!")
	}

	// 接著去跑每一筆資料
	var wrapper DataWrapper
	wrapper.DataType = 18

	MapData := make(map[string]interface{})
	wrapper.Data = MapData

	for i := 0; i < int(HeaderInt.(int16)); i++ {
		var PropertyName string
		PropertyName, Index = ReadingUTFString(HexData, Index)

		// var wrapper2 DataWrapper
		// wrapper2, ReadingIndex = DecodeObject(HexData, ReadingIndex)
		MapData[PropertyName], Index = DecodeObject(HexData, Index)

		// panic(wrapper.Data)
		// ResultData[PropertyName] = wrapper
	}
	return wrapper, Index
}
