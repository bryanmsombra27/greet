package main

import (
	"flag"
	"fmt"

	lgx "github.com/bryansombra27/greet/internal/logx"
	"github.com/bryansombra27/greet/pkg/greet"
	"github.com/fatih/color"
)

func main() {
	name:= flag.String("name","mundo","Nombre a saludar")

	flag.Parse()
	fmt.Println("hola mundo")
	msg:= greet.Hello(*name)
	
	fmt.Println(msg)
	lgx.Info("Keso!")
	color.Green("hola, %s saludos",*name)

}