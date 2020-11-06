package bean

//type -100402: invoice
type HtlcRequestInvoice struct {
	NetType string `json:"net_type"` //解析用
	HtlcRequestFindPathInfo
	typeLengthValue
}

type HtlcRequestFindPathInfo struct {
	RecipientNodePeerId string   `json:"recipient_node_peer_id"`
	RecipientUserPeerId string   `json:"recipient_user_peer_id"`
	H                   string   `json:"h"`
	ExpiryTime          JsonDate `json:"expiry_time"`
	PropertyId          int64    `json:"property_id"`
	Amount              float64  `json:"amount"`
	Description         string   `json:"description"`
	IsPrivate           bool     `json:"is_private"`
}

//type --100401: alice tell carl ,she wanna transfer some money to Carl
type HtlcRequestFindPath struct {
	Invoice string `json:"invoice"`
	HtlcRequestFindPathInfo
	typeLengthValue
}

//type 消息 --100040: Alice请求创建htlc交易
type CreateHtlcTxForC3a struct {
	Amount                           float64 `json:"amount"`
	Memo                             string  `json:"memo"`
	H                                string  `json:"h"`
	CltvExpiry                       int     `json:"cltv_expiry"` //发起者设定的总的等待的区块个数
	RoutingPacket                    string  `json:"routing_packet"`
	LastTempAddressPrivateKey        string  `json:"last_temp_address_private_key"` //	上个RSMC委托交易用到的临时地址的私钥
	CurrRsmcTempAddressIndex         int     `json:"curr_rsmc_temp_address_index"`
	CurrRsmcTempAddressPubKey        string  `json:"curr_rsmc_temp_address_pub_key"` //	创建Cnx中的toRsmc的部分使用的临时地址的公钥
	CurrHtlcTempAddressIndex         int     `json:"curr_htlc_temp_address_index"`
	CurrHtlcTempAddressPubKey        string  `json:"curr_htlc_temp_address_pub_key"` //	创建Cnx中的toHtlc的部分使用的临时地址的公钥
	CurrHtlcTempAddressForHt1aIndex  int     `json:"curr_htlc_temp_address_for_ht1a_index"`
	CurrHtlcTempAddressForHt1aPubKey string  `json:"curr_htlc_temp_address_for_ht1a_pub_key"` //	创建Ht1a中生成ht1a的输出的Rmsc的临时地址的公钥
	typeLengthValue
}

//type 响应 --100040: 需要alice签名C3a的交易数据
type CreateHtlcTxForC3aResult struct {
	ChannelId              string                  `json:"channel_id"` //the global channel id.
	C3aRsmcRawData         NeedClientSignRawTxData `json:"c3a_rsmc_raw_data"`
	C3aCounterpartyRawData NeedClientSignRawTxData `json:"c3a_counterparty_raw_data"`
	C3aHtlcRawData         NeedClientSignRawTxData `json:"c3a_htlc_raw_data"`
}

//type 消息 --100100: 需要alice签名C3a的交易数据
type AliceSignedHtlcDataForC3a struct {
	ChannelId                       string `json:"channel_id"` //the global channel id.
	C3aRsmcPartialSignedHex         string `json:"c3a_rsmc_partial_signed_hex"`
	C3aCounterpartyPartialSignedHex string `json:"c3a_counterparty_partial_signed_hex"`
	C3aHtlcPartialSignedHex         string `json:"c3a_htlc_partial_signed_hex"`
	typeLengthValue
}

//type p2p消息 --40 Alice新增htlc交易C3a的请求，p2p推给bob
type CreateHtlcTxForC3aOfP2p struct {
	ChannelId                        string                  `json:"channel_id"` //the global channel id.
	C3aCounterpartyPartialSignedData NeedClientSignRawTxData `json:"c3a_counterparty_partial_signed_data"`
	C3aRsmcPartialSignedData         NeedClientSignRawTxData `json:"c3a_rsmc_partial_signed_data"`
	C3aHtlcPartialSignedData         NeedClientSignRawTxData `json:"c3a_htlc_partial_signed_data"`
	CreateHtlcTxForC3a
}

//type obd推送 --110040 obd主动推送alice的C3a的信息给bob的客户端
type CreateHtlcTxForC3aToBob struct {
	ChannelId                        string                  `json:"channel_id"` //the global channel id.
	C3aCounterpartyPartialSignedData NeedClientSignRawTxData `json:"c3a_counterparty_partial_signed_data"`
	C3aRsmcPartialSignedData         NeedClientSignRawTxData `json:"c3a_rsmc_partial_signed_data"`
	C3aHtlcPartialSignedData         NeedClientSignRawTxData `json:"c3a_htlc_partial_signed_data"`
	Amount                           float64                 `json:"amount"`
	Memo                             string                  `json:"memo"`
}

//type 消息 --100041 bob签名C3a的结果
type BobSignedC3a struct {
	ChannelId                        string `json:"channel_id"` //the global channel id.
	C3aCompleteSignedRsmcHex         string `json:"c3a_complete_signed_rsmc_hex"`
	C3aCompleteSignedCounterpartyHex string `json:"c3a_complete_signed_counterparty_hex"`
	C3aCompleteSignedHtlcHex         string `json:"c3a_complete_signed_htlc_hex"`
	LastTempAddressPrivateKey        string `json:"last_temp_address_private_key"` //	上个RSMC委托交易用到的临时私钥
	CurrRsmcTempAddressIndex         int    `json:"curr_rsmc_temp_address_index"`
	CurrRsmcTempAddressPubKey        string `json:"curr_rsmc_temp_address_pub_key"` //	创建Cnx中的toRsmc的部分使用的临时地址的公钥
	CurrHtlcTempAddressIndex         int    `json:"curr_htlc_temp_address_index"`
	CurrHtlcTempAddressPubKey        string `json:"curr_htlc_temp_address_pub_key"` //	创建Cnx中的toHtlc的部分使用的临时地址的公钥
	CurrHtlcTempAddressForHeIndex    int    `json:"curr_htlc_temp_address_for_he_index"`
	CurrHtlcTempAddressForHePubKey   string `json:"curr_htlc_temp_address_for_he_pub_key"` // These keys of HE1b output. Example Bob send R to Alice, these is Bob3's.
}

//type 消息 --100101 需要bob签名C3b，及C3a的toRsmc的Rd和Br，toHtlc的Br，Ht1a，Hlock
type NeedBobSignHtlcTxOfC3b struct {
	ChannelId              string                    `json:"channel_id"` //the global channel id.
	C3aRsmcRdRawData       NeedClientSignRawTxData   `json:"c3a_rsmc_rd_raw_data"`
	C3aRsmcBrRawData       NeedClientSignRawBRTxData `json:"c3a_rsmc_br_raw_data"`
	C3aHtlcHtRawData       NeedClientSignRawTxData   `json:"c3a_htlc_ht_raw_data"`
	C3aHtlcHlockRawData    NeedClientSignRawTxData   `json:"c3a_htlc_hlock_raw_data"`
	C3aHtlcBrRawData       NeedClientSignRawBRTxData `json:"c3a_htlc_br_raw_data"`
	C3bRsmcRawData         NeedClientSignRawTxData   `json:"c3b_rsmc_raw_data"`
	C3bCounterpartyRawData NeedClientSignRawTxData   `json:"c3b_counterparty_raw_data"`
	C3bHtlcRawData         NeedClientSignRawTxData   `json:"c3b_htlc_raw_data"`
}

//type 返回值 --100101
type NeedBobSignHtlcTxOfC3bResult struct {
	ChannelId                       string `json:"channel_id"` //the global channel id.
	C3aRsmcRdPartialSignedHex       string `json:"c3a_rsmc_rd_partial_signed_hex"`
	C3aRsmcBrPartialSignedHex       string `json:"c3a_rsmc_br_partial_signed_hex"`
	C3aHtlcHtPartialSignedHex       string `json:"c3a_htlc_ht_partial_signed_hex"`
	C3aHtlcHlockPartialSignedHex    string `json:"c3a_htlc_hlock_partial_signed_hex"`
	C3aHtlcBrPartialSignedHex       string `json:"c3a_htlc_br_partial_signed_hex"`
	C3bRsmcPartialSignedHex         string `json:"c3b_rsmc_partial_signed_hex"`
	C3bCounterpartyPartialSignedHex string `json:"c3b_counterparty_partial_signed_hex"`
	C3bHtlcPartialSignedHex         string `json:"c3b_htlc_partial_signed_hex"`
}

//type p2p消息 --41 推送bob完成C3a的完整签名，C3a的子交易的部分签名，C3b的部分签名给alice进行二签和更新保存
type NeedAliceSignHtlcTxOfC3bP2p struct {
	ChannelId                        string                  `json:"channel_id"` //the global channel id.
	C3aCompleteSignedRsmcHex         string                  `json:"c3a_complete_signed_rsmc_hex"`
	C3aCompleteSignedCounterpartyHex string                  `json:"c3a_complete_signed_counterparty_hex"`
	C3aCompleteSignedHtlcHex         string                  `json:"c3a_complete_signed_htlc_hex"`
	C3aRsmcRdPartialSignedData       NeedClientSignRawTxData `json:"c3a_rsmc_rd_partial_signed_data"`
	C3aHtlcHtPartialSignedData       NeedClientSignRawTxData `json:"c3a_htlc_ht_partial_signed_data"`
	C3aHtlcHlockPartialSignedData    NeedClientSignRawTxData `json:"c3a_htlc_hlock_partial_signed_data"`
	C3bRsmcPartialSignedData         NeedClientSignRawTxData `json:"c3b_rsmc_partial_signed_data"`
	C3bCounterpartyPartialSignedData NeedClientSignRawTxData `json:"c3b_counterparty_partial_signed_data"`
	C3bHtlcPartialSignedData         NeedClientSignRawTxData `json:"c3b_htlc_partial_signed_data"`
}

//type 消息 --110041 把41的消息推送给Alice
type NeedAliceSignHtlcTxOfC3b struct {
	ChannelId                        string                  `json:"channel_id"` //the global channel id.
	C3aRsmcRdPartialSignedData       NeedClientSignRawTxData `json:"c3a_rsmc_rd_partial_signed_data"`
	C3aHtlcHtPartialSignedData       NeedClientSignRawTxData `json:"c3a_htlc_ht_partial_signed_data"`
	C3aHtlcHlockPartialSignedData    NeedClientSignRawTxData `json:"c3a_htlc_hlock_partial_signed_data"`
	C3bRsmcPartialSignedData         NeedClientSignRawTxData `json:"c3b_rsmc_partial_signed_data"`
	C3bCounterpartyPartialSignedData NeedClientSignRawTxData `json:"c3b_counterparty_partial_signed_data"`
	C3bHtlcPartialSignedData         NeedClientSignRawTxData `json:"c3b_htlc_partial_signed_data"`
}

//type 消息 --100102 alice对c3b的签名结果
type AliceSignedHtlcTxOfC3bResult struct {
	ChannelId                        string `json:"channel_id"` //the global channel id.
	C3aRsmcRdCompleteSignedHex       string `json:"c3a_rsmc_rd_complete_signed_hex"`
	C3aHtlcHtCompleteSignedHex       string `json:"c3a_htlc_ht_complete_signed_hex"`
	C3aHtlcHlockCompleteSignedHex    string `json:"c3a_htlc_hlock_complete_signed_hex"`
	C3bRsmcCompleteSignedHex         string `json:"c3b_rsmc_complete_signed_hex"`
	C3bCounterpartyCompleteSignedHex string `json:"c3b_counterparty_complete_signed_hex"`
	C3bHtlcCompleteSignedHex         string `json:"c3b_htlc_complete_signed_hex"`
}

//type 响应 --100102 继续子交易的签名
type NeedAliceSignHtlcSubTxOfC3b struct {
	ChannelId           string                    `json:"channel_id"` //the global channel id.
	C3aHtlcHtrdRawData  NeedClientSignRawTxData   `json:"c3a_htlc_htrd_raw_data"`
	C3aHtlcHtbrRawData  NeedClientSignRawBRTxData `json:"c3a_htlc_htbr_raw_data"`
	C3bRsmcRdRawData    NeedClientSignRawTxData   `json:"c3b_rsmc_rd_raw_data"`
	C3bRsmcBrRawData    NeedClientSignRawBRTxData `json:"c3b_rsmc_br_raw_data"`
	C3bHtlcHtdRawData   NeedClientSignRawTxData   `json:"c3b_htlc_htd_raw_data"`
	C3bHtlcHlockRawData NeedClientSignRawTxData   `json:"c3b_htlc_hlock_raw_data"`
	C3bHtlcBrRawData    NeedClientSignRawBRTxData `json:"c3b_htlc_br_raw_data"`
}

//type 消息 --100103
type AliceSignHtlcSubTxOfC3bResult struct {
	ChannelId                    string `json:"channel_id"` //the global channel id.
	C3aHtlcHtrdPartialSignedHex  string `json:"c3a_htlc_htrd_partial_signed_hex"`
	C3aHtlcHtbrPartialSignedHex  string `json:"c3a_htlc_htbr_partial_signed_hex"`
	C3bRsmcRdPartialSignedHex    string `json:"c3b_rsmc_rd_partial_signed_hex"`
	C3bRsmcBrPartialSignedHex    string `json:"c3b_rsmc_br_partial_signed_hex"`
	C3bHtlcHtdPartialSignedHex   string `json:"c3b_htlc_htd_partial_signed_hex"`
	C3bHtlcHlockPartialSignedHex string `json:"c3b_htlc_hlock_partial_signed_hex"`
	C3bHtlcBrPartialSignedHex    string `json:"c3b_htlc_br_partial_signed_hex"`
}

// type p2p消息 42 Alice对c3b完成签名，把结果通过p2p推送给bob所在的obd
type NeedBobSignHtlcSubTxOfC3bP2p struct {
	ChannelId                        string                    `json:"channel_id"` //the global channel id.
	C3bCompleteSignedRsmcHex         string                    `json:"c3b_complete_signed_rsmc_hex"`
	C3bCompleteSignedCounterpartyHex string                    `json:"c3b_complete_signed_counterparty_hex"`
	C3bCompleteSignedHtlcHex         string                    `json:"c3b_complete_signed_htlc_hex"`
	C3bRsmcRdPartialData             NeedClientSignRawTxData   `json:"c3b_rsmc_rd_partial_data"`
	C3bHtlcHtdPartialData            NeedClientSignRawTxData   `json:"c3b_htlc_htd_partial_data"`
	C3bHtlcHlockPartialData          NeedClientSignRawTxData   `json:"c3b_htlc_hlock_partial_data"`
	C3aHtlcHtrdPartialData           NeedClientSignRawTxData   `json:"c3a_htlc_htrd_partial_data"`
	C3aHtlcHtbrPartialData           NeedClientSignRawBRTxData `json:"c3a_htlc_htbr_partial_data"`
	C3aHtlcHedRawData                NeedClientSignRawTxData   `json:"c3a_htlc_hed_raw_data"`
}

// type 110042 需要bob签名C3b的子交易及C3a的ht的子交易
type NeedBobSignHtlcSubTxOfC3b struct {
	ChannelId               string                  `json:"channel_id"` //the global channel id.
	C3aHtlcHtrdPartialData  NeedClientSignRawTxData `json:"c3a_htlc_htrd_partial_data"`
	C3aHtlcHedRawData       NeedClientSignRawTxData `json:"c3a_htlc_hed_raw_data"`
	C3bRsmcRdPartialData    NeedClientSignRawTxData `json:"c3b_rsmc_rd_partial_data"`
	C3bHtlcHtdPartialData   NeedClientSignRawTxData `json:"c3b_htlc_htd_partial_data"`
	C3bHtlcHlockPartialData NeedClientSignRawTxData `json:"c3b_htlc_hlock_partial_data"`
}

// type 消息 100104 bob完成签名C3b的子交易及C3a的ht的子交易
type BobSignedHtlcSubTxOfC3b struct {
	ChannelId                     string `json:"channel_id"` //the global channel id.
	C3aHtlcHtrdCompleteSignedHex  string `json:"c3a_htlc_htrd_complete_signed_hex"`
	C3aHtlcHedPartialSignedHex    string `json:"c3a_htlc_hed_partial_signed_hex"`
	C3bRsmcRdCompleteSignedHex    string `json:"c3b_rsmc_rd_complete_signed_hex"`
	C3bHtlcHtdCompleteSignedHex   string `json:"c3b_htlc_htd_complete_signed_hex"`
	C3bHtlcHlockCompleteSignedHex string `json:"c3b_htlc_hlock_complete_signed_hex"`
}

// type p2p消息 43 bob完成对C3b的签名，把C3a的htrd和hed签名结果，以及C3b的Hlock的子交易He裸交易传递给alice签名
type NeedAliceSignHeSubTxOfC3bP2p struct {
	ChannelId                    string                  `json:"channel_id"` //the global channel id.
	C3aHtlcHtrdCompleteSignedHex string                  `json:"c3a_htlc_htrd_complete_signed_hex"`
	C3aHtlcHedPartialData        NeedClientSignRawTxData `json:"c3a_htlc_hed_partial_data"` //等待R的签名
	C3bHtlcHeRawData             NeedClientSignRawTxData `json:"c3b_htlc_he_raw_data"`      // Hlock的子交易
}

// type obd推送消息 110043 需要alice签名C3b的he
type NeedAliceSignHeSubTxOfC3b struct {
	ChannelId        string                  `json:"channel_id"`             //the global channel id.
	C3bHtlcHeRawData NeedClientSignRawTxData `json:"c_3_b_htlc_he_raw_data"` // Hlock的之交易
}

// type 消息 100105 Alice完成对c3b的He的签名
type AliceSignHeSubTxOfC3b struct {
	ChannelId                 string `json:"channel_id"` //the global channel id.
	C3bHtlcHePartialSignedHex string `json:"c3b_htlc_he_partial_signed_hex"`
}

// type 响应 100105
type CreateHtlcC3aResult struct {
	ChannelId string `json:"channel_id"` //the global channel id.
}

// type p2p消息 44
type AliceSignedHeSubTxOfC3bP2p struct {
	ChannelId                  string                  `json:"channel_id"`                      //the global channel id.
	C3bHtlcHePartialSignedData NeedClientSignRawTxData `json:"c3b_htlc_he_partial_signed_data"` //等待R的签名
}

// type obd推送消息 100044
type CreateHtlcC3bResult struct {
}

// 正向H传递完成

// 开始反向R的传递

// type 消息 100045
type HtlcBobSendR struct {
	ChannelId string `json:"channel_id"`
	R         string `json:"r"`
	typeLengthValue
}

// type 响应 100045
type HtlcBobSendRResult struct {
	ChannelId          string                    `json:"channel_id"`
	C3bHtlcHerdRawData NeedClientSignRawTxData   `json:"c3b_htlc_herd_raw_data"`
	C3bHtlcHebrRawData NeedClientSignRawBRTxData `json:"c3b_htlc_hebr_raw_data"`
}

// type 消息 100106
type BobSignHerdAndHebrForC3b struct {
	ChannelId                   string `json:"channel_id"`
	C3bHtlcHerdPartialSignedHex string `json:"c3b_htlc_herd_partial_signed_hex"`
	C3bHtlcHebrPartialSignedHex string `json:"c3b_htlc_hebr_partial_signed_hex"`
}

// type p2p消息 45
type NeedAliceSignHerdTxOfC3bP2p struct {
	ChannelId                    string                    `json:"channel_id"` //the global channel id.
	C3bHtlcHerdPartialSignedData NeedClientSignRawTxData   `json:"c3b_htlc_herd_partial_signed_data"`
	C3bHtlcHebrPartialSignedData NeedClientSignRawBRTxData `json:"c3b_htlc_hebr_partial_signed_data"`
}

// type obd推送消息 110045
type NeedAliceSignHerdTxOfC3b struct {
	ChannelId                    string                  `json:"channel_id"` //the global channel id.
	C3bHtlcHerdPartialSignedData NeedClientSignRawTxData `json:"c3b_htlc_herd_partial_signed_data"`
}

// type 消息 100107
type AliceSignHerdTxOfC3a struct {
	ChannelId                    string `json:"channel_id"` //the global channel id.
	C3bHtlcHerdCompleteSignedHex string `json:"c3b_htlc_herd_complete_signed_hex"`
}

// type 响应 100107 反向R传递的最终结果
type AfterVerifyROfC3a struct {
	ChannelId string `json:"channel_id"` //the global channel id.
}

// type p2p消息 46 Alice啊我那次herd的签名，发送结果给bob所在的obd
type AliceSignedHerdTxOfC3bP2p struct {
	ChannelId                    string `json:"channel_id"` //the global channel id.
	C3bHtlcHerdCompleteSignedHex string `json:"c3b_htlc_herd_complete_signed_hex"`
}

// type 响应 100046 反向R传递的最终结果
type AliceSignHerdTxOfC3b struct {
	ChannelId string `json:"channel_id"` //the global channel id.
}

//
//
//
//
//
//
//
//
//
//

// type 40 payer start htlc tx
type AddHtlcRequest struct {
	PropertyId                           int64   `json:"property_id"`
	Amount                               float64 `json:"amount"`
	Memo                                 string  `json:"memo"`
	H                                    string  `json:"h"`
	CltvExpiry                           int     `json:"cltv_expiry"` //发起者设定的总的等待的区块个数
	RoutingPacket                        string  `json:"routing_packet"`
	ChannelAddressPrivateKey             string  `json:"channel_address_private_key"`   //	开通通道用到的地址的私钥
	LastTempAddressPrivateKey            string  `json:"last_temp_address_private_key"` //	上个RSMC委托交易用到的临时地址的私钥
	CurrRsmcTempAddressIndex             int     `json:"curr_rsmc_temp_address_index"`
	CurrRsmcTempAddressPubKey            string  `json:"curr_rsmc_temp_address_pub_key"`     //	创建Cnx中的toRsmc的部分使用的临时地址的公钥
	CurrRsmcTempAddressPrivateKey        string  `json:"curr_rsmc_temp_address_private_key"` //	创建Cnx中的toRsmc的部分使用的临时地址的私钥
	CurrHtlcTempAddressIndex             int     `json:"curr_htlc_temp_address_index"`
	CurrHtlcTempAddressPubKey            string  `json:"curr_htlc_temp_address_pub_key"`     //	创建Cnx中的toHtlc的部分使用的临时地址的公钥
	CurrHtlcTempAddressPrivateKey        string  `json:"curr_htlc_temp_address_private_key"` //	创建Cnx中的toHtlc的部分使用的临时地址的私钥
	CurrHtlcTempAddressForHt1aIndex      int     `json:"curr_htlc_temp_address_for_ht1a_index"`
	CurrHtlcTempAddressForHt1aPubKey     string  `json:"curr_htlc_temp_address_for_ht1a_pub_key"`     //	创建Ht1a中生成ht1a的输出的Rmsc的临时地址的公钥
	CurrHtlcTempAddressForHt1aPrivateKey string  `json:"curr_htlc_temp_address_for_ht1a_private_key"` //	创建Ht1a中生成ht1a的输出的Rmsc的临时地址的私钥
	typeLengthValue
}

//type -100041: bob sign the request for the interNode
type HtlcSignGetH struct {
	PayerCommitmentTxHash         string `json:"payer_commitment_tx_hash"`
	ChannelAddressPrivateKey      string `json:"channel_address_private_key"`   //	开通通道用到的私钥
	LastTempAddressPrivateKey     string `json:"last_temp_address_private_key"` //	上个RSMC委托交易用到的临时私钥
	CurrRsmcTempAddressIndex      int    `json:"curr_rsmc_temp_address_index"`
	CurrRsmcTempAddressPubKey     string `json:"curr_rsmc_temp_address_pub_key"`     //	创建Cnx中的toRsmc的部分使用的临时地址的公钥
	CurrRsmcTempAddressPrivateKey string `json:"curr_rsmc_temp_address_private_key"` //	创建Cnx中的toRsmc的部分使用的临时地址的私钥
	CurrHtlcTempAddressIndex      int    `json:"curr_htlc_temp_address_index"`
	CurrHtlcTempAddressPubKey     string `json:"curr_htlc_temp_address_pub_key"`     //	创建Cnx中的toHtlc的部分使用的临时地址的公钥
	CurrHtlcTempAddressPrivateKey string `json:"curr_htlc_temp_address_private_key"` //	创建Cnx中的toHtlc的部分使用的临时地址的私钥
	typeLengthValue
}

// -42 msg
type AfterBobSignAddHtlcToAlice struct {
	ChannelId                      string `json:"channel_id"`
	PayerCommitmentTxHash          string `json:"payer_commitment_tx_hash"`
	PayerSignedRsmcHex             string `json:"payer_signed_rsmc_hex"`
	PayerSignedToCounterpartyHex   string `json:"payer_signed_to_counterparty_hex"`
	PayerSignedHtlcHex             string `json:"payer_signed_htlc_hex"`
	PayerRsmcRdHex                 string `json:"payer_rsmc_rd_hex"`
	PayerLockByHForBobHex          string `json:"payer_lock_by_h_for_bob_hex"`
	PayerHt1aHex                   string `json:"payer_ht_1_a_hex"`
	PayeeLastTempAddressPrivateKey string `json:"payee_last_temp_address_private_key"`
	PayeeCurrRsmcTempAddressPubKey string `json:"payee_curr_rsmc_temp_address_pub_key"`
	PayeeCurrHtlcTempAddressPubKey string `json:"payee_curr_htlc_temp_address_pub_key"`
	PayeeCommitmentTxHash          string `json:"payee_commitment_tx_hash"`
	PayeeRsmcHex                   string `json:"payee_rsmc_hex"`
	PayeeToCounterpartyTxHex       string `json:"payee_to_counterparty_tx_hex"`
	PayeeHtlcHex                   string `json:"payee_htlc_hex"`
}

// -43 付款人签名收款人的承诺交易的三个hex及创建对应的子交易
type AfterAliceSignAddHtlcToBob struct {
	PayerCommitmentTxHash                 string `json:"payer_commitment_tx_hash"`
	PayerCurrHtlcTempAddressForHt1aPubKey string `json:"payer_curr_htlc_temp_address_for_ht1a_pub_key"`
	PayerHt1aSignedHex                    string `json:"payer_ht1a_signed_hex"`
	PayeeCommitmentTxHash                 string `json:"payee_commitment_tx_hash"`
	PayeeSignedRsmcHex                    string `json:"payee_signed_rsmc_hex"`
	PayeeRsmcRdHex                        string `json:"payee_rsmc_rd_hex"`
	PayeeSignedToCounterpartyHex          string `json:"payee_signed_to_counterparty_hex"`
	PayeeSignedHtlcHex                    string `json:"payee_signed_htlc_hex"`
	PayeeHtdHex                           string `json:"payee_htd_hex"`
	PayeeHlockHex                         string `json:"payee_hlock_hex"`
}

// -44 收款人更加签名后的ht1a，创建这个交易的RD
type PayeeCreateHt1aRDForPayer struct {
	PayerCommitmentTxHash string `json:"payer_commitment_tx_hash"`
	PayerHt1aRDHex        string `json:"payer_ht1a_rd_hex"`
}

//type -45: Send R to previous node. and create commitment transactions.
type HtlcSendR struct {
	ChannelId                            string `json:"channel_id"`
	R                                    string `json:"r"`
	ChannelAddressPrivateKey             string `json:"channel_address_private_key"` // The key of Sender. Example Bob send R to Alice, the Sender is Bob.
	CurrHtlcTempAddressForHE1bIndex      int    `json:"curr_htlc_temp_address_for_he1b_index"`
	CurrHtlcTempAddressForHE1bPubKey     string `json:"curr_htlc_temp_address_for_he1b_pub_key"` // These keys of HE1b output. Example Bob send R to Alice, these is Bob3's.
	CurrHtlcTempAddressForHE1bPrivateKey string `json:"curr_htlc_temp_address_for_he1b_private_key"`
	typeLengthValue
}

//type -46: Middleman node check out if R is correct
type HtlcCheckRAndCreateTx struct {
	ChannelId                string `json:"channel_id"`
	R                        string `json:"r"`
	MsgHash                  string `json:"msg_hash"`
	ChannelAddressPrivateKey string `json:"channel_address_private_key"` // The key of creator tx. Example Bob send R to Alice, that is Alice's.
	typeLengthValue
}

// -47
type HtlcRPayerVerifyRInfoToPayee struct {
	ChannelId            string `json:"channel_id"`
	PayerHlockTxHex      string `json:"payer_hlock_tx_hex"`
	PayerHed1aHex        string `json:"payer_hed1a_hex"`
	PayeeSignedHerd1bHex string `json:"payee_signed_herd1b_hex"`
}

// -48
type HtlcRPayeeSignHed1aToPayer struct {
	ChannelId           string `json:"channel_id"`
	PayerSignedHed1aHex string `json:"payer_signed_hed1a_hex"`
}