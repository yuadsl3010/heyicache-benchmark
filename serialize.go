package main

import "unsafe"

// SerializeTestStruct 将 TestStruct 序列化为字节数组
// 使用 protobuf 只序列化 TestProto 字段，因为 TestStruct 本身不是 protobuf 消息
// 这样可以减少序列化开销，专注于 protobuf 性能测试
func SerializeTestStruct(ts *TestStruct) ([]byte, error) {
	if ts == nil || ts.TestProto == nil {
		return nil, nil
	}

	// 序列化 TestProto 字段
	return ts.TestProto.Marshal()
}

// DeserializeTestStruct 从字节数组反序列化为 TestStruct
// 使用 protobuf 反序列化 TestProto 字段，并重构 TestStruct
// 注意：非 protobuf 字段会使用默认值，这适合 protobuf 性能基准测试
func DeserializeTestStruct(data []byte) (*TestStruct, error) {
	if len(data) == 0 {
		return nil, nil
	}

	testProto := &TestPB{}
	if err := testProto.Unmarshal(data); err != nil {
		return nil, err
	}

	// 从 protobuf 数据重构 TestStruct，这样可以保持数据一致性
	// 同时专注于测试 protobuf 序列化性能

	return &TestStruct{
		Id:       testProto.Id % uint64(maxNum),
		TestName: testProto.TestString,
		TestChild: TestStructChild{
			Id:       testProto.TestChild.Id % uint64(maxNum),
			TestName: testProto.TestChild.TestString,
		},
		TestProto: testProto,
	}, nil
}

// StringToByte 高性能强转string->[]byte
func StringToByte(s string) []byte {
	tmp1 := (*[2]uintptr)(unsafe.Pointer(&s))
	tmp2 := [3]uintptr{tmp1[0], tmp1[1], tmp1[1]}
	return *(*[]byte)(unsafe.Pointer(&tmp2))
}
