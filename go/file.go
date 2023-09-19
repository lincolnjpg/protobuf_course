package main

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

func writeToFile(fileName string, pb proto.Message) {
	if out, err := proto.Marshal(pb); err != nil {
		log.Fatalln("Deu ruim pra fazer o marshal", err)
		return
	} else if err := os.WriteFile(fileName, out, 0644); err != nil {
		log.Fatalln("Deu ruim pra escrever o arquivo no disco", err)
		return
	}

	fmt.Println("Deu bom pra escrever o arquivo no disco")
}

func readFromFile(fileName string, pb proto.Message) {
	in, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Deu ruim pra ler o arquivo no disco", err)
		return
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Deu ruim pra fazer o unmarshal da mensagem", err)
		return
	}

	fmt.Println(pb)
}
