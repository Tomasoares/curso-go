package api

type CsvReader interface {
	ReadFile(filename string) []Animal
}
