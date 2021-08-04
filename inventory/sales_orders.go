package inventory

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

// https://www.zoho.com/inventory/api/v1/#Sales_Orders_Create_a_Sales_Order
func (c *API) CreateSalesOrder(request SalesOrder, enablePortal bool) (data CreateSalesOrderResponse, err error) {

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
	CustomerID            int64          `json:"customer_id"`
	SalesorderNumber      string         `json:"salesorder_number"`
	Date                  string         `json:"date"`
	ShipmentDate          string         `json:"shipment_date"`
	CustomFields          []CustomFields `json:"custom_fields"`
	ReferenceNumber       string         `json:"reference_number"`
	LineItems             []LineItems    `json:"line_items"`
	Notes                 string         `json:"notes"`
	Terms                 string         `json:"terms"`
	Discount              string         `json:"discount"`
	IsDiscountBeforeTax   bool           `json:"is_discount_before_tax"`
	DiscountType          string         `json:"discount_type"`
	ShippingCharge        int            `json:"shipping_charge"`
	DeliveryMethod        string         `json:"delivery_method"`
	Adjustment            int            `json:"adjustment"`
	PricebookID           int64          `json:"pricebook_id"`
	SalespersonID         int64          `json:"salesperson_id"`
	AdjustmentDescription string         `json:"adjustment_description"`
	IsInclusiveTax        bool           `json:"is_inclusive_tax"`
	ExchangeRate          int            `json:"exchange_rate"`
	TemplateID            int64          `json:"template_id"`
	Documents             []Documents    `json:"documents"`
	BillingAddressID      int64          `json:"billing_address_id"`
	ShippingAddressID     int64          `json:"shipping_address_id"`
	PlaceOfSupply         string         `json:"place_of_supply"`
	GstTreatment          string         `json:"gst_treatment"`
	GstNo                 string         `json:"gst_no"`
}

type LineItems struct {
	ItemID        int64  `json:"item_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Rate          int    `json:"rate"`
	Quantity      int    `json:"quantity"`
	Unit          string `json:"unit"`
	TaxID         int64  `json:"tax_id"`
	TaxName       string `json:"tax_name"`
	TaxType       string `json:"tax_type"`
	TaxPercentage int    `json:"tax_percentage"`
	ItemTotal     int    `json:"item_total"`
	WarehouseID   int64  `json:"warehouse_id"`
	HsnOrSac      int    `json:"hsn_or_sac"`
}
type Documents struct {
	CanSendInMail     bool   `json:"can_send_in_mail"`
	FileName          string `json:"file_name"`
	FileType          string `json:"file_type"`
	FileSizeFormatted string `json:"file_size_formatted"`
	AttachmentOrder   int    `json:"attachment_order"`
	DocumentID        int64  `json:"document_id"`
	FileSize          int    `json:"file_size"`
}

type CreateSalesOrderResponse struct {
	Code       int64      `json:"code"`
	Message    string     `json:"message"`
	SalesOrder SalesOrder `json:"sales_order"`
}
