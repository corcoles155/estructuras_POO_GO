package main

import (
	"fmt"
	"time"

	us "./user"
)	

type empleado struct {
	us.Usuario //empleado hereda de Usuario
}

func main()  {
	pablo := new(empleado)
	pablo.AltaUsuario(1, "Pablo", time.Now(), true)

	fmt.Println(pablo.Usuario)
}	
