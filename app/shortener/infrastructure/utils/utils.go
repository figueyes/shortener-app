package utils

import (
	"crypto/sha256"
	"encoding/json"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/log"
	"github.com/itchyny/base58-go"
	"reflect"
)

func CreateSha256(input string) []byte {
	sha := sha256.New()
	sha.Write([]byte(input))
	return sha.Sum(nil)
}

func CreateBase58(input []byte) string {
	e := base58.BitcoinEncoding
	encoded, err := e.Encode(input)
	if err != nil {
		log.Error(err.Error())
		panic(err.Error())
	}
	return string(encoded)
}

func BoolAddr(b bool) *bool {
	boolVar := b
	return &boolVar
}

func ConvertEntity(in, out interface{}) interface{} {
	str, _ := json.Marshal(in)
	err2 := json.Unmarshal(str, out)

	if err2 != nil {
		return nil
	}
	return out
}

func EntityToJson(entity interface{}) string {
	str, err := json.Marshal(entity)
	if err != nil {
		return "{}"
	}
	return string(str)
}

func JsonToEntity(jsonIn string, entity interface{}) {
	err := json.Unmarshal([]byte(jsonIn), entity)

	if err != nil {
		entity = nil
	}
}

func IsNilFixed(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
