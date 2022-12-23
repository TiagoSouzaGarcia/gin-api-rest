package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TiagoSouzaGarcia/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
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
	fmt.Println(string(respostaBody))
	fmt.Println(mockDaResposta)
}
