package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TiagoSouzaGarcia/api-go-gin/controllers"
	"github.com/TiagoSouzaGarcia/api-go-gin/database"
	"github.com/TiagoSouzaGarcia/api-go-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	//Enxugando as msgs.
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	//Requisicao que será passada
	req, _ := http.NewRequest("GET", "/Tiago", nil)
	//Resposta que será gravada com NewRecorder(). implementa a interface de response writer.
	resposta := httptest.NewRecorder()
	//Realiza a requisição. Tem como parametros a resposta da requisição e a requisição
	r.ServeHTTP(resposta, req)

	//Validacao do resultado
	/* if resposta.Code != http.StatusOK {
		t.Fatalf("Status error: valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK)
	} */

	//Verificando o status code da resposta
	assert.Equal(t, http.StatusOK, resposta.Code, "Status error: valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK)
	//Verificando o conteúdo da resposta
	mockDaResposta := `{"API diz:":"E ai Tiago, tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
	//impressão dos conteudos dos testes
	/* fmt.Println(string(respostaBody))
	fmt.Println(mockDaResposta) */
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibirTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
	//Impressão dos alunos
	fmt.Println(resposta.Body)
}

func TestBuscandoPorCPF(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("alunos/cpf/:cpf", controllers.BuscarAlunoPorCpf)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
