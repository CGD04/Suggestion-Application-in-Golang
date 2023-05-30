package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdmin(t *testing.T){
	url := "http://localhost:8080/login"
	var data = []byte(`{"email":"john@gmail.com","password":"#@Furpa77"}`)
	//create request object
	r,_ :=http.NewRequest("POST",url,bytes.NewBuffer(data))
	r.Header.Set("Content-type","application/json")
	// create client
	client := &http.Client{}
	//send post request, specified by the method inside r
	resp,err := client.Do(r)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	body,_ := io.ReadAll(resp.Body)

	assert.Equal(t,http.StatusOK, resp.StatusCode)
	expResp := `{"message":"successful"}`
	assert.JSONEq(t,expResp,string(body))
}

