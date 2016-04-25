package flight

import "testing"

func TestGoFlight(t *testing.T) {
	mock := MockFlight{}
	resp, err := mock.GetFlight(Request{})
	if err != nil {
		t.Error(err)
	}
	if resp.Trips.Data.Airport[0].Code != "EWR" {
		t.Errorf("Error parsing data, wanted [EWR] - received [%s]", resp.Trips.Data.Airport[0].Code)
	}
}
