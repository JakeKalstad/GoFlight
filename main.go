package flight

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Passengers struct {
	Kind              string `json:kind`
	AdultCount        int    `json:adultCount`
	ChildCount        int    `json:childCount`
	InfantInLapCount  int    `json:infantInLapCount`
	InfantInSeatCount int    `json:infantInSeatCount`
	SeniorCount       int    `json:seniorCount`
}

type PermittedDepartureTime struct {
	Kind         string `json:kind`
	EarliestTime string `json:earliestTime`
	LatestTime   string `json:latestTime`
}

type RequestSlice struct {
	Kind                   string                 `json:kind`
	Origin                 string                 `json:origin`
	Destination            string                 `json:destination`
	Date                   string                 `json:date`
	MaxStops               int                    `json:maxStops`
	MaxConnectionDuration  int                    `json:maxConnectionDuration`
	PreferredCabin         string                 `json:preferredCabin`
	PermittedDepartureTime PermittedDepartureTime `json:permittedDepartureTime`
	PermittedCarrier       []string               `json:permittedCarrier`
	ProhibitedCarrier      []string               `json:prohibitedCarrier`
}

type Request struct {
	Passengers  Passengers     `json:passengers`
	Slice       []RequestSlice `json:slice`
	MaxPrice    string         `json:maxPrice`
	SaleCountry string         `json:saleCountry`
	Refundable  string         `json:refundable`
}

type City struct {
	Kind    string `json:kind`
	Code    string `json:code`
	Name    string `json:name`
	Country string `json:country`
}

type Airport struct {
	Kind string `json:kind`
	Name string `json:name`
	Code string `json:code`
	City string `json:city`
}

type Tax struct {
	Kind       string `json:kind`
	Name       string `json:name`
	ID         string `json:id`
	ChargeType string `json:chargeType`
	Code       string `json:code`
	Country    string `json:country`
	SalePrice  string `json:salePrice`
}

type Aircraft struct {
	Kind string `json:kind`
	Name string `json:name`
	Code string `json:code`
}

type Carrier struct {
	Kind string `json:kind`
	Name string `json:name`
	Code string `json:code`
}

type TripData struct {
	Kind     string   `json:kind`
	Airport  Airport  `json:airport`
	City     City     `json:city`
	Aircraft Aircraft `json:aircraft`
	Tax      Tax      `json:tax`
	Carrier  Carrier  `json:carrier`
}

type TripOption struct {
	Kind      string `json:kind`
	ID        string `json:id`
	SaleTotal string `json:saleTotal`
}
type Fare struct {
	Kind        string  `json:kind`
	ID          string  `json:id`
	Carrier     Carrier `json:carrier`
	Origin      string  `json:origin`
	Destination string  `json:destination`
	BasisCode   string  `json:basisCode`
	Private     bool    `json:private`
}

type Price struct {
	Kind  string `json:kind`
	Fares Fare   `json:fare`
}

type BagDescriptor struct {
	Kind           string   `json:kind`
	CommercialName string   `json:commercialName`
	Count          int      `json:count`
	Description    []string `json:description`
	Subcode        string   `json:subcode`
}

type FreeBaggageOption struct {
	Kind       string        `json:kind`
	Descriptor BagDescriptor `json:bagDescriptor`
}

type SegmentPrice struct {
	Kind               string              `json:kind`
	FareID             string              `json:fareId`
	SegmentID          string              `json:segmentId`
	FreeBaggageOptions []FreeBaggageOption `json:freeBaggageOption`
	Kilos              int                 `json:kilos`
	KilosPerPiece      int                 `json:kilosPerPiece`
	Pieces             int                 `json:pieces`
	Pounds             int                 `json:pounds`
}

type Trips struct {
	Kind                string         `json:kind`
	RequestId           string         `json:requestId`
	Data                TripData       `json:data`
	Options             TripOption     `json:tripOption`
	Pricing             []Price        `json:pricing`
	SegmentPricing      []SegmentPrice `json:segmentPricing`
	BaseFareTotal       string         `json:baseFareTotal`
	SaleFareTotal       string         `json:saleFareTotal`
	SaleTaxTotal        string         `json:saleTaxTotal`
	SaleTotal           string         `json:saleTotal`
	Passengers          Passengers     `json:passengers`
	Tax                 Tax            `json:tax`
	FareCalculation     string         `json:fareCalculation`
	LatestTicketingTime string         `json:latestTicketingTime`
	PTC                 string         `json:ptc`
	Refundable          bool           `json:refundable`
}

type Flight struct {
	Carrier string `json:carrier`
	Number  string `json:number`
}

type Leg struct {
	Kind                string `json:kind`
	Duration            int    `json:duration`
	ID                  string `json:id`
	ArrivalTime         string `json:arrivalTime`
	DepartureTime       string `json:departureTime`
	Origin              string `json:origin`
	Destination         string `json:destination`
	OriginTerminal      string `json:originTerminal`
	DestinationTerminal string `json:destinationTerminal`
	OperatingDisclosure string `json:operatingDisclosure`
	OnTimePerformance   int    `json:onTimePerformance`
	Mileage             int    `json:mileage`
	Meal                string `json:meal`
	Secure              bool   `json:secure`
	ConnectionDuration  int    `json:connectionDuration`
	ChangePlane         bool   `json:changePlane`
}

type Segment struct {
	Kind                        string `json:kind`
	Duration                    string `json:duration`
	Flight                      Flight `json:flight`
	ID                          string `json:id`
	Cabin                       string `json:cabin`
	BookingCode                 string `json:bookingCode`
	BookingCodeCount            int    `json:bookingCodeCount`
	MarriedSegmentGroup         string `json:marriedSegmentGroup`
	SubjectToGovernmentApproval bool   `json:subjectToGovernmentApproval`
	Legs                        []Leg  `json:leg`
	ConnectionDuration          int    `json:connectionDuration`
}

type ResponseSlice struct {
	Kind     string    `json:kind`
	Duration string    `json:duration`
	Segment  []Segment `json:segment`
}

type Response struct {
	Kind  string          `json:kind`
	Trips Trips           `json:trips`
	Slice []ResponseSlice `json:slice`
}

type GoFlyer interface {
	GetFlight(Request) (Response, error)
}

type GoFlight struct {
	ApiKey string
}

func (flight *GoFlight) GetFlight(apiRequest Request) (Response, error) {
	url := fmt.Sprintf("https://www.googleapis.com/qpxExpress/v1/trips/search?key=%s", flight.ApiKey)
	fmt.Println("URL:>", url)
	jsonData, err := json.Marshal(apiRequest)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	var response Response
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &response)
	return response, err
}
