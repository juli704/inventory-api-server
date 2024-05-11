package test

import (
	"testing"
)

func TestGetContainers(t *testing.T) {

	tearDownTest := SetupPopulatedTest(t, "container", []string{``, ``, ``})
	defer tearDownTest(t)

	TestGetAll(t, "container", []string{
		`{"id":1,"itemCount":0}`,
		`{"id":2,"itemCount":0}`,
		`{"id":3,"itemCount":0}`,
	})

}

func TestPostContainer(t *testing.T) {

	tearDownTest := SetupEmptyTest(t, "container")
	defer tearDownTest(t)

	TestPost(t, "container", []PostTestCase{
		{``, `{"containerId":1}`, 200},
		{``, `{"containerId":2}`, 200},
		{``, `{"containerId":3}`, 200},
	})

}

func TestGetContainerById(t *testing.T) {

	tearDownTest := SetupPopulatedTest(t, "container", []string{``, ``, ``})
	defer tearDownTest(t)

	TestGetById(t, "container", 2, `{"id":2,"itemCount":0}`)

}

func TestPutContainerById(t *testing.T) {

	tearDownTest := SetupPopulatedTest(t, "container", []string{``, ``, ``})
	defer tearDownTest(t)

	TestPutById(t, "container", 1, `{"itemCount":0}`)

}

func TestDeleteContainerById(t *testing.T) {

	tearDownTest := SetupPopulatedTest(t, "container", []string{``, ``, ``})
	defer tearDownTest(t)

	TestDeleteById(t, "container", 1)

}
