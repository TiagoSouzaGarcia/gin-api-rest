package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestBuscaAlunoPorIDHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscarAlunoPorId)
	//Descricao da rota. Conversao do ID de int para string com Itoa.
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome)
	assert.Equal(t, "12345678901", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletarAluno)
	pathDeBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDeBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditaUmALunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	//Patch edita todos os campos
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	//Objeto que será passado no método.
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "47123456789", RG: "123456700"}
	//Marshal converte o vaor para json
	valorJson, _ := json.Marshal(aluno)
	pathParaEditar := "/alunos/" + strconv.Itoa(ID)
	//Dessa vez será passada algo para o corpo da requisição. A json será passada para bytes.
	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMockAtualizado models.Aluno
	//Unmarshal vai transformar os valores em bytes da resposta para um Json e armazenar no endereco de memória determinado.
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, "47123456789", alunoMockAtualizado.CPF)
	assert.Equal(t, "Nome do Aluno Teste", alunoMockAtualizado.Nome)
	assert.Equal(t, "123456700", alunoMockAtualizado.RG)
}
