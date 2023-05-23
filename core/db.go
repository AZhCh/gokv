package core

type DB struct {
	dict       *Dict
	expireDict *Dict
}

type Dict struct {
	data map[string]DictEntry
}

type DictEntry struct {
	key       string
	value     any
	entryType uint8
}
