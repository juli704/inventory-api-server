package test

import (
	"testing"
)

func TestGetCategories(t *testing.T) {

	tearDownTest := SetupPopulatedTest(t, "category", []string{
		`{"name":"Grains"}`,
		`{"name":"Spices"}`,
		`{"name":"Grains"}`,
	})
	defer tearDownTest(t)

	TestGetAll(t, "category", []string{
		`{"id":1,"name":"Grains"}`,
		`{"id":2,"name":"Spices"}`,
		`{"id":3,"name":"Grains"}`,
	})

}

func TestPostCategory(t *testing.T) {

	tearDownTest := SetupEmptyTest(t, "category")
	defer tearDownTest(t)

	TestPost(t, "category", []PostTestCase{
		{`{"name":"Grains"}`, `{"categoryId":1}`, 200},
		{`{"name":"Spices"}`, `{"categoryId":2}`, 200},
		{`{"name":"Grains"}`, `{"categoryId":3}`, 200},
		{`{"name":}`, ``, 400},
	})

}

func TestGetCategoryById(t *testing.T) {

	tearDownTest := SetupPopulatedTest(t, "category", []string{
		`{"name":"Grains"}`,
		`{"name":"Spices"}`,
		`{"name":"Grains"}`,
	})
	defer tearDownTest(t)

	TestGetById(t, "category", 2, `{"id":2,"name":"Spices"}`)

}

func TestPutCategoryById(t *testing.T) {

	tearDownTest := SetupPopulatedTest(t, "category", []string{
		`{"name":"Grains"}`,
		`{"name":"Spices"}`,
		`{"name":"Grains"}`,
	})
	defer tearDownTest(t)

	TestPutById(t, "category", 1, `{"name":"Lentils"}`)

}

func TestDeleteCategoryById(t *testing.T) {

	tearDownTest := SetupPopulatedTest(t, "category", []string{
		`{"name":"Grains"}`,
		`{"name":"Spices"}`,
		`{"name":"Grains"}`,
	})
	defer tearDownTest(t)

	TestDeleteById(t, "category", 1)

}
