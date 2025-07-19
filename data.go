package main

import (
	"fmt"
)

func GetKey(num int) string {
	// 创建一个唯一的键
	return fmt.Sprintf("test_key_%d", num)
}

// NewTestStruct 创建一个填充了数据的 TestStruct 指针
// 参数 num 用于生成不同的测试数据
func NewTestStruct(num int) (string, *TestStruct) {
	// 创建 TestPBChild 实例
	testPBChild := &TestPBChild{
		Id:         uint64(num + 1000),
		TestString: fmt.Sprintf("pb_child_string_%d", num),
		TestStrings: []string{
			fmt.Sprintf("pb_child_str1_%d", num),
			fmt.Sprintf("pb_child_str2_%d", num),
			fmt.Sprintf("pb_child_str3_%d", num),
		},
		TestMap: map[string]string{
			fmt.Sprintf("key1_%d", num): fmt.Sprintf("value1_%d", num),
			fmt.Sprintf("key2_%d", num): fmt.Sprintf("value2_%d", num),
		},
		TestUint64S: []uint64{
			uint64(num * 100),
			uint64(num * 200),
			uint64(num * 300),
		},
		TestBytes: []byte(fmt.Sprintf("bytes1_%d", num)),
		TestFloats: []float32{
			float32(num) * 1.1,
			float32(num) * 2.2,
			float32(num) * 3.3,
		},
	}

	// 创建多个 TestPBChild 实例用于数组
	testPBChildren := []*TestPBChild{
		{
			Id:          uint64(num + 2000),
			TestString:  fmt.Sprintf("pb_children_string1_%d", num),
			TestStrings: []string{fmt.Sprintf("pb_children_str_%d", num)},
			TestMap: map[string]string{
				fmt.Sprintf("child_key_%d", num): fmt.Sprintf("child_value_%d", num),
			},
			TestUint64S: []uint64{uint64(num * 400)},
			TestBytes:   []byte(fmt.Sprintf("child_bytes_%d", num)),
			TestFloats:  []float32{float32(num) * 4.4},
		},
		{
			Id:          uint64(num + 3000),
			TestString:  fmt.Sprintf("pb_children_string2_%d", num),
			TestStrings: []string{fmt.Sprintf("pb_children_str2_%d", num)},
			TestMap: map[string]string{
				fmt.Sprintf("child_key2_%d", num): fmt.Sprintf("child_value2_%d", num),
			},
			TestUint64S: []uint64{uint64(num * 500)},
			TestBytes:   []byte(fmt.Sprintf("child_bytes2_%d", num)),
			TestFloats:  []float32{float32(num) * 5.5},
		},
	}

	// 创建 TestPB 实例
	testPB := &TestPB{
		Id:         uint64(num + 10000),
		TestString: fmt.Sprintf("pb_string_%d", num),
		TestStrings: []string{
			fmt.Sprintf("pb_str1_%d", num),
			fmt.Sprintf("pb_str2_%d", num),
		},
		TestMap: map[string]string{
			fmt.Sprintf("pb_key1_%d", num): fmt.Sprintf("pb_value1_%d", num),
			fmt.Sprintf("pb_key2_%d", num): fmt.Sprintf("pb_value2_%d", num),
		},
		TestUint64S: []uint64{
			uint64(num * 1000),
			uint64(num * 2000),
		},
		TestBytes: []byte(fmt.Sprintf("pb_bytes1_%d", num)),
		TestFloats: []float32{
			float32(num) * 10.1,
			float32(num) * 20.2,
		},
		TestChild:    testPBChild,
		TestChildren: testPBChildren,
	}

	// 创建 TestStructChild 实例
	testStructChild := TestStructChild{
		Id:       uint64(num + 100),
		TestName: fmt.Sprintf("struct_child_name_%d_00000000", num),
		TestSkip: fmt.Sprintf("struct_child_skip_%d_00000000", num),
	}

	// 创建多个 TestStructChild 实例用于数组
	testStructChildren := []TestStructChild{
		{
			Id:       uint64(num + 200),
			TestName: fmt.Sprintf("struct_children_name1_%d", num),
			TestSkip: fmt.Sprintf("struct_children_skip1_%d", num),
		},
		{
			Id:       uint64(num + 300),
			TestName: fmt.Sprintf("struct_children_name2_%d", num),
			TestSkip: fmt.Sprintf("struct_children_skip2_%d", num),
		},
	}

	// 创建 TestStructChild 指针实例
	testStructChildPtr := &TestStructChild{
		Id:       uint64(num + 400),
		TestName: fmt.Sprintf("struct_child_ptr_name_%d", num),
		TestSkip: fmt.Sprintf("struct_child_ptr_skip_%d", num),
	}

	// 创建多个 TestStructChild 指针实例用于数组
	testStructChildrenPtr := []*TestStructChild{
		{
			Id:       uint64(num + 500),
			TestName: fmt.Sprintf("struct_children_ptr_name1_%d", num),
			TestSkip: fmt.Sprintf("struct_children_ptr_skip1_%d", num),
		},
		{
			Id:       uint64(num + 600),
			TestName: fmt.Sprintf("struct_children_ptr_name2_%d", num),
			TestSkip: fmt.Sprintf("struct_children_ptr_skip2_%d", num),
		},
	}

	// 创建并返回 TestStruct 实例
	return GetKey(num), &TestStruct{
		Id:              uint64(num),
		TestName:        fmt.Sprintf("test_name_%d", num),
		TestSkip:        fmt.Sprintf("test_skip_%d", num),
		TestChild:       testStructChild,
		TestChildren:    testStructChildren,
		TestChildPtr:    testStructChildPtr,
		TestChildrenPtr: testStructChildrenPtr,
		TestProto:       testPB,
	}
}

// freecache and bigcache test the protobuf field
func CheckTestStruct(num int, data *TestStruct, onlyCheckPB bool) bool {
	if data == nil {
		fmt.Println("CheckTestStruct: data is nil")
		return false
	}

	_, right := NewTestStruct(num)
	if !onlyCheckPB {
		// 检查基础字段
		if data.Id != right.Id {
			fmt.Printf("CheckTestStruct: data.Id mismatch, got %d, want %d\n", data.Id, right.Id)
			return false
		}

		if data.TestName != right.TestName {
			fmt.Printf("CheckTestStruct: data.TestName mismatch, got %s, want %s\n", data.TestName, right.TestName)
			return false
		}

		// 检查 TestChild
		if data.TestChild.Id != right.TestChild.Id {
			fmt.Printf("CheckTestStruct: data.TestChild.Id mismatch, got %d, want %d\n", data.TestChild.Id, right.TestChild.Id)
			return false
		}
		if data.TestChild.TestName != right.TestChild.TestName {
			fmt.Printf("CheckTestStruct: data.TestChild.TestName mismatch, got %s, want %s\n", data.TestChild.TestName, right.TestChild.TestName)
			return false
		}

		// 检查 TestChildren
		if len(data.TestChildren) != len(right.TestChildren) {
			fmt.Printf("CheckTestStruct: len(data.TestChildren) mismatch, got %d, want %d\n", len(data.TestChildren), len(right.TestChildren))
			return false
		}
		for i, expected := range right.TestChildren {
			if data.TestChildren[i].Id != expected.Id ||
				data.TestChildren[i].TestName != expected.TestName {
				fmt.Printf("CheckTestStruct: data.TestChildren[%d] mismatch, got (Id=%d, Name=%s), want (Id=%d, Name=%s)\n",
					i, data.TestChildren[i].Id, data.TestChildren[i].TestName, expected.Id, expected.TestName)
				return false
			}
		}

		// 检查 TestChildPtr
		if data.TestChildPtr == nil {
			fmt.Println("CheckTestStruct: data.TestChildPtr is nil")
			return false
		}
		if data.TestChildPtr.Id != right.TestChildPtr.Id {
			fmt.Printf("CheckTestStruct: data.TestChildPtr.Id mismatch, got %d, want %d\n", data.TestChildPtr.Id, right.TestChildPtr.Id)
			return false
		}
		if data.TestChildPtr.TestName != right.TestChildPtr.TestName {
			fmt.Printf("CheckTestStruct: data.TestChildPtr.TestName mismatch, got %s, want %s\n", data.TestChildPtr.TestName, right.TestChildPtr.TestName)
			return false
		}

		// 检查 TestChildrenPtr
		if len(data.TestChildrenPtr) != len(right.TestChildrenPtr) {
			fmt.Printf("CheckTestStruct: len(data.TestChildrenPtr) mismatch, got %d, want %d\n", len(data.TestChildrenPtr), len(right.TestChildrenPtr))
			return false
		}
		for i, expected := range right.TestChildrenPtr {
			if data.TestChildrenPtr[i] == nil {
				fmt.Printf("CheckTestStruct: data.TestChildrenPtr[%d] is nil\n", i)
				return false
			}
			if data.TestChildrenPtr[i].Id != expected.Id ||
				data.TestChildrenPtr[i].TestName != expected.TestName {
				fmt.Printf("CheckTestStruct: data.TestChildrenPtr[%d] mismatch, got (Id=%d, Name=%s), want (Id=%d, Name=%s)\n",
					i, data.TestChildrenPtr[i].Id, data.TestChildrenPtr[i].TestName, expected.Id, expected.TestName)
				return false
			}
		}
	}

	// 检查 TestProto
	if data.TestProto == nil {
		fmt.Println("CheckTestStruct: data.TestProto is nil")
		return false
	}

	// 检查 TestProto 基础字段
	if data.TestProto.Id != right.TestProto.Id {
		fmt.Printf("CheckTestStruct: data.TestProto.Id mismatch, got %d, want %d\n", data.TestProto.Id, right.TestProto.Id)
		return false
	}
	if data.TestProto.TestString != right.TestProto.TestString {
		fmt.Printf("CheckTestStruct: data.TestProto.TestString mismatch, got %s, want %s\n", data.TestProto.TestString, right.TestProto.TestString)
		return false
	}

	// 检查 TestProto.TestStrings
	if len(data.TestProto.TestStrings) != len(right.TestProto.TestStrings) {
		fmt.Printf("CheckTestStruct: len(data.TestProto.TestStrings) mismatch, got %d, want %d\n", len(data.TestProto.TestStrings), len(right.TestProto.TestStrings))
		return false
	}
	for i, expected := range right.TestProto.TestStrings {
		if data.TestProto.TestStrings[i] != expected {
			fmt.Printf("CheckTestStruct: data.TestProto.TestStrings[%d] mismatch, got %s, want %s\n", i, data.TestProto.TestStrings[i], expected)
			return false
		}
	}

	// // 检查 TestProto.TestMap
	// if len(data.TestProto.TestMap) != 2 {
	// 	fmt.Printf("CheckTestStruct: len(data.TestProto.TestMap) mismatch, got %d, want 2\n", len(data.TestProto.TestMap))
	// 	return false
	// }
	// expectedMap := map[string]string{
	// 	fmt.Sprintf("pb_key1_%d", num): fmt.Sprintf("pb_value1_%d", num),
	// 	fmt.Sprintf("pb_key2_%d", num): fmt.Sprintf("pb_value2_%d", num),
	// }
	// for key, expectedValue := range expectedMap {
	// 	if actualValue, exists := data.TestProto.TestMap[key]; !exists || actualValue != expectedValue {
	// 		fmt.Printf("CheckTestStruct: data.TestProto.TestMap[%s] mismatch, got %s, want %s\n", key, actualValue, expectedValue)
	// 		return false
	// 	}
	// }

	// 检查 TestProto.TestUint64S
	if len(data.TestProto.TestUint64S) != len(right.TestProto.TestUint64S) {
		fmt.Printf("CheckTestStruct: len(data.TestProto.TestUint64S) mismatch, got %d, want %d\n", len(data.TestProto.TestUint64S), len(right.TestProto.TestUint64S))
		return false
	}
	for i, expected := range right.TestProto.TestUint64S {
		if data.TestProto.TestUint64S[i] != expected {
			fmt.Printf("CheckTestStruct: data.TestProto.TestUint64S[%d] mismatch, got %d, want %d\n", i, data.TestProto.TestUint64S[i], expected)
			return false
		}
	}

	// 检查 TestProto.TestBytes
	for i, expected := range right.TestProto.TestBytes {
		if data.TestProto.TestBytes[i] != expected {
			fmt.Printf("CheckTestStruct: data.TestProto.TestBytes[%d] mismatch, got %v, want %v\n", i, data.TestProto.TestBytes[i], expected)
			return false
		}
	}

	// 检查 TestProto.TestFloats
	if len(data.TestProto.TestFloats) != len(right.TestProto.TestFloats) {
		fmt.Printf("CheckTestStruct: len(data.TestProto.TestFloats) mismatch, got %d, want %d\n", len(data.TestProto.TestFloats), len(right.TestProto.TestFloats))
		return false
	}
	for i, expected := range right.TestProto.TestFloats {
		if data.TestProto.TestFloats[i] != expected {
			fmt.Printf("CheckTestStruct: data.TestProto.TestFloats[%d] mismatch, got %f, want %f\n", i, data.TestProto.TestFloats[i], expected)
			return false
		}
	}

	// 检查 TestProto.TestChild
	if data.TestProto.TestChild == nil {
		fmt.Println("CheckTestStruct: data.TestProto.TestChild is nil")
		return false
	}
	if !checkTestPBChild(data.TestProto.TestChild, right.TestProto.TestChild) {
		fmt.Println("CheckTestStruct: data.TestProto.TestChild checkTestPBChild failed")
		return false
	}

	// 检查 TestProto.TestChildren
	if len(data.TestProto.TestChildren) != len(right.TestProto.TestChildren) {
		fmt.Printf("CheckTestStruct: len(data.TestProto.TestChildren) mismatch, got %d, want %d\n", len(data.TestProto.TestChildren), len(right.TestProto.TestChildren))
		return false
	}
	if !checkTestPBChild(data.TestProto.TestChildren[0], right.TestProto.TestChildren[0]) {
		fmt.Println("CheckTestStruct: data.TestProto.TestChildren[0] checkTestPBChild failed")
		return false
	}
	if !checkTestPBChild(data.TestProto.TestChildren[1], right.TestProto.TestChildren[1]) {
		fmt.Println("CheckTestStruct: data.TestProto.TestChildren[1] checkTestPBChild failed")
		return false
	}

	return true
}

// checkTestPBChild 辅助函数，用于检查 TestPBChild 结构
func checkTestPBChild(child *TestPBChild, right *TestPBChild) bool {
	if child == nil {
		fmt.Println("checkTestPBChild: child is nil")
		return false
	}

	if child.Id != right.Id {
		fmt.Printf("checkTestPBChild: Id mismatch, got %d, want %d\n", child.Id, right.Id)
		return false
	}

	if child.TestString != right.TestString {
		fmt.Printf("checkTestPBChild: TestString mismatch, got %s, want %s\n", child.TestString, right.TestString)
		return false
	}
	if len(child.TestStrings) != len(right.TestStrings) {
		fmt.Printf("checkTestPBChild: len(TestStrings) mismatch, got %d, want %d\n", len(child.TestStrings), len(right.TestStrings))
		return false
	}
	for i, expected := range right.TestStrings {
		if child.TestStrings[i] != expected {
			fmt.Printf("checkTestPBChild: TestStrings[%d] mismatch, got %s, want %s\n", i, child.TestStrings[i], expected)
			return false
		}
	}

	if len(child.TestUint64S) != len(right.TestUint64S) {
		fmt.Printf("checkTestPBChild: len(TestUint64S) mismatch, got %d, want %d\n", len(child.TestUint64S), len(right.TestUint64S))
		return false
	}
	for i, expected := range right.TestUint64S {
		if child.TestUint64S[i] != expected {
			fmt.Printf("checkTestPBChild: TestUint64S[%d] mismatch, got %d, want %d\n", i, child.TestUint64S[i], expected)
			return false
		}
	}

	for i, expected := range right.TestBytes {
		if child.TestBytes[i] != expected {
			fmt.Printf("checkTestPBChild: TestBytes[%d] mismatch, got %v, want %v\n", i, child.TestBytes[i], expected)
			return false
		}
	}

	if len(child.TestFloats) != len(right.TestFloats) {
		fmt.Printf("checkTestPBChild: len(TestFloats) mismatch, got %d, want %d\n", len(child.TestFloats), len(right.TestFloats))
		return false
	}
	for i, expected := range right.TestFloats {
		if child.TestFloats[i] != expected {
			fmt.Printf("checkTestPBChild: TestFloats[%d] mismatch, got %f, want %f\n", i, child.TestFloats[i], expected)
			return false
		}
	}

	return true
}
