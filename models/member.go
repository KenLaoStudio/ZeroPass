package models

type Member struct {
	Address    string
	Identity_1 []byte
	Identity_2 []byte
	Licence    []byte
	Verify     bool
}
