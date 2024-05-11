package test

import (
	"testing"
)

func TestGetItems(t *testing.T) {

	tearDownTest := SetupPopulatedTest(t, "item", []string{
		`{"name":"Wild Rice","description":"5 Kilograms","categoryId":0,"containerId":0}`,
		`{"name":"Garam Masala","description":"10 Kilograms","categoryId":1,"containerId":1}`,
		`{"name":"Jasmin Rice","description":"5 Kilograms","categoryId":0,"containerId":2}`,
	})
	defer tearDownTest(t)

	TestGetAll(t, "item", []string{
		`{"categoryId":0,"containerId":0,"description":"5 Kilograms","id":1,"name":"Wild Rice"}`,
		`{"categoryId":1,"containerId":1,"description":"10 Kilograms","id":2,"name":"Garam Masala"}`,
		`{"categoryId":0,"containerId":2,"description":"5 Kilograms","id":3,"name":"Jasmin Rice"}`,
	})

}

func TestPostItem(t *testing.T) {

	tearDownTest := SetupEmptyTest(t, "item")
	defer tearDownTest(t)

	TestPost(t, "item", []PostTestCase{
		{`{"name":"Wild Rice","description":"5 Kilograms","categoryId":0,"containerId":0}`, `{"itemId":1}`, 200},
		{`{"name":"Garam Masala","description":"10 Kilograms","categoryId":1,"containerId":1}`, `{"itemId":2}`, 200},
		{`{"name":}`, ``, 400},
	})

}

func TestGetItemById(t *testing.T) {

	tearDownTest := SetupPopulatedTest(t, "item", []string{
		`{"name":"Wild Rice","description":"5 Kilograms","categoryId":0,"containerId":0}`,
		`{"name":"Garam Masala","description":"10 Kilograms","categoryId":1,"containerId":1}`,
		`{"name":"Jasmin Rice","description":"5 Kilograms","categoryId":0,"containerId":2}`,
	})
	defer tearDownTest(t)

	TestGetById(t, "item", 2, `{"categoryId":1,"containerId":1,"description":"10 Kilograms","id":2,"name":"Garam Masala"}`)

}

func TestPutItemById(t *testing.T) {

	tearDownTest := SetupPopulatedTest(t, "item", []string{
		`{"name":"Wild Rice","description":"5 Kilograms","categoryId":0,"containerId":0}`,
		`{"name":"Garam Masala","description":"10 Kilograms","categoryId":1,"containerId":1}`,
		`{"name":"Jasmin Rice","description":"5 Kilograms","categoryId":0,"containerId":2}`,
	})
	defer tearDownTest(t)

	TestPutById(t, "item", 1, `{"name":"Garam Masala","description":"10 Kilograms","categoryId":1,"containerId":1}`)

}

func TestDeleteItemById(t *testing.T) {

	tearDownTest := SetupPopulatedTest(t, "item", []string{
		`{"name":"Wild Rice","description":"5 Kilograms","categoryId":0,"containerId":0}`,
		`{"name":"Garam Masala","description":"10 Kilograms","categoryId":1,"containerId":1}`,
		`{"name":"Jasmin Rice","description":"5 Kilograms","categoryId":0,"containerId":2}`,
	})
	defer tearDownTest(t)

	TestDeleteById(t, "item", 1)

}
