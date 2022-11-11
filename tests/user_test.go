package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func UserTest(server *gin.Engine, t *testing.T) {
	fmt.Println("============ USER Start ============")
	t.Run("Create User", func(t *testing.T) {
		jsonStr := []byte(`{"email":"test@gmail.com","password":"123456789","name":"test"}`)
		req, _ := http.NewRequest("POST", "/v1/user/create", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)
		expectedStatus := http.StatusOK
		assert.Equal(t, expectedStatus, w.Code)
	})
	t.Run("User signin", func(t *testing.T) {
		jsonStr := []byte(`{"email":"test@gmail.com","password":"123456789"}`)
		req, _ := http.NewRequest("POST", "/v1/user/signin", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)
		fmt.Println("f: ", w.Body.String())
		gToken := gjson.Get(w.Body.String(), "data.token")
		fmt.Println("token: ", gToken.String())
		expectedStatus := http.StatusOK
		assert.Equal(t, expectedStatus, w.Code)
	})
	fmt.Println("============ USER End ============")
}
