package test

import "testing"

func TestGetContainers(t *testing.T) {

	tearDown := SetupTest(t, []TableSetup{
		{"container", []string{``, ``, ``}},
	})
	defer tearDown(t)

	TestGetAll(t, "container", []string{
		`{"id":1,"itemCount":0}`,
		`{"id":2,"itemCount":0}`,
		`{"id":3,"itemCount":0}`,
	})

}

func TestPostContainer(t *testing.T) {

	tearDown := SetupTest(t, []TableSetup{
		{"container", []string{}},
	})
	defer tearDown(t)

	TestPost(t, "container", []PostTestCase{
		{``, `{"containerId":1}`, 200},
		{``, `{"containerId":2}`, 200},
		{``, `{"containerId":3}`, 200},
	})

}

func TestGetContainerById(t *testing.T) {

	tearDown := SetupTest(t, []TableSetup{
		{"container", []string{``, ``, ``}},
	})
	defer tearDown(t)

	TestGetById(t, "container", 2, `{"id":2,"itemCount":0}`)

}

func TestPutContainerById(t *testing.T) {

	tearDown := SetupTest(t, []TableSetup{
		{"container", []string{``, ``, ``}},
	})
	defer tearDown(t)

	TestPutById(t, "container", 1, `{"itemCount":0}`)

}

func TestDeleteContainerById(t *testing.T) {

	tearDown := SetupTest(t, []TableSetup{
		{"container", []string{``, ``, ``}},
	})
	defer tearDown(t)

	TestDeleteById(t, "container", 1)

}
