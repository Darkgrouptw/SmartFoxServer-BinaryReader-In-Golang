package main

import "./SFSReader"

func main() {
	TestData := "12000100024473110002120006000a556e6974734c6576656c0a00000008010304010101010100094465636b496e64657804000000000006426f73734944040000b5a50009426c6f674c6576656c040000000600084465707574794944040000b5a60005556e6974730c000800002af900002aff00002ee100002ee20000520e0000791b00007d01000080ea120006000a556e6974734c6576656c0a00000008010401010101010100094465636b496e64657804000000030006426f73734944040000b5a50009426c6f674c6576656c040000000600084465707574794944040000b5a60005556e6974730c000800002af900002ee100002ee2000032ca0000520b0000520c0000791a00007924"
	// TestData := "12000100024473110002120006000a556e6974734c6576656c0a00000008141414141414141400094465636b496e64657804000000000006426f73734944040000b5a50009426c6f674c6576656c040000000500084465707574794944040000b5a60005556e6974730c000800002af900002afe0000520b0000520d0000791a0000792400007d01000080ea120006000a556e6974734c6576656c0a00000008141414141414141400094465636b496e64657804000000000006426f73734944040000b5a50009426c6f674c6576656c040000000500084465707574794944040000b5a60005556e6974730c000800002af900002afe0000520b0000520d0000791a0000792400007d01000080ea"
	// TestData := "12000100024473110002120006000a556e6974734c6576656c0a00000008010101010101010800094465636b496e64657804000000010006426f73734944040000b5a50009426c6f674c6576656c040000000400084465707574794944040000b5a60005556e6974730c000800002af900002aff00002ee100002ee200002ee30000520a0000791a00007d01120006000a556e6974734c6576656c0a00000008141414141414141400094465636b496e64657804000000010006426f73734944040000b5a50009426c6f674c6576656c040000000100084465707574794944040000b5a60005556e6974730c000800002af900002aff00002ee100002ee20000520a0000520e00007d01000080ea"
	print(SFSReader.DecodeDataToMap(TestData))
}
