package handlers

type TernaDto struct {
	Titulacion string `json:"titulacion"`
	Curso      int    `json:"curso"`
	Grupo      int    `json:"grupo"`
}

type EntryDTO struct{}

type ErrorHttp struct {
	Message string `json:"message"`
}
