package main

import (
	"fmt"
	"reflect"

	pb "course.proto.go.com/proto"
	"google.golang.org/protobuf/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          7,
		IsSimple:    true,
		Name:        "Junin",
		SampleLists: []int32{1, 2, 3, 4, 5, 6, 7},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{
			Id:   3,
			Name: "Three",
		},
		MultipleDummies: []*pb.Dummy{
			{
				Id:   34,
				Name: "Trinta e quatro",
			},
			{
				Id:   8,
				Name: "Eight",
			},
		},
	}
}

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_EYE_COLOR_BROWN,
	}
}

func doOneOf(message interface{}) {
	switch m := message.(type) {
	case *pb.OneOf_Id:
		fmt.Println(m.Id)
	case *pb.OneOf_Name:
		fmt.Println(m.Name)
	default:
		fmt.Println(fmt.Errorf("message has unexpected type: %v", m))
	}
}

func doMaps() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IDWrapper{
			"key1": {Id: 31},
			"key2": {Id: 32},
			"key3": {Id: 33},
		},
	}
}

func doMaps2() *pb.MapExample2 {
	return &pb.MapExample2{
		Ids: map[string]float32{
			"key1": 3.13,
			"key2": 8.29,
			"key3": 39.05,
		},
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"
	writeToFile(path, p)
	message := &pb.Simple{}
	readFromFile(path, message)
	fmt.Println(message)
}

func doToJSON(p proto.Message) string {
	jsonString := toJSON(p)
	return jsonString
}

func doFromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(jsonString, message)
	return message
}

func main() {
	//fmt.Println(doSimple())
	//fmt.Println(doComplex())
	//fmt.Println(doEnum())
	//doOneOf(&pb.OneOf_Id{Id: 99})
	//doOneOf(&pb.OneOf_Name{Name: "noventa e nove"})
	//fmt.Println(doMaps())
	//fmt.Println(doMaps2())
	//doFile(doSimple())
	// json simple
	jsonString := doToJSON(doSimple())
	message := doFromJSON(jsonString, reflect.TypeOf(pb.Simple{}))
	fmt.Println(jsonString)
	fmt.Println(message)

	// json complex
	jsonString = doToJSON(doComplex())
	message = doFromJSON(jsonString, reflect.TypeOf(pb.Complex{}))
	fmt.Println(jsonString)
	fmt.Println(message)

	// json unknown field - discardUnknown
	fmt.Println(doFromJSON(`{"id": 1, "otherField": true}`, reflect.TypeOf(pb.Simple{})))
}
