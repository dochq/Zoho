package inventory

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

//https://www.zoho.com/inventory/api/v1/#Contacts_Create_a_Contact
//func (c *API) CreateContact(request interface{}, OrganizationID string, params map[string]zoho.Parameter) (data ListContactsResponse, err error) {
func (c *API) CreateContact(request Contact, enablePortal bool) (data CreateContactResponse, err error) {

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
	ContactName      string           `json:"contact_name"`
	CompanyName      string           `json:"company_name"`
	PaymentTerms     int              `json:"payment_terms"`
	CurrencyID       int64            `json:"currency_id"`
	Website          string           `json:"website"`
	ContactType      string           `json:"contact_type"`
	CustomFields     []CustomFields   `json:"custom_fields"`
	BillingAddress   Address          `json:"billing_address"`
	ShippingAddress  Address          `json:"shipping_address"`
	ContactPersons   []ContactPerson  `json:"contact_persons"`
	DefaultTemplates DefaultTemplates `json:"default_templates"`
	LanguageCode     string           `json:"language_code"`
	Notes            string           `json:"notes"`
	TaxExemptionID   int64            `json:"tax_exemption_id"`
	TaxAuthorityID   int64            `json:"tax_authority_id"`
	TaxID            int64            `json:"tax_id"`
	IsTaxable        bool             `json:"is_taxable"`
	Facebook         string           `json:"facebook"`
	Twitter          string           `json:"twitter"`
	PlaceOfContact   string           `json:"place_of_contact"`
	GstNo            string           `json:"gst_no"`
	GstTreatment     string           `json:"gst_treatment"`
}
type CustomFields struct {
	Value string `json:"value"`
	Index int    `json:"index"`
}
type Address struct {
	Attention string `json:"attention"`
	Address   string `json:"address"`
	Street2   string `json:"street2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
}

type ContactPerson struct {
	Salutation       string `json:"salutation"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Mobile           string `json:"mobile"`
	IsPrimaryContact bool   `json:"is_primary_contact"`
}
type DefaultTemplates struct {
	InvoiceTemplateID           int64  `json:"invoice_template_id"`
	InvoiceTemplateName         string `json:"invoice_template_name"`
	EstimateTemplateID          int64  `json:"estimate_template_id"`
	EstimateTemplateName        string `json:"estimate_template_name"`
	CreditnoteTemplateID        int64  `json:"creditnote_template_id"`
	CreditnoteTemplateName      string `json:"creditnote_template_name"`
	InvoiceEmailTemplateID      int64  `json:"invoice_email_template_id"`
	InvoiceEmailTemplateName    string `json:"invoice_email_template_name"`
	EstimateEmailTemplateID     int64  `json:"estimate_email_template_id"`
	EstimateEmailTemplateName   string `json:"estimate_email_template_name"`
	CreditnoteEmailTemplateID   int64  `json:"creditnote_email_template_id"`
	CreditnoteEmailTemplateName string `json:"creditnote_email_template_name"`
}

type CreateContactResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Contact struct {
		ContactID                        string  `json:"contact_id"`
		ContactName                      string  `json:"contact_name"`
		CompanyName                      string  `json:"company_name"`
		HasTransaction                   bool    `json:"has_transaction"`
		ContactType                      string  `json:"contact_type"`
		IsTaxable                        bool    `json:"is_taxable"`
		TaxID                            string  `json:"tax_id"`
		TaxName                          string  `json:"tax_name"`
		TaxPercentage                    float64 `json:"tax_percentage"`
		TaxExemptionID                   string  `json:"tax_exemption_id"`
		TaxAuthorityID                   string  `json:"tax_authority_id"`
		GSTNo                            string  `json:"gst_no"`
		GSTTreatment                     string  `json:"gst_treatment"`
		IsLinkedWithZohocrm              bool    `json:"is_linked_with_zohocrm"`
		Website                          string  `json:"website"`
		PrimaryContactID                 string  `json:"primary_contact_id"`
		PaymentTerms                     int64   `json:"payment_terms"`
		PaymentTermsLabel                string  `json:"payment_terms_label"`
		CurrencyID                       string  `json:"currency_id"`
		CurrencyCode                     string  `json:"currency_code"`
		CurrencySymbol                   string  `json:"currency_symbol"`
		LanguageCode                     string  `json:"language_code"`
		OutstandingReceivableAmount      float64 `json:"outstanding_receivable_amount"`
		OutstandingReceivableAmountBcy   float64 `json:"outstanding_receivable_amount_bcy"`
		UnusedCreditsReceivableAmount    float64 `json:"unused_credits_receivable_amount"`
		UnusedCreditsReceivableAmountBcy float64 `json:"unused_credits_receivable_amount_bcy"`
		Status                           string  `json:"status"`
		PaymentReminderEnabled           bool    `json:"payment_reminder_enabled"`
		CustomFields                     []struct {
			Value string `json:"value"`
			Index int64  `json:"index"`
			Label string `json:"label"`
		} `json:"custom_fields"`
		BillingAddress   Address          `json:"billing_address"`
		ShippingAddress  Address          `json:"shipping_address"`
		Facebook         string           `json:"facebook"`
		Twitter          string           `json:"twitter"`
		ContactPersons   []ContactPerson  `json:"contact_persons"`
		DefaultTemplates DefaultTemplates `json:"default_templates"`
		Notes            string           `json:"notes"`
		CreatedTime      string           `json:"created_time"`
		LastModifiedTime string           `json:"last_modified_time"`
	} `json:"contact"`
}
