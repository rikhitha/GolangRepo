package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) Findall() ([]Customer, error) {

	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {

	customers := []Customer{
		{Id: "1001", Name: "Ashish", City: "Chennai", Zipcode: "600100", DateOfBirth: "25-12-1998", Status: "available"},
		{Id: "1001", Name: "abi", City: "Chennai", Zipcode: "600100", DateOfBirth: "25-11-1998", Status: "away"},
	}

	return CustomerRepositoryStub{customers: customers}
}
