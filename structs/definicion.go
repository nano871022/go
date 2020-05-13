package structs

type Persona struct{
  Nombre string
  Edad int
}

type Estudiante struct{
  Persona
  Semestre int8
}
type Profesor struct{
  Estudiante
  Semestre int8
}
