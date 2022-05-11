package docker

import "testing"

func TestExtractHost(t *testing.T) {
	id := "vibrant_maxwell"
	port := "5432"

	host, err := ExtractHost(id, port)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Extracted host: ", host)

}
func TestStartContainer(t *testing.T) {
	//docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres  -d postgres

}
