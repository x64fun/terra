package main

// import (
// 	"encoding/json"
// 	"log"

// 	"github.com/google/uuid"
// 	"github.com/x64fun/terra/example/pb"
// 	"github.com/x64fun/terra/internal/gotype"
// 	"google.golang.org/protobuf/proto"
// 	"google.golang.org/protobuf/types/known/structpb"
// )

// func init() {
// 	log.SetFlags(log.Lshortfile)
// }
// func main() {
// 	func() {
// 		s := &pb.Setting{
// 			ID: &gotype.UUID{Data: []byte(uuid.New().String())},
// 		}
// 		b, err := proto.Marshal(s)
// 		if err != nil {
// 			panic(err)
// 		}
// 		log.Println(string(b))
// 		s2 := &pb.Setting{}
// 		if err := proto.Unmarshal(b, s2); err != nil {
// 			panic(err)
// 		}
// 		log.Println(s2)
// 	}()

// 	r := &pb.Request{
// 		Name: "x64fun",
// 	}
// 	var err error
// 	r.Detail, err = structpb.NewStruct(map[string]interface{}{
// 		"firstName": "John",
// 		"lastName":  "Smith",
// 		"isAlive":   true,
// 		"age":       27,
// 		"address": map[string]interface{}{
// 			"streetAddress": "21 2nd Street",
// 			"city":          "New York",
// 			"state":         "NY",
// 			"postalCode":    "10021-3100",
// 		},
// 		"phoneNumbers": []interface{}{
// 			map[string]interface{}{
// 				"type":   "home",
// 				"number": "212 555-1234",
// 			},
// 			map[string]interface{}{
// 				"type":   "office",
// 				"number": "646 555-4567",
// 			},
// 		},
// 		"children": []interface{}{},
// 		"spouse":   nil,
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	buf, err := json.Marshal(r)
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Println(string(buf))

// 	r2 := &pb.Request{}
// 	jsonBuf := []byte(`{
// 		"name": "thelark",
// 		"detail": {
// 			"hello": "world"
// 		}
// 	}`)
// 	if err := json.Unmarshal(jsonBuf, r2); err != nil {
// 		panic(err)
// 	}
// 	log.Println(r2.GetDetail().AsMap())
// 	r2.GetDetail().MarshalJSON()
// }
