package test

import (
	"bytes"
	"fmt"
	sw "inventory-api-server/go"
	"io"
	"net/http"
	"strings"
	"testing"
)

func constructJsonArray(jsonObjects []string) string {

	array := "["
	for _, obj := range jsonObjects {
		array += obj + ","
	}

	// Replace the trailing , with a ]
	array = array[:len(array)-1] + "]"

	return array

}

func TestGetAll(t *testing.T, table string, expectedResponseObjects []string) {

	// Get all records //

	resp, err := http.Get(fmt.Sprintf(`%s/%s`, sw.GetApiServerAddress(), table))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Error("Expected status code", http.StatusOK, "but got", resp.StatusCode)
	}

	// Decode response //

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	stringBody := string(body)

	// Assert response //

	expectedResponse := constructJsonArray(expectedResponseObjects)

	if stringBody != expectedResponse {
		t.Error("Expected", expectedResponse, "but got", stringBody)
	}

}

type PostTestCase struct {
	postBody     string
	responseBody string
	statusCode   int
}

func TestPost(t *testing.T, table string, testCases []PostTestCase) {

	for _, testCase := range testCases {

		// Post //

		requestBody := bytes.NewBuffer([]byte(testCase.postBody))

		resp, err := http.Post(fmt.Sprintf(`%s/%s`, sw.GetApiServerAddress(), table), "application/json", requestBody)
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		// Assert //

		if testCase.statusCode != resp.StatusCode {
			t.Error("Expected status code", testCase.statusCode, "but got", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		// Remove newline from empty response
		trimmedBody := strings.TrimSpace(string(body))

		if !strings.Contains(trimmedBody, testCase.responseBody) {
			t.Error("Expected response body to be \"" + testCase.responseBody + "\" but got \"" + trimmedBody + "\"")
		}

	}

}

func TestGetById(t *testing.T, table string, id int, expectedResponseObject string) {

	// Get one record //

	resp, err := http.Get(fmt.Sprintf(`%s/%s/%v`, sw.GetApiServerAddress(), table, id))
	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Error("Expected status code", http.StatusOK, "but got", resp.StatusCode)
	}

	// Decode response //

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	stringBody := string(body)

	// Assert response //

	if stringBody != expectedResponseObject {
		t.Error("Expected response object", expectedResponseObject, "but got", stringBody)
	}

}

func TestPutById(t *testing.T, table string, id int, requestBody string) {

	// Update record //

	requestBodyBytes := bytes.NewBuffer([]byte(requestBody))

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf(`%s/%s/%v`, sw.GetApiServerAddress(), table, id), requestBodyBytes)
	if err != nil {
		t.Error(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Error("Expected status code", http.StatusOK, "but got", resp.StatusCode)
	}

}

func TestDeleteById(t *testing.T, table string, id int) {

	// Delete record //

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf(`%s/%s/%v`, sw.GetApiServerAddress(), table, id), nil)
	if err != nil {
		t.Error(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Error("Expected status code", http.StatusOK, "but got", resp.StatusCode)
	}

	// Assert gone //

	resp, err = http.Get(fmt.Sprintf("%s/%s/%v", sw.GetApiServerAddress(), table, id))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 404 {
		t.Error("Expected status code 404 but got", resp.StatusCode)
	}

}
