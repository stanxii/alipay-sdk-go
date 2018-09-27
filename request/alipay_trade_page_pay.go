package request

type AlipayTradePagePayRequest struct {
	OutTradeNo           string                         `json:"out_trade_no"`
	ProductCode          string                         `json:"product_code"`
	TotalAmount          string                         `json:"total_amount"`
	Subject              string                         `json:"subject"`
	Body                 string                         `json:"body,omitempty"`
	GoodsDetail          string                         `json:"goods_detail,omitempty"`
	PassbackParams       string                         `json:"passback_params,omitempty"`
	ExtendParams         AlipayTradePagePayExtendParams `json:"extend_params,omitempty"`
	GoodsType            string                         `json:"goods_type,omitempty"`
	TimeoutExpress       string                         `json:"timeout_express,omitempty"`
	EnablePayChannels    string                         `json:"enable_pay_channels,omitempty"`
	DisablePayChannels   string                         `json:"disable_pay_channels,omitempty"`
	AuthToken            string                         `json:"auth_token,omitempty"`
	QrPayMode            string                         `json:"qr_pay_mode,omitempty"`
	QrcodeWidth          string                         `json:"qrcode_width,omitempty"`
	SysServiceProviderId string                         `json:"sys_service_provider_id,omitempty"`
	HbFqNum              string                         `json:"hb_fq_num,omitempty"`
	HbFqSellerPercent    string                         `json:"hb_fq_seller_percent,omitempty"`
}

type AlipayTradePagePayExtendParams struct {
	SysServiceProviderId string `json:"sys_service_provider_id,omitempty"`
	HbFqNum              string `json:"hb_fq_num,omitempty"`
	HbFqSellerPercent    string `json:"hb_fq_seller_percent,omitempty"`
}

var (
	AlipayTradePagePayMethod = "alipay.trade.page.pay"
)