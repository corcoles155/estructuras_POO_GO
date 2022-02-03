package User

import "time"

type Usuario struct {
	Id int
	Nombre string
	FechaAlta time.Time
	Status bool
}

//(this *Usuario) --> Cuando se haga mención a la palabra this dentro de la función me estaré refiriendo a la estructura Usuario
func (this *Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, status bool)  {
	this.Id=id
	this.Nombre=nombre
	this.FechaAlta=fechaalta
	this.Status=status
}