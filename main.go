package main

import (
	"fmt"
	"ReadConfig/ReadUNIX"
)

func main(){
	configPath := "ReadUNIX/config.txt"
	//ReadUNIX
	myConfig := new(ReadUNIX.Config)
	myConfig.InitConfig(configPath)
	fmt.Println(myConfig.Read("num"))
}