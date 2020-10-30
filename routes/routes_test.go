package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cavdy-play/go_mongo/config"
	"github.com/cavdy-play/go_mongo/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type Todos struct {
	Message string             `json:"string"`
	Status  string             `json:"string"`
	Todos   []controllers.Todo `json:"data"`
}

type header struct {
	Key   string
	Value string
}

func performRequest(r http.Handler, method, path string, headers ...header) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func performPost(r http.Handler, method, path string, buffer *bytes.Buffer, headers ...header) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, buffer)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func setup() *gin.Engine {
	config.Connect()
	router := gin.Default()
	Routes(router)
	return router
}

func TestRouteWelcome(t *testing.T) {
	message := ""
	r := setup()

	var response map[string]string
	w := performRequest(r, "GET", "/")

	json.Unmarshal([]byte(w.Body.String()), &response)
	message = response["message"]

	assert.Equal(t, "Welcome To API", message)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRouteNotFound(t *testing.T) {
	message := ""
	r := setup()

	var response map[string]string
	w := performRequest(r, "GET", "/notFound")

	json.Unmarshal([]byte(w.Body.String()), &response)
	message = response["message"]

	assert.Equal(t, "Route Not Found", message)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestRouteTodos(t *testing.T) {
	message := ""
	r := setup()

	var response map[string]string
	w := performRequest(r, "GET", "/todos")

	json.Unmarshal([]byte(w.Body.String()), &response)
	message = response["message"]
	_, hasData := response["data"]

	if w.Code == http.StatusOK {
		assert.Equal(t, "All Todos", message)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, true, hasData)
	} else {
		assert.Equal(t, "Something went wrong", message)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, false, hasData)
	}
}

func TestRouteInsertTodo(t *testing.T) {
	message := ""
	r := setup()

	var response map[string]string

	var jsonStr = []byte(`{ "Title": "TEST", "Body": "Body Test", "Completed": "Complete"}`)

	w := performPost(r, "POST", "/todo", bytes.NewBuffer(jsonStr), header{
		Key: "Content-Type", Value: "application/json",
	})

	json.Unmarshal([]byte(w.Body.String()), &response)
	message = response["message"]

	if w.Code == http.StatusCreated {
		assert.Equal(t, "Todo created Successfully", message)
		assert.Equal(t, http.StatusCreated, w.Code)
	} else {
		assert.Equal(t, "Something went wrong", message)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	}
}

func TestRouteGetTodoById(t *testing.T) {
	message := ""
	r := setup()

	var responseTodos Todos
	todo := controllers.Todo{}
	q := performRequest(r, "GET", "/todos")
	json.Unmarshal([]byte(q.Body.String()), &responseTodos)

	if q.Code == http.StatusOK {
		todo = responseTodos.Todos[0]
		url := fmt.Sprintf("/todo/%v", todo.ID)

		t.Logf("URL : %v", url)

		var response map[string]string
		w := performRequest(r, "GET", url)

		json.Unmarshal([]byte(w.Body.String()), &response)
		message = response["message"]
		_, hasData := response["data"]

		if w.Code == http.StatusOK {
			assert.Equal(t, "Single Todo", message)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, true, hasData)
		} else {
			assert.Equal(t, "Todo not found", message)
			assert.Equal(t, http.StatusNotFound, w.Code)
			assert.Equal(t, false, hasData)
		}
	} else {
		t.Error("Aucune donnée pour tester")
	}
}

func TestRouteUpdateTodo(t *testing.T) {
	message := ""
	r := setup()

	var responseTodos Todos
	todo := controllers.Todo{}
	q := performRequest(r, "GET", "/todos")
	json.Unmarshal([]byte(q.Body.String()), &responseTodos)

	if q.Code == http.StatusOK {
		todo = responseTodos.Todos[0]
		url := fmt.Sprintf("/todo/%v", todo.ID)

		t.Logf("URL : %v", url)

		var response map[string]string
		var jsonStr = []byte(`{"Completed": "Not Complete or maybe"}`)
		w := performPost(r, "PUT", url, bytes.NewBuffer(jsonStr), header{
			Key: "Content-Type", Value: "application/json",
		})

		json.Unmarshal([]byte(w.Body.String()), &response)
		message = response["message"]

		if w.Code == http.StatusOK {
			assert.Equal(t, "Todo Edited Successfully", message)
			assert.Equal(t, http.StatusOK, w.Code)
		} else {
			assert.Equal(t, "Something went wrong", message)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		}
	} else {
		t.Error("Aucune donnée pour tester")
	}
}

func TestRouteDeleteTodo(t *testing.T) {
	message := ""
	r := setup()

	var responseTodos Todos
	todo := controllers.Todo{}
	q := performRequest(r, "GET", "/todos")
	json.Unmarshal([]byte(q.Body.String()), &responseTodos)

	if q.Code == http.StatusOK {
		todo = responseTodos.Todos[0]
		url := fmt.Sprintf("/todo/%v", todo.ID)

		t.Logf("URL : %v", url)

		var response map[string]string
		w := performRequest(r, "DELETE", url)

		json.Unmarshal([]byte(w.Body.String()), &response)
		message = response["message"]

		if w.Code == http.StatusOK {
			assert.Equal(t, "Todo deleted successfully", message)
			assert.Equal(t, http.StatusOK, w.Code)
		} else {
			assert.Equal(t, "Something went wrong", message)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		}
	} else {
		t.Error("Aucune donnée pour tester")
	}
}
