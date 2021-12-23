package helper

import "strings"

//Response é usado para retorno de formato json estático
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

//EmptyObj é usado quando os dados não podem ser nulos no json
type EmptyObj struct{}

//BuildResponse é a função responsával por injetar valor de dados para resposta dinâmica de sucesso
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

//BuildErrorResponse função responsável por injetar o valor dos dados na resposta dinâmica com falha
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}