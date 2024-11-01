package luksdk

func New(signSecret string) *SDK {
	return &SDK{signSecret: signSecret}
}

type SDK struct {
	signSecret string
}

// VerifySignature 验证签名是否正确
//   - sign: 待验证的签名
//   - params: 待签名的请求结构体（不含签名字段，如 sign）
func (sdk *SDK) VerifySignature(sign string, params any) error {
	verify := signature(sdk.signSecret, params)
	if verify != sign {
		return ErrInvalidSignature
	}
	return nil
}

// GenerateSignature 生成签名
func (sdk *SDK) GenerateSignature(params any) string {
	return signature(sdk.signSecret, params)
}

// GetChannelToken CFGame向接入方获取用户令牌
func (sdk *SDK) GetChannelToken(request *GetChannelTokenRequest, successHandler ...RequestHandler[*GetChannelTokenRequest, *GetChannelTokenResponse]) *Response[*GetChannelTokenResponse] {
	return generateHandler(sdk.signSecret, request.Sign, request, successHandler...)
}

// RefreshChannelToken 刷新用户令牌过期时间
func (sdk *SDK) RefreshChannelToken(request *RefreshChannelTokenRequest, successHandler ...RequestHandler[*RefreshChannelTokenRequest, *RefreshChannelTokenResponse]) *Response[*RefreshChannelTokenResponse] {
	return generateHandler(sdk.signSecret, request.Sign, request, successHandler...)
}

// GetChannelUserInfo 获取渠道用户信息
func (sdk *SDK) GetChannelUserInfo(request *GetChannelUserInfoRequest, successHandler ...RequestHandler[*GetChannelUserInfoRequest, *GetChannelUserInfoResponse]) *Response[*GetChannelUserInfoResponse] {
	return generateHandler(sdk.signSecret, request.Sign, request, successHandler...)
}

// CreateChannelOrder 向渠道下订单
func (sdk *SDK) CreateChannelOrder(request *CreateChannelOrderRequest, successHandler ...RequestHandler[*CreateChannelOrderRequest, CreateChannelOrderResponse]) *Response[CreateChannelOrderResponse] {
	return generateHandler(sdk.signSecret, request.Sign, request, successHandler...)
}

// NotifyChannelOrder 下注开奖通知结果
func (sdk *SDK) NotifyChannelOrder(request *NotifyChannelOrderRequest, successHandler ...RequestHandler[*NotifyChannelOrderRequest, NotifyChannelOrderResponse]) *Response[NotifyChannelOrderResponse] {
	return generateHandler(sdk.signSecret, request.Sign, request, successHandler...)
}

// NotifyGame 向渠道通知游戏状态
func (sdk *SDK) NotifyGame(request *NotifyGameRequest, successHandler ...RequestHandler[*NotifyGameRequest, *NotifyGameResponse]) *Response[*NotifyGameResponse] {
	return generateHandler(sdk.signSecret, request.Sign, request, successHandler...)
}
