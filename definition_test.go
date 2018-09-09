package sophos_test

import "testing"

func TestDefinition_Get(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	_, err := client.GetEndpointSwag(dnsMock{})
	if err != nil {
		t.Error(err)
	}

	d := dnsMock{}.Definition()
	_, err = d.GetSwag(client, errOption)
	if err == nil {
		t.Error(err)
	}
}
