package main

type TestStruct struct {
	Id              uint64
	TestName        string
	TestSkip        string
	TestChild       TestStructChild
	TestChildren    []TestStructChild
	TestChildPtr    *TestStructChild
	TestChildrenPtr []*TestStructChild
	TestProto       *TestPB
	Flag            uint8
}

type TestStructChild struct {
	Id       uint64
	TestName string
	TestSkip string `heyicache:"skip"`
}
