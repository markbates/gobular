package models

type Checker struct {
	Expression string
	TestString string
}

type Result struct {
	Num     int
	Line    string
	Matches []string
}
