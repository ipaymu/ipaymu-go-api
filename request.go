package ipaymu_go_api

type RequestDirectMaster struct {
	Name          *string       `json:"name"`
	Phone         *string       `json:"phone"`
	Email         *string       `json:"email"`
	Amount        float64       `json:"amount"`
	NotifyUrl     *string       `json:"notifyUrl"`
	Expired       *int8         `json:"expired"`
	ExpiredType   *ExpiredType  `json:"expiredType"`
	Comments      *string       `json:"comments"`
	ReferenceId   *string       `json:"referenceId"`
	PaymentMethod PaymentMethod `json:"paymentMethod"`
}

// AddBuyer sets the buyer's name, phone, and email for a direct payment request.
//
// Parameters:
// - name: A string representing the buyer's name.
// - phone: A string representing the buyer's phone number.
// - email: A string representing the buyer's email address.
//
// This function does not return any value. It modifies the RequestDirectMaster struct's Name, Phone, and Email fields.
func (r *RequestDirectMaster) AddBuyer(name, phone, email string) {
    r.Name = &name
    r.Phone = &phone
    r.Email = &email
}

type Product struct {
	Product []string  `json:"product,omitempty"`
	Qty     []int8    `json:"qty,omitempty"`
	Price   []float64 `json:"price,omitempty"`
}

type ProductWithCOD struct {
	Product
	Weight []float32 `json:"weight,omitempty"`
	Width  []float32 `json:"width,omitempty"`
	Height []float32 `json:"height,omitempty"`
	Length []float32 `json:"length,omitempty"`
}

type RequestDirectVA struct {
	RequestDirectMaster
	PaymentChannel PaymentChannelVA `json:"paymentChannel"`
	Product
}

// NewRequestDirectVA creates a new instance of RequestDirectVA with the specified payment channel.
//
// This function initializes a new RequestDirectVA struct and sets the payment method to VirtualAccount.
// It also assigns the provided payment channel to the RequestDirectVA struct.
//
// Parameters:
// - channel: A PaymentChannelVA representing the desired payment channel for the direct payment request.
//
// Return:
// - A pointer to a new RequestDirectVA instance with the specified payment channel.
func NewRequestDirectVA(channel PaymentChannelVA) *RequestDirectVA {
    req := &RequestDirectVA{}
    req.PaymentMethod = VirtualAccount
    req.PaymentChannel = channel
    return req
}

type RequestDirectCOD struct {
	RequestDirectMaster
	PaymentChannel PaymentChannelCOD `json:"paymentChannel"`
	*ProductWithCOD
	DeliveryArea    string `json:"deliveryArea"`
	DeliveryAddress string `json:"deliveryAddress"`
}

// NewRequestDirectCOD creates a new instance of RequestDirectCOD with the specified COD payment channel.
//
// This function initializes a new RequestDirectCOD struct and sets the payment method to COD.
// It also assigns the provided COD payment channel to the RequestDirectCOD struct.
//
// Parameters:
// - codChannel: A PaymentChannelCOD representing the desired COD payment channel for the direct payment request.
//
// Return:
// - A pointer to a new RequestDirectCOD instance with the specified COD payment channel.
func NewRequestDirectCOD(codChannel PaymentChannelCOD) *RequestDirectCOD {
    req := &RequestDirectCOD{}
    req.PaymentMethod = COD
    req.PaymentChannel = codChannel
    return req
}

type RequestDirectConStore struct {
	RequestDirectMaster
	PaymentChannel PaymentChannelConStore `json:"paymentChannel"`
}

type RequestRedirect struct {
	Description   *string        `json:"description"`
	ReturnUrl     *string        `json:"returnUrl"`
	NotifyUrl     *string        `json:"notifyUrl"`
	CancelUrl     *string        `json:"cancelUrl"`
	ReferenceId   *string        `json:"referenceId"`
	Product       []string       `json:"product"`
	Qty           []int8         `json:"qty"`
	Price         []float64      `json:"price"`
	Weight        []float32      `json:"weight"`
	Dimension     []string       `json:"dimension"`
	BuyerName     *string        `json:"buyerName"`
	BuyerEmail    *string        `json:"buyerEmail"`
	BuyerPhone    *string        `json:"buyerPhone"`
	PickupArea    *string        `json:"pickupArea"`
	PickupAddress *string        `json:"pickupAddress"`
	PaymentMethod *PaymentMethod `json:"paymentMethod"`
}

func NewRequestRedirect() *RequestRedirect {
	return &RequestRedirect{}
}

// AddProduct adds a product to the RequestRedirect struct.
//
// This function appends the provided product details to the respective fields of the RequestRedirect struct.
// It accepts the following parameters:
// - product: A string representing the name of the product.
// - qty: An int8 representing the quantity of the product.
// - price: A float64 representing the price of the product.
// - weight: A pointer to a float32 representing the weight of the product. If nil, it is ignored.
// - dimension: A pointer to a string representing the dimension of the product. If nil, it is ignored.
//
// The function does not return any value. It modifies the RequestRedirect struct's Product, Qty, Price, Weight, and Dimension fields.
func (r *RequestRedirect) AddProduct(product string, qty int8, price float64, weight *float32, dimension *string) {
    r.Product = append(r.Product, product)
    r.Qty = append(r.Qty, qty)
    r.Price = append(r.Price, price)
    if weight != nil {
        r.Weight = append(r.Weight, *weight)
    }
    if dimension != nil {
        r.Dimension = append(r.Dimension, *dimension)
    }
}

type RequestTransactionHistory struct {
	ID         *int              `json:"id"`
	Status     *PaymentStatus    `json:"status"`
	Date       *FilterDate       `json:"date"`
	StartDate  *string           `json:"startdate"` // Format Y-m-d
	EndDate    *string           `json:"enddate"`   // Format Y-m-d
	Page       *int8             `json:"page"`
	OrderBy    *FilterOrderBy    `json:"orderBy"`
	Order      *FilterOrder      `json:"order"`
	Limit      *int8             `json:"limit"` // Max Limit 20
	Lang       *FilterLanguage   `json:"lang"`
	BulkId     *string           `json:"bulkId"` // joined with comma (,)
	Account    *string           `json:"account"`
	LockStatus *FilterLockStatus `json:"lockStatus"`
}

func NewRequestTransactionHistory() *RequestTransactionHistory {
	return &RequestTransactionHistory{}
}

type RequestCallBack struct {
	TrxID       int    `json:"trx_id"`
	Status      string `json:"status"`
	StatusCode  int8   `json:"status_code"`
	SID         string `json:"sid"`
	ReferenceID string `json:"reference_id"`
}
