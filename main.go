package main

import (
	"encoding/json"
	"fmt"

	"github.com/jncmaguire/example-proto-custom-resolver/resolver"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
	any := new(anypb.Any)

	err := json.Unmarshal([]byte(`{"type_url":"type.googleapis.com/pet.Pet","value":"Cg5QYXRjaHlTY3JhdGNoeRIGCPDg59cEGmUKVwoIQW5pbWFsaWESCENob3JkYXRhGghNYW1tYWxpYSIJQ2Fybml2b3JhKgpGZWxpZm9ybWlhMgdGZWxpZGFlOgdGZWxpbmFlQgVGZWxpc0oHRi5jYXR1cxIDY2F0EgVraXR0eSJjClwKCEFuaW1hbGlhEghDaG9yZGF0YRoITWFtbWFsaWEiCFByaW1hdGVzKgpIYXBsb3JoaW5pMglIb21pbmlkYWU6CUhvbWluaW5hZUIESG9tb0oKSC4gc2FwaWVucxIDbWFuKm4KB2NhdC1tb20SYwpcCghBbmltYWxpYRIIQ2hvcmRhdGEaCE1hbW1hbGlhIghQcmltYXRlcyoKSGFwbG9yaGluaTIJSG9taW5pZGFlOglIb21pbmluYWVCBEhvbW9KCkguIHNhcGllbnMSA21hbipyCglodW1hbi1tb20SZQpXCghBbmltYWxpYRIIQ2hvcmRhdGEaCE1hbW1hbGlhIglDYXJuaXZvcmEqCkZlbGlmb3JtaWEyB0ZlbGlkYWU6B0ZlbGluYWVCBUZlbGlzSgdGLmNhdHVzEgNjYXQSBWtpdHR5"}`), any)

	if err != nil {
		panic(err)
	}

	r, err := resolver.New()

	if err != nil {
		panic(err)
	}

	message, err := anypb.UnmarshalNew(any, proto.UnmarshalOptions{Resolver: r})
	if err != nil {
		panic(err)
	}

	fmt.Println(processMessageAsMap(message))
}
