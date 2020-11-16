package SFSReader

// 資料型態
type DataWrapper struct {
	DataType int
	Data     interface{}
}

// 型態表
const (
	// 0 開始
	NULL = iota
	BOOL
	BYTE
	SHORT
	INT
	LONG
	FLOAT
	DOUBLE
	UTF_STRING

	// 9 開始是 Array
	BOOL_ARRAY
	BYTE_ARRAY
	SHORT_ARRAY
	INT_ARRAY
	LONG_ARRAY
	FLOAT_ARRAY
	DOUBLE_ARRAY
	UTF_STRING_ARRAY
	SFSDATA_ARRAY

	// 18
	SFSOBJECT
)

func DecodeObject(HexData []byte, Index int) (DataWrapper, int) {
	var wrapper DataWrapper

	// 讀取資料
	var DataTypeInterface interface{}
	DataTypeInterface, Index = ReadingByte(HexData, Index)
	wrapper.DataType = int(DataTypeInterface.(byte))

	// 根據形態來列表
	switch wrapper.DataType {
	case NULL:
		wrapper.Data = nil

	// Bool
	case BOOL:
		wrapper.Data, Index = ReadingBool(HexData, Index)
	case BOOL_ARRAY:
		wrapper.Data, Index = ReadingBoolArray(HexData, Index)

	// Byte
	case BYTE:
		wrapper.Data, Index = ReadingByte(HexData, Index)
	case BYTE_ARRAY:
		wrapper.Data, Index = ReadingByteArray(HexData, Index)

	// Short
	case SHORT:
		wrapper.Data, Index = ReadingShort(HexData, Index)
	case SHORT_ARRAY:
		wrapper.Data, Index = ReadingShortArray(HexData, Index)

	// Int
	case INT:
		wrapper.Data, Index = ReadingInt(HexData, Index)
	case INT_ARRAY:
		wrapper.Data, Index = ReadingIntArray(HexData, Index)

	// Long
	case LONG:
		wrapper.Data, Index = ReadingLong(HexData, Index)
	case LONG_ARRAY:
		wrapper.Data, Index = ReadingLongArray(HexData, Index)

	// Float
	case FLOAT:
		wrapper.Data, Index = ReadingFloat(HexData, Index)
	case FLOAT_ARRAY:
		wrapper.Data, Index = ReadingFloatArray(HexData, Index)

	// Double
	case DOUBLE:
		wrapper.Data, Index = ReadingDouble(HexData, Index)
	case DOUBLE_ARRAY:
		wrapper.Data, Index = ReadingDoubleArray(HexData, Index)

	// 17
	case SFSDATA_ARRAY:
		wrapper.Data, Index = ReadingSFSDataArray(HexData, Index)

	//18
	case SFSOBJECT:
		Index--
		wrapper.Data, Index = ReadingSFSObject(HexData, Index)
		wrapper = wrapper.Data.(DataWrapper)

	default:
		panic(wrapper.DataType)
	}
	return wrapper, Index
}

// 拿取 Array 資料
func SFSGetArray(Data interface{}) []DataWrapper {
	wrapper := Data.(DataWrapper)
	if wrapper.DataType == SFSDATA_ARRAY {
		return wrapper.Data.([]DataWrapper)
	}
	panic("Type Wrong !!")
}

// 拿取 Int Array 的資料
func SFSGetIntArray(Data interface{}) []int32 {
	wrapper := Data.(DataWrapper)
	if wrapper.DataType == INT_ARRAY {
		return wrapper.Data.([]int32)
	}
	panic("Type Wrong !!")
}

// 拿 Object 資料
func SFSGetObject(Data interface{}) map[string]interface{} {
	wrapper := Data.(DataWrapper)
	if wrapper.DataType == SFSOBJECT {
		return wrapper.Data.(map[string]interface{})
	}
	panic("Type Wrong !!")
}

// func SFSGetFromArrayIndex(Data interface{}) DataWrapper
