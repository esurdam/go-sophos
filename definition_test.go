package sophos_test

import "testing"

func TestDefinition_Get(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	d := dnsMock{}.Definition()
	_, err := d.GetSwag(client)
	if err != nil {
		t.Error(err)
	}

	_, err = d.GetSwag(client, errOption)
	if err == nil {
		t.Error(err)
	}
}