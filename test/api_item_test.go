package test

import (
	"testing"
)

func TestGetItems(t *testing.T) {

	tearDown := SetupTest(t, []TableSetup{
		{"category", []string{
			`{"name":"Grains"}`,
			`{"name":"Spices"}`,
		}},
		{"container", []string{
			``, ``, ``,
		}},
		{"item", []string{
			`{"categoryId":1,"containerId":1,"description":"5 Kilograms","id":1,"name":"Wild Rice"}`,
			`{"categoryId":2,"containerId":2,"description":"10 Kilograms","id":2,"name":"Garam Masala"}`,
			`{"categoryId":1,"containerId":3,"description":"5 Kilograms","id":3,"name":"Jasmin Rice"}`,
		}},
	})
	defer tearDown(t)

	TestGetAll(t, "item", []string{
		`{"categoryId":1,"containerId":1,"description":"5 Kilograms","id":1,"name":"Wild Rice"}`,
		`{"categoryId":2,"containerId":2,"description":"10 Kilograms","id":2,"name":"Garam Masala"}`,
		`{"categoryId":1,"containerId":3,"description":"5 Kilograms","id":3,"name":"Jasmin Rice"}`,
	})

}

func TestPostItem(t *testing.T) {

	SetupTest(t, []TableSetup{
		{"category", []string{
			`{"name":"Grains"}`,
			`{"name":"Spices"}`,
		}},
		{"container", []string{
			``, ``, ``,
		}},
	})

	TestPost(t, "item", []PostTestCase{
		{`{"name":"Wild Rice","description":"5 Kilograms","categoryId":1,"containerId":1}`, `{"itemId":1}`, 200},
		{`{"name":"Garam Masala","description":"10 Kilograms","categoryId":2,"containerId":2}`, `{"itemId":2}`, 200},
		{`{"name":}`, ``, 400},
	})

}

func TestGetItemById(t *testing.T) {

	tearDown := SetupTest(t, []TableSetup{
		{"category", []string{
			`{"name":"Grains"}`,
			`{"name":"Spices"}`,
		}},
		{"container", []string{
			``, ``, ``,
		}},
		{"item", []string{
			`{"categoryId":1,"containerId":1,"description":"5 Kilograms","id":1,"name":"Wild Rice"}`,
			`{"categoryId":2,"containerId":2,"description":"10 Kilograms","id":2,"name":"Garam Masala"}`,
			`{"categoryId":1,"containerId":3,"description":"5 Kilograms","id":3,"name":"Jasmin Rice"}`,
		}},
	})
	defer tearDown(t)

	TestGetById(t, "item", 2, `{"categoryId":2,"containerId":2,"description":"10 Kilograms","id":2,"name":"Garam Masala"}`)

}

func TestPutItemById(t *testing.T) {

	tearDown := SetupTest(t, []TableSetup{
		{"category", []string{
			`{"name":"Grains"}`,
			`{"name":"Spices"}`,
		}},
		{"container", []string{
			``, ``, ``,
		}},
		{"item", []string{
			`{"categoryId":1,"containerId":1,"description":"5 Kilograms","id":1,"name":"Wild Rice"}`,
			`{"categoryId":2,"containerId":2,"description":"10 Kilograms","id":2,"name":"Garam Masala"}`,
			`{"categoryId":1,"containerId":3,"description":"5 Kilograms","id":3,"name":"Jasmin Rice"}`,
		}},
	})
	defer tearDown(t)

	TestPutById(t, "item", 1, `{"name":"Oregano","description":"5 Kilograms","categoryId":2,"containerId":1}`)

}

func TestDeleteItemById(t *testing.T) {

	tearDown := SetupTest(t, []TableSetup{
		{"category", []string{
			`{"name":"Grains"}`,
			`{"name":"Spices"}`,
		}},
		{"container", []string{
			``, ``, ``,
		}},
		{"item", []string{
			`{"categoryId":1,"containerId":1,"description":"5 Kilograms","id":1,"name":"Wild Rice"}`,
			`{"categoryId":2,"containerId":2,"description":"10 Kilograms","id":2,"name":"Garam Masala"}`,
			`{"categoryId":1,"containerId":3,"description":"5 Kilograms","id":3,"name":"Jasmin Rice"}`,
		}},
	})
	defer tearDown(t)

	TestDeleteById(t, "item", 1)

}
