package inventory

import (
	"fmt"
	"math/rand"

	zoho "github.com/schmorrison/Zoho"
)

var (
	InventoryAPIEndpoint string = "https://inventory.zoho.%s/api/v1/"
)

const (
	InvoiceAPIEndpointHeader string = "X-com-zoho-invoice-organizationid"
	ContactsModule           string = "contacts"
	SalesOrderModule         string = "salesorders"
)

type CustomFieldRequest struct {
	CustomfieldID string `json:"customfield_id,omitempty"`
	Label         string `json:"label"`
	Value         string `json:"value,omitempty"`
}

// API is used for interacting with the Zoho inventory API
// the exposed methods are primarily access to inventory modules which provide access to inventory Methods
type API struct {
	*zoho.Zoho
	id string
}

// New returns a *invoice.API with the provided zoho.Zoho as an embedded field
func New(z *zoho.Zoho) *API {
	id := func() string {
		var id []byte
		keyspace := "abcdefghijklmnopqrutuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		for i := 0; i < 25; i++ {
			id = append(id, keyspace[rand.Intn(len(keyspace))])
		}
		return string(id)
	}()

	API := API{
		Zoho: z,
		id:   id,
	}
	InventoryAPIEndpoint = fmt.Sprintf("https://inventory.zoho.%s/api/v1/", z.ZohoTLD)
	return &API
}
