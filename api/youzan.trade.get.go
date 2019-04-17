package api

import (
    "github.com/uniplaces/carbon"
    "youzango"
)

func NewTradeMethod(tid string, accessToken string) *Method {
    data := struct {
        Tid string `json:"tid"`
    }{
        Tid: tid,
    }
    jsonData, err := youzango.BuildJson(data)
    if err != nil {
        panic(err)
    }
    method := NewMethod(accessToken)
    method.Name = "youzan.trade.get"
    method.Version = "4.0.0"
    method.JsonData = jsonData
    return method
}

type TradeResponse struct {
    BaseResponse
    // 订单数据
    Data Trade `json:"data"`
}

// Trade
type Trade struct {
    // 交易基础信息结构体
    FullOrderInfo FullOrderInfo `json:"full_order_info"`
    // 订单退款信息结构体
    RefundOrder RefundOrder `json:"refund_order"`
    // 订单发货详情结构体
    DeliveryOrder DeliveryOrder `json:"delivery_order"`
    // 订单优惠详情结构体
    OrderPromotion OrderPromotion `json:"order_promotion"`
}

// FullOrderInfo
type FullOrderInfo struct {
    // 交易基础信息
    OrderInfo OrderInfo `json:"order_info"`
    // 交易来源信息
    SourceInfo SourceInfo `json:"source_info"`
    // 订单买家信息
    BuyerInfo BuyerInfo `json:"buyer_info"`
    // 订单支付信息
    PayInfo PayInfo `json:"pay_info"`
    // 标记信息
    RemarkInfo RemarkInfo `json:"remark_info"`
    // 收货地址信息
    AddressInfo AddressInfo `json:"address_info"`
    // 交易明细结
    Orders []Order `json:"orders"`
    // 送礼订单子单
    ChildInfo ChildInfo `json:"child_info"`
}

// OrderInfo 交易基础信息
type OrderInfo struct {
    // 主订单状态
    // WAIT_BUYER_PAY （等待买家付款，定金预售描述：定金待付、等待尾款支付开始、尾款待付）；
    // TRADE_PAID（订单已支付 ）；
    // WAIT_CONFIRM（待确认，包含待成团、待接单等等。即：买家已付款，等待成团或等待接单）；
    // WAIT_SELLER_SEND_GOODS（等待卖家发货，即：买家已付款）；
    // WAIT_BUYER_CONFIRM_GOODS (等待买家确认收货，即：卖家已发货) ；
    // TRADE_SUCCESS（买家已签收以及订单成功）；
    // TRADE_CLOSED（交易关闭）；
    // PS：TRADE_PAID状态仅代表当前订单已支付成功，表示瞬时状态，稍后会自动修改成后面的状态。如果不关心此状态请再次请求详情接口获取下一个状态。
    Status string `json:"status"`
    // 主订单类型
    // 0:普通订单;
    // 1:送礼订单;
    // 2:代付;
    // 3:分销采购单;
    // 4:赠品;
    // 5:心愿单;
    // 6:二维码订单;
    // 7:合并付货款;
    // 8:1分钱实名认证;
    // 9:品鉴;
    // 10:拼团;
    // 15:返利;
    // 35:酒店;
    // 40:外卖;
    // 41:堂食点餐;
    // 46:外卖买单;
    // 51:全员开店;
    // 61:线下收银台订单;
    // 71:美业预约单;
    // 72:美业服务单;
    // 75:知识付费;
    // 81:礼品卡;
    // 100:批发
    Type string `json:"type"`
    // 订单号
    Tid string `json:"tid"`
    // 主订单状态 描述
    StatusStr string `json:"status_str"`
    // 支付类型
    // 0:默认值,未支付;
    // 1:微信自有支付;
    // 2:支付宝wap;
    // 3:支付宝wap;
    // 5:财付通;
    // 7:代付;
    // 8:联动优势;
    // 9:货到付款;
    // 10:大账号代销;
    // 11:受理模式;
    // 12:百付宝;
    // 13:sdk支付;
    // 14:合并付货款;
    // 15:赠品;
    // 16:优惠兑换;
    // 17:自动付货款;
    // 18:爱学贷;
    // 19:微信wap;
    // 20:微信红包支付;
    // 21:返利;
    // 22:ump红包;
    // 24:易宝支付;
    // 25:储值卡;
    // 27:qq支付;
    // 28:有赞E卡支付;
    // 29:微信条码;
    // 30:支付宝条码;
    // 33:礼品卡支付;
    // 35:会员余额;
    // 72:微信扫码二维码支付;
    // 100:代收账户;
    // 300:储值账户;
    // 400:保证金账户;
    // 101:收款码;
    // 102:微信;
    // 103:支付宝;
    // 104:刷卡;
    // 105:二维码台卡;
    // 106:储值卡;
    // 107:有赞E卡;
    // 110:标记收款-自有微信支付;
    // 111:标记收款-自有支付宝;
    // 112:标记收款-自有POS刷卡;
    // 113:通联刷卡支付;
    // 200:记账账户;
    // 201:现金
    PayType string `json:"pay_type"`
    // 店铺类型
    // 0:微商城;
    // 1:微小店;
    // 2:爱学贷微商城;
    // 3:批发店铺;
    // 4:批发商城;
    // 5:外卖;
    // 6:美业;
    // 7:超级门店;
    // 8:收银;
    // 9:收银加微商城;
    // 10:零售总部;
    // 99:有赞开放平台平台型应用创建的店铺
    TeamType string `json:"team_type"`
    // 关闭类型
    // 0:未关闭;
    // 1:过期关闭;
    // 2:标记退款;
    // 3:订单取消;
    // 4:买家取消;
    // 5:卖家取消;
    // 6:部分退款;
    // 10:无法联系上买家;
    // 11:买家误拍或重拍了;
    // 12:买家无诚意完成交易;
    // 13:已通过银行线下汇款;
    // 14:已通过同城见面交易;
    // 15:已通过货到付款交易;
    // 16:已通过网上银行直接汇款;
    // 17:已经缺货无法交易
    CloseType string `json:"close_type"`
    // 物流类型
    // 0:快递发货;
    // 1:到店自提;
    // 2:同城配送;
    // 9:无需发货（虚拟商品订单）
    ExpressType string `json:"express_type"`
    // 订单打标
    OrderTags OrderTags `json:"order_tags"`
    // 订单扩展信息
    OrderExtra OrderExtra `json:"order_extra"`
    // 订单创建时间
    Created string `json:"created"`
    // 订单更新时间
    UpdateTime string `json:"update_time"`
    // 订单过期时间（未付款将自动关单）
    ExpiredTime string `json:"expired_time"`
    // 订单支付时间
    PayTime string `json:"pay_time"`
    // 订单发货时间（当所有商品发货后才会更新）
    ConsignTime string `json:"consign_time"`
    // 订单确认时间（多人拼团成团）
    ConfirmTime string `json:"confirm_time"`
    // 退款状态
    // 0:未退款;
    // 1:部分退款中;
    // 2:部分退款成功;
    // 11:全额退款中;
    // 12:全额退款成功
    RefundState int `json:"refund_state"`
    // 是否零售订单
    IsRetailOrder bool `json:"is_retail_order"`
    // 支付类型。取值范围：
    // WEIXIN (微信自有支付)
    // WEIXIN_DAIXIAO (微信代销支付)
    // ALIPAY (支付宝支付)
    // BANKCARDPAY (银行卡支付)
    // PEERPAY (代付)
    // CODPAY (货到付款)
    // BAIDUPAY (百度钱包支付)
    // PRESENTTAKE (直接领取赠品)
    // COUPONPAY(优惠券/码全额抵扣)
    // BULKPURCHASE(来自分销商的采购)
    // MERGEDPAY(合并付货款)
    // ECARD(有赞E卡支付)
    // PURCHASE_PAY (采购单支付)
    // MARKPAY (标记收款)
    // OFCASH (现金支付)
    // PREPAIDCARD (储值卡余额支付)
    // ENCHASHMENT_GIFT_CARD(礼品卡支付)
    SuccessTime string `json:"success_time"`
    // 网点id
    OfflineId int `json:"offline_id"`
    // 订单成功时间
    PayTypeStr string `json:"pay_type_str"`
}

func (o *OrderInfo) ToCarbon(prop string) (*carbon.Carbon, error) {
    return youzango.GetCarbon(o, prop)
}

// OrderTags
type OrderTags struct {
    // 是否虚拟订单
    IsVirtual bool `json:"is_virtual"`
    // 是否采购单
    IsPurchaseOrder bool `json:"is_purchase_order"`
    // 是否分销单
    IsFenxiaoOrder bool `json:"is_fenxiao_order"`
    // 是否会员订单
    IsMember bool `json:"is_member"`
    // 是否预订单
    IsPreorder bool `json:"is_preorder"`
    // 是否线下订单
    IsOfflineOrder bool `json:"is_offline_order"`
    // 是否多门店订单
    IsMultiStore bool `json:"is_multi_store"`
    // 是否结算
    IsSettle bool `json:"is_settle"`
    // 是否支付
    IsPayed bool `json:"is_payed"`
    // 是否担保交易
    IsSecuredTransactions bool `json:"is_secured_transactions"`
    // 是否享受免邮
    IsPostageFree bool `json:"is_postage_free"`
    // 是否有维权
    IsFeedback bool `json:"is_feedback"`
    // 是否有退款
    IsRefund bool `json:"is_refund"`
    // 是否定金预售
    IsDownPaymentPre bool `json:"is_down_payment_pre"`
}

// OrderExtra
type OrderExtra struct {
    // 是否来自购物车 是：true 不是：false
    IsFromCart string `json:"is_from_cart"`
    // 收银员id
    CashierId string `json:"cashier_id"`
    // 收银员名字
    CashierName string `json:"cashier_name"`
    // 发票抬头
    InvoiceTitle string `json:"invoice_title"`
    // 结算时间
    SettleTime string `json:"settle_time"`
    // 是否父单(分销合并订单) 是：1 其他：null
    IsParentOrder string `json:"is_parent_order"`
    // 是否子单(分销买家订单) 是：1 其他：null
    IsSubOrder string `json:"is_sub_order"`
    // 分销单订单号
    FxOrderNo string `json:"fx_order_no"`
    // 分销店铺id
    FxKdtId string `json:"fx_kdt_id"`
    // 父单号
    ParentOrderNo string `json:"parent_order_no"`
    // 采购单订单号
    PurchaseOrderNo string `json:"purchase_order_no"`
    // 美业分店id
    DeptId string `json:"dept_id"`
    // 下单设备号
    CreateDeviceId string `json:"create_device_id"`
    // 是否是积分订单：1：是 0：不是
    IsPointsOrder string `json:"is_points_order"`
    // 海淘身份证信息：332527XXXXXXXXX
    IdCardNumber string `json:"id_card_number"`
    // 下单人昵称
    BuyerName string `json:"buyer_name"`
    // 是否会员订单
    IsMember string `json:"is_member"`
    // 团购返现优惠金额
    TmCash int `json:"tm_cash"`
    // 团购返现最大返现金额
    TCash int `json:"t_cash"`
    // 订单返现金额
    Cash int `json:"cash"`
    // 虚拟总单号：一次下单发生拆单时，会生成一个虚拟总单号
    OrderCombineId string `json:"order_combine_id"`
    // 拆单时店铺维度的虚拟总单号：发生拆单时，单个店铺生成了多笔订单会生成一个店铺维度的虚拟总单号
    KdtDimensionCombineId string `json:"kdt_dimension_combine_id"`
    // 使用了同一张优惠券&优惠码的多笔订单对应的虚拟总单号
    PromotionCombineId string `json:"promotion_combine_id"`
    // 身份证姓名信息 （订购人的身份证号字段可通过订单详情4.0接口“id_card_number ”获取）
    IdCardName string `json:"id_card_name"`
    // 分销单外部支付流水号
    FxOuterTransactionNo string `json:"fx_outer_transaction_no"`
    // 分销单内部支付流水号
    FxInnerTransactionNo string `json:"fx_inner_transaction_no"`
}

func (o *OrderExtra) ToCarbon(prop string) (*carbon.Carbon, error) {
    return youzango.GetCarbon(o, prop)
}

type SourceInfo struct {
    // 是否来自线下订单
    IsOfflineOrder bool `json:"is_offline_order"`
    // 平台
    Source Source `json:"source"`
    // 订单标记
    // wx_apps:微信小程序买家版
    // wx_shop:微信小程序商家版
    // wx_wm:微信小程序外卖
    // wap_wm:移动端外卖
    // super_store:超级门店
    // weapp_spotlight:新微信小程序买家版
    // wx_meiye:美业小程序
    // wx_apps_maidan:小程序餐饮买单
    // wx_apps_diancan:小程序堂食
    // weapp_youzan:有赞小程序
    // retail_free_buy:零售自由购
    // weapp_owl:知识付费小程序
    // app_spotlight:有赞精选app
    // retail_scan_buy:零售扫码购
    // weapp_plugin:小程序插件 除以上之外为其他
    OrderMark string `json:"order_mark"`
    // 订单唯一识别码
    BookKey string `json:"book_key"`
    // 活动类型：如群团购：”mall_group_buy“
    BizSource string `json:"biz_source"`
}

type Source struct {
    // 平台
    // wx:微信;
    // merchant_3rd:商家自有app;
    // buyer_v:买家版;
    // browser:系统浏览器;
    // alipay:支付宝;
    // qq:腾讯QQ;
    // wb:微博;
    // other:其他
    Platform string `json:"platform"`
    // 微信平台细分
    // wx_gzh:微信公众号;
    // yzdh:有赞大号;
    // merchant_xcx:商家小程序;
    // yzdh_xcx:有赞大号小程序;
    // direct_buy:直接购买
    WxEntrance string `json:"wx_entrance"`
}

type BuyerInfo struct {
    // 买家id
    BuyerId int `json:"buyer_id"`
    // 买家手机号
    BuyerPhone string `json:"buyer_phone"`
    // 粉丝类型 1:自有粉丝; 9:代销粉丝
    FansType int `json:"fans_type"`
    // 粉丝id
    FansId int `json:"fans_id"`
    // 粉丝昵称
    FansNickname string `json:"fans_nickname"`
    // 微信H5和微信小程序（有赞小程序和小程序插件）的订单会返回微信weixin_openid，三方App（有赞APP开店）的订单会返回open_user_id，2019年1月30号后的订单支持返回该参数
    OuterUserId string `json:"outer_user_id"`
}

type PayInfo struct {
    // 优惠前商品总价
    TotalFee string `json:"total_fee"`
    // 邮费
    PostFee string `json:"post_fee"`
    // 最终支付价格 payment=orders.payment的总和
    Payment string `json:"payment"`
    // 有赞支付流水号
    Transaction []string `json:"transaction"`
    // 外部支付单号
    OuterTransactions []string `json:"outer_transactions"`
    // 多阶段支付信息
    PhasePayments []PhasePayment `json:"phase_payments"`
}

func (p *PayInfo) GetPrice(prop string) (int, error) {
    return youzango.GetPrice(p, prop)
}

type PhasePayment struct {
    // 支付阶段
    Phase int `json:"phase"`
    // 支付开始时间
    PayStartTime string `json:"pay_start_time"`
    // 支付结束时间
    PayEndTime string `json:"pay_end_time"`
    // 阶段支付金额
    RealPrice string `json:"real_price"`
    // 外部支付流水号
    OuterTransactionNo string `json:"outer_transaction_no"`
    // 内部支付流水号
    InnerTransactionNo string `json:"inner_transaction_no"`
}

func (p *PhasePayment) GetPrice(prop string) (int, error) {
    return youzango.GetPrice(p, prop)
}

func (p *PhasePayment) ToCarbon(prop string) (*carbon.Carbon, error) {
    return youzango.GetCarbon(p, prop)
}

type RemarkInfo struct {
    // 订单买家留言
    BuyerMessage string `json:"buyer_message"`
    // 订单标星等级 0-5
    Star int `json:"star"`
    // 订单商家备注
    TradeMemo string `json:"trade_memo"`
}

type AddressInfo struct {
    // 收货人姓名
    ReceiverName string `json:"receiver_name"`
    // 收货人手机号
    ReceiverTel string `json:"receiver_tel"`
    // 省
    DeliveryProvince string `json:"delivery_province"`
    // 市
    DeliveryCity string `json:"delivery_city"`
    // 区
    DeliveryDistrict string `json:"delivery_district"`
    // 详细地址
    DeliveryAddress string `json:"delivery_address"`
    // 字段为json格式，需要开发者自行解析
    // lng、lon（经纬度）；
    // checkOutTime（酒店退房时间）；
    // recipients（入住人）；
    // checkInTime（酒店入住时间）；
    // idCardNumber（海淘身份证信息）；
    // areaCode（邮政编码）
    AddressExtra string `json:"address_extra"`
    // 邮政编码
    DeliveryPostalCode string `json:"delivery_postal_code"`
    // 到店自提信息 json格式
    SelfFetchInfo string `json:"self_fetch_info"`
    // 同城送预计送达时间-开始时间 非同城送以及没有开启定时达的订单不返回
    DeliveryStartTime string `json:"delivery_start_time"`
    // 同城送预计送达时间-结束时间 非同城送以及没有开启定时达的订单不返回
    DeliveryEndTime string `json:"delivery_end_time"`
}

func (a *AddressInfo) ToCarbon(prop string) (*carbon.Carbon, error) {
    return youzango.GetCarbon(a, prop)
}

type Order struct {
    // 订单明细id
    Oid string `json:"oid"`
    // 订单类型
    // 0:普通类型商品;
    // 1:拍卖商品;
    // 5:餐饮商品;
    // 10:分销商品;
    // 20:会员卡商品;
    // 21:礼品卡商品;
    // 23:有赞会议商品;
    // 24:周期购;
    // 30:收银台商品;
    // 31:知识付费商品;
    // 35:酒店商品;
    // 40:普通服务类商品;
    // 182:普通虚拟商品;
    // 183:电子卡券商品;
    // 201:外部会员卡商品;
    // 202:外部直接收款商品;
    // 203:外部普通商品;
    // 205:mock不存在商品;
    // 206:小程序二维码
    ItemType int `json:"item_type"`
    // 商品名称
    Title string `json:"title"`
    // 商品数量
    Num int `json:"num"`
    // 商家编码
    OuterSkuId string `json:"outer_sku_id"`
    // 商品留言
    BuyerMessages string `json:"buyer_messages"`
    // 单商品原价
    Price string `json:"price"`
    // 商品优惠后总价
    TotalFee string `json:"total_fee"`
    // 商品最终均摊价
    Payment string `json:"payment"`
    // 商品id
    ItemId int `json:"item_id"`
    // 规格id（无规格商品为0）
    SkuId int `json:"sku_id"`
    // 规格信息（无规格商品为空）
    SkuPropertiesName string `json:"sku_properties_name"`
    // 商品编码
    OuterItemId string `json:"outer_item_id"`
    // 商品积分价（非积分商品则为0）
    PointsPrice string `json:"points_price"`
    // 商品图片地址
    PicPath string `json:"pic_path"`
    // 商品详情链接
    GoodsUrl string `json:"goods_url"`
    // 商品别名
    Alias string `json:"alias"`
    // 是否赠品
    IsPresent bool `json:"is_present"`
    // 单商品现价，减去了商品的优惠金额
    DiscountPrice string `json:"discount_price"`
    // 商品唯一编码
    SkuUniqueCode string `json:"sku_unique_code"`
    // 0 全款预售，1 定金预售
    PreSaleType string `json:"pre_sale_type"`
    // 是否为预售商品 如果是预售商品则为1
    IsPreSale string `json:"is_pre_sale"`
    // 是否是跨境海淘订单("1":是,"0":否)
    IsCrossBorder string `json:"is_cross_border"`
    // 海关编号
    CustomsCode string `json:"customs_code"`
    // 海淘商品贸易模式
    CrossBorderTradeMode string `json:"cross_border_trade_mode"`
    // 子订单号
    SubOrderNo string `json:"sub_order_no"`
    // 分销单金额 ，单位元
    FenxiaoPrice string `json:"fenxiao_price"`
    // 分销单实付金额 ，单位元
    FenxiaoPayment string `json:"fenxiao_payment"`
}

func (o *Order) ToCarbon(prop string) (*carbon.Carbon, error) {
    return youzango.GetCarbon(o, prop)
}

type ChildInfo struct {
    // 	送礼编号
    GiftNo string `json:"gift_no"`
    // 送礼标记
    GiftSign string `json:"gift_sign"`
    // 子单详情
    ChildOrders []ChildOrder `json:"child_orders"`
}

type ChildOrder struct {
    // 订单号
    Tid string `json:"tid"`
    // 领取人姓名
    UserName string `json:"user_name"`
    // 领取人电话
    UserTel string `json:"user_tel"`
    // 省
    Province string `json:"province"`
    // 市
    City string `json:"city"`
    // 区
    County string `json:"county"`
    // 收货地址详情
    AddressDetail string `json:"address_detail"`
    // 老送礼订单状态：WAIT_EXPRESS(5, "待发货"), EXPRESS(6, "已发货"), SUCCESS(100, "成功")
    OrderState string `json:"order_state"`
}

type RefundOrder struct {
    // 退款类型 1:退款 - 买家申请退款; 2:退款 - 商家主动退款; 3:退款 - 一键退款
    RefundType int `json:"refund_type"`
    // 退款金额
    RefundFee string `json:"refund_fee"`
    // 退款id
    RefundId string `json:"refund_id"`
    // 退款状态
    // 1:买家已经申请退款，等待卖家同意;
    // 10:卖家拒绝退款;
    // 20:卖家已经同意退货，等待买家退货;
    // 30:买家已经退货，等待卖家确认收货;
    // 40:卖家未收到货,拒绝退款;
    // 50:退款关闭;
    // 60:退款成功
    RefundState int `json:"refund_state"`
    // 退款交易明细
    Oids []Oid `json:"oids"`
}

type Oid struct {
    // 交易明细id
    Oid string `json:"oid"`
}

type DeliveryOrder struct {
    // 改字段已弃用 包裹id已移至dists中的dist_id字段
    PkId int `json:"pk_id"`
    // 物流状态 0:待发货; 1:已发货
    ExpressState int `json:"express_state"`
    // 物流类型 0:手动发货; 1:系统自动发货
    ExpressType int `json:"express_type"`
    // 发货明细
    Oids []Oid `json:"oids"`
    // 包裹信息
    Dists []Dist `json:"dists"`
}

type Dist struct {
    // 包裹id
    DistId string `json:"dist_id"`
    // 包裹详情
    ExpressInfo ExpressInfo `json:"express_info"`
}

type ExpressInfo struct {
    // 物流类型
    ExpressId int `json:"express_id"`
    // 物流编号
    ExpressNo string `json:"express_no"`
}

type OrderPromotion struct {
    ExpressInfo
}
