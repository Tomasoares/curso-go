package api

type JsonConverter interface {
	convert(animals []Animal) []byte
}
