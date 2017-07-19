package main

import (
	"github.com/marshome/p-vision/services/vision"
	"fmt"
)

func main(){
	service:=vision.NewVisionService()
	err:=service.Annotate("./testdata/1.jpg")
	if err!=nil{
		fmt.Println(err)
	}
}
