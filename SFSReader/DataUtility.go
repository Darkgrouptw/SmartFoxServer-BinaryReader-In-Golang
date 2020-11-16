package SFSReader

import "bytes"
import "compress/zlib"
import "encoding/hex"
import "io/ioutil"
import "../Common"

// 解縮縮資料
func DecompressData(DataStr string) []byte {
	// 轉成 Byte Array
	ByteData, Error := hex.DecodeString(DataStr)
	Common.Check(Error)

	// 這裡會使用 zlib 的原因是因為
	// Flate 的前面兩個 Byte 是代表使用 Default 壓縮，且沒帶 Dictionary
	// 所以你也可以使用
	// reader, Error := flate.NewReader(bytes.NewReader(ByteData[2:]))
	//
	// 來源：
	// https://stackoverflow.com/questions/29513472/golang-compress-flate-module-cant-decompress-valid-deflate-compressed-http-b
	reader, Error := zlib.NewReader(bytes.NewReader(ByteData))
	defer reader.Close()
	Common.Check(Error)
	DecodeBytes, Error := ioutil.ReadAll(reader)
	Common.Check(Error)

	return DecodeBytes
	// return hex.EncodeToString(DecodeBytes)
}

// 壓縮資料
// func CompressData(DataStr string) string {
// 	// 轉成 Byte Array
// 	ByteData, Error := hex.DecodeString(DataStr)
// 	Common.Check(Error)

// 	Buffer := bytes.NewReader(ByteData)
// 	reader, Error := zlib.NewWriter(&Buffer)
// 	defer reader.Close()
// 	Common.Check(Error)
// 	DecodeBytes, Error := ioutil.ReadAll(reader)
// 	Common.Check(Error)

// 	return hex.EncodeToString(DecodeBytes)
// }

// 轉換牌組資料
func DecodeDataToMap(DataStr string) map[string]interface{} {
	// 轉成 Byte Array
	HexData, Error := hex.DecodeString(DataStr)
	Common.Check(Error)

	// 接下來這段是根據 C# 裡面的 Script 去改寫
	// 會有一些前置作業去判斷是否有效
	// 接著才會去讀取資料
	//
	// http://docs2x.smartfoxserver.com/DevelopmentBasics/sfsobject-sfsarray
	// 前面的 Title
	DataWrapper, _ := ReadingSFSObject(HexData, 0)
	return DataWrapper.Data.(map[string]interface{})
}
