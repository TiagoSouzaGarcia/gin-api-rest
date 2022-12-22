package main

import (
	"github.com/TiagoSouzaGarcia/api-go-gin/database"
	"github.com/TiagoSouzaGarcia/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	/* Mock de Alunos
	models.Alunos = []models.Aluno{
		{Nome: "Gui Lima", CPF: "00000000000", RG: "4700000000"},
		{Nome: "Ana", CPF: "11111111111", RG: "4800000000"},
	} */
	routes.HandleRequests()
}
