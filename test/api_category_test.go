package test

import "testing"

func TestGetCategories(t *testing.T) {

	tearDown := SetupTest(t, []TableSetup{
		{"category", []string{
			`{"name":"Grains"}`,
			`{"name":"Spices"}`,
			`{"name":"Seeds"}`,
		}},
	})
	defer tearDown(t)

	TestGetAll(t, "category", []string{
		`{"id":1,"name":"Grains"}`,
		`{"id":2,"name":"Spices"}`,
		`{"id":3,"name":"Seeds"}`,
	})

}

func TestPostCategory(t *testing.T) {

	tearDown := SetupTest(t, []TableSetup{})
	defer tearDown(t)

	TestPost(t, "category", []PostTestCase{
		{`{"name":"Grains"}`, `{"categoryId":1}`, 200},
		{`{"name":"Spices"}`, `{"categoryId":2}`, 200},
		{`{"name":"Seeds"}`, `{"categoryId":3}`, 200},
		{`{"name":}`, ``, 400},
	})

}

func TestGetCategoryById(t *testing.T) {

	tearDown := SetupTest(t, []TableSetup{
		{"category", []string{
			`{"name":"Grains"}`,
			`{"name":"Spices"}`,
			`{"name":"Seeds"}`,
		}},
	})
	defer tearDown(t)

	TestGetById(t, "category", 2, `{"id":2,"name":"Spices"}`)

}

func TestPutCategoryById(t *testing.T) {

	tearDown := SetupTest(t, []TableSetup{
		{"category", []string{
			`{"name":"Grains"}`,
			`{"name":"Spices"}`,
			`{"name":"Seeds"}`,
		}},
	})
	defer tearDown(t)

	TestPutById(t, "category", 1, `{"name":"Lentils"}`)

}

func TestDeleteCategoryById(t *testing.T) {

	tearDown := SetupTest(t, []TableSetup{
		{"category", []string{
			`{"name":"Grains"}`,
			`{"name":"Spices"}`,
			`{"name":"Seeds"}`,
		}},
	})
	defer tearDown(t)

	TestDeleteById(t, "category", 1)

}
