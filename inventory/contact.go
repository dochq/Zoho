package inventory

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

//https://www.zoho.com/inventory/api/v1/#Contacts_Create_a_Contact
//func (c *API) CreateContact(request interface{}, OrganizationID string, params map[string]zoho.Parameter) (data ListContactsResponse, err error) {
func (c *API) CreateContact(request Contact) (data CreateContactResponse, err error) {

	endpoint := zoho.Endpoint{
		Name:         ContactsModule,
		URL:          fmt.Sprintf(InventoryAPIEndpoint+"%s", ContactsModule),
		Method:       zoho.HTTPPost,
		ResponseData: &CreateContactResponse{},
		URLParameters: map[string]zoho.Parameter{
			"filter_by": "",
		},
		RequestBody: &request,
		BodyFormat:  zoho.JSON_STRING,
		Headers: map[string]string{
			InvoiceAPIEndpointHeader: c.OrganizationID,
		},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CreateContactResponse{}, fmt.Errorf("Failed to create contact: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CreateContactResponse); ok {
		// Check if the request succeeded
		if v.Contact.ContactID == "" {
			return *v, fmt.Errorf("Failed to create contact: %s", v.Message)
		}

		return *v, nil
	}
	return CreateContactResponse{}, fmt.Errorf("Data retrieved was not 'CreateContactResponse'")
}

type Contact struct {
	ContactID        string           `json:"contact_id,omitempty"`
	ContactName      string           `json:"contact_name,omitempty"`
	CompanyName      string           `json:"company_name,omitempty"`
	PaymentTerms     int              `json:"payment_terms,omitempty"`
	CurrencyID       string           `json:"currency_id,omitempty"`
	Website          string           `json:"website,omitempty"`
	ContactType      string           `json:"contact_type,omitempty"`
	BillingAddress   Address          `json:"billing_address,omitempty"`
	ShippingAddress  Address          `json:"shipping_address,omitempty"`
	ContactPersons   []ContactPerson  `json:"contact_persons,omitempty"`
	DefaultTemplates DefaultTemplates `json:"default_templates,omitempty"`
	LanguageCode     string           `json:"language_code,omitempty"`
	Notes            string           `json:"notes,omitempty"`
	TaxExemptionID   int64            `json:"tax_exemption_id,omitempty"`
	TaxAuthorityID   int64            `json:"tax_authority_id,omitempty"`
	TaxID            int64            `json:"tax_id,omitempty"`
	IsTaxable        bool             `json:"is_taxable,omitempty"`
	Facebook         string           `json:"facebook,omitempty"`
	Twitter          string           `json:"twitter,omitempty"`
	PlaceOfContact   string           `json:"place_of_contact,omitempty"`
	GstNo            string           `json:"gst_no,omitempty"`
	GstTreatment     string           `json:"gst_treatment,omitempty"`
}

type Address struct {
	Attention string `json:"attention,omitempty"`
	Address   string `json:"address,omitempty"`
	Street2   string `json:"street2,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Zip       string `json:"zip,omitempty"`
	Country   string `json:"country,omitempty"`
}

type ContactPerson struct {
	Salutation       string `json:"salutation,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Email            string `json:"email,omitempty"`
	Phone            string `json:"phone,omitempty"`
	Mobile           string `json:"mobile,omitempty"`
	IsPrimaryContact bool   `json:"is_primary_contact,omitempty"`
}
type DefaultTemplates struct {
	InvoiceTemplateID           int64  `json:"invoice_template_id,omitempty"`
	InvoiceTemplateName         string `json:"invoice_template_name,omitempty"`
	EstimateTemplateID          int64  `json:"estimate_template_id,omitempty"`
	EstimateTemplateName        string `json:"estimate_template_name,omitempty"`
	CreditnoteTemplateID        int64  `json:"creditnote_template_id,omitempty"`
	CreditnoteTemplateName      string `json:"creditnote_template_name,omitempty"`
	InvoiceEmailTemplateID      int64  `json:"invoice_email_template_id,omitempty"`
	InvoiceEmailTemplateName    string `json:"invoice_email_template_name,omitempty"`
	EstimateEmailTemplateID     int64  `json:"estimate_email_template_id,omitempty"`
	EstimateEmailTemplateName   string `json:"estimate_email_template_name,omitempty"`
	CreditnoteEmailTemplateID   int64  `json:"creditnote_email_template_id,omitempty"`
	CreditnoteEmailTemplateName string `json:"creditnote_email_template_name,omitempty"`
}

type CreateContactResponse struct {
	Code    int     `json:"code,omitempty"`
	Message string  `json:"message,omitempty"`
	Contact Contact `json:"contact,omitempty"`
}
