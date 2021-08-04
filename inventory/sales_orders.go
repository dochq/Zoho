package inventory

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

// https://www.zoho.com/inventory/api/v1/#Sales_Orders_Create_a_Sales_Order
func (c *API) CreateSalesOrder(request SalesOrder) (data CreateSalesOrderResponse, err error) {

	endpoint := zoho.Endpoint{
		Name:         ContactsModule,
		URL:          fmt.Sprintf(InventoryAPIEndpoint+"%s", SalesOrderModule),
		Method:       zoho.HTTPPost,
		ResponseData: &CreateSalesOrderResponse{},
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
		return CreateSalesOrderResponse{}, fmt.Errorf("Failed to create sales order: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CreateSalesOrderResponse); ok {
		// Check if the request succeeded
		if v.Code != 0 {
			return *v, fmt.Errorf("Failed to create sales order: %s", v.Message)
		}

		return *v, nil
	}
	return CreateSalesOrderResponse{}, fmt.Errorf("Data retrieved was not 'CreateSalesOrderResponse'")
}

type SalesOrder struct {
	CustomerID            int64       `json:"customer_id,omitempty"`
	SalesorderNumber      string      `json:"salesorder_number,omitempty"`
	Date                  string      `json:"date,omitempty"`
	ShipmentDate          string      `json:"shipment_date,omitempty"`
	ReferenceNumber       string      `json:"reference_number,omitempty"`
	LineItems             []LineItems `json:"line_items,omitempty"`
	Notes                 string      `json:"notes,omitempty"`
	Terms                 string      `json:"terms,omitempty"`
	Discount              string      `json:"discount,omitempty"`
	IsDiscountBeforeTax   bool        `json:"is_discount_before_tax,omitempty"`
	DiscountType          string      `json:"discount_type,omitempty"`
	ShippingCharge        int         `json:"shipping_charge,omitempty"`
	DeliveryMethod        string      `json:"delivery_method,omitempty"`
	Adjustment            int         `json:"adjustment,omitempty"`
	PricebookID           int64       `json:"pricebook_id,omitempty"`
	SalespersonID         int64       `json:"salesperson_id,omitempty"`
	AdjustmentDescription string      `json:"adjustment_description,omitempty"`
	IsInclusiveTax        bool        `json:"is_inclusive_tax,omitempty"`
	ExchangeRate          int         `json:"exchange_rate,omitempty"`
	TemplateID            int64       `json:"template_id,omitempty"`
	Documents             []Documents `json:"documents,omitempty"`
	BillingAddressID      int64       `json:"billing_address_id,omitempty"`
	ShippingAddressID     int64       `json:"shipping_address_id,omitempty"`
	PlaceOfSupply         string      `json:"place_of_supply,omitempty"`
	GstTreatment          string      `json:"gst_treatment,omitempty"`
	GstNo                 string      `json:"gst_no,omitempty"`
}

type LineItems struct {
	ItemID        int64  `json:"item_id,omitempty"`
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
	Rate          int    `json:"rate,omitempty"`
	Quantity      int    `json:"quantity,omitempty"`
	Unit          string `json:"unit,omitempty"`
	TaxID         int64  `json:"tax_id,omitempty"`
	TaxName       string `json:"tax_name,omitempty"`
	TaxType       string `json:"tax_type,omitempty"`
	TaxPercentage int    `json:"tax_percentage,omitempty"`
	ItemTotal     int    `json:"item_total,omitempty"`
	WarehouseID   int64  `json:"warehouse_id,omitempty"`
	HsnOrSac      int    `json:"hsn_or_sac,omitempty"`
}
type Documents struct {
	CanSendInMail     bool   `json:"can_send_in_mail,omitempty"`
	FileName          string `json:"file_name,omitempty"`
	FileType          string `json:"file_type,omitempty"`
	FileSizeFormatted string `json:"file_size_formatted,omitempty"`
	AttachmentOrder   int    `json:"attachment_order,omitempty"`
	DocumentID        int64  `json:"document_id,omitempty"`
	FileSize          int    `json:"file_size,omitempty"`
}

type CreateSalesOrderResponse struct {
	Code       int64      `json:"code,omitempty"`
	Message    string     `json:"message,omitempty"`
	SalesOrder SalesOrder `json:"sales_order,omitempty"`
}
