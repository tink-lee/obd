package omnicore

import (
	"github.com/omnilaboratory/obd/bean"
	"github.com/omnilaboratory/obd/tool"
	"log"
	"testing"
)

func TestVerifyOmniTxHex(t *testing.T) {
	hex1 := "02000000019a66e2388b670aba78f8b11a2cf663a52233b0dc304be57347889b5b64702ad50000000000ffffffff033c1b00000000000017a914c57dc754b4ea659897b2f17254703d6e669568b6870000000000000000166a146f6d6e6900000000000000890000000010089d40220200000000000017a914c57dc754b4ea659897b2f17254703d6e669568b68700000000"
	transaction := DecodeRawTransaction(hex1, tool.GetCoreNet())
	log.Println(transaction)
	hex2 := "02000000019a66e2388b670aba78f8b11a2cf663a52233b0dc304be57347889b5b64702ad500000000db00483045022100b681bbf4524ab2237644dadc446221671f5746a8c33bd18852fab5f8256f2bb002202e7de1905670918de003119b2a6644e96a3563ab387e8e5999adcae04483411f01483045022100d4fd6bfca414f3a29ff0c08f63de82cd4fb84d7feada82dc721511f21703e7be0220134b0b6d6706d0e39d03e0523ee733a48ff9b5fb4b248d7f4e2b8f4461eda55b014752210238b21ae976c7059057b2b973fec5be8c7b37cea1b15aff52ecf309d5f527e76e2103c9b8dbcf3940eee1445686aaff6f6acf5eb5cd6490ee868aab13312a8fa8f7be52aeffffffff033c1b00000000000017a914c57dc754b4ea659897b2f17254703d6e669568b6870000000000000000166a146f6d6e6900000000000000890000000010089d40220200000000000017a914c57dc754b4ea659897b2f17254703d6e669568b68700000000"
	rawTransaction := DecodeRawTransaction(hex2, tool.GetCoreNet())
	log.Println(rawTransaction)
	txHex := VerifySignatureFromTxHex(hex1, hex2)
	log.Println(txHex)
}

func TestSign(t *testing.T) {
	redeemhex := "0200000002ac5a7c14bf3a63944b50333e12263213fdf59eacfc1a60af27298324db736aea00000000930000483045022100c4acce9704328d6ffd85e1c4cd358b999c9590adff9db3ddf24a388f282cc8cf022024660b7321d7502e4ca71b989cda642d676714c400a6bbc4b5f2dd205f512d0c0147522103d2586577e2f4460b1c299f4f74b719982c1982f31e48ca1ef406347472fa611221038097033cb34a88b8bfc052adbbfefa8e92c33b7635e30c7a79d90ff4917c6c0b52aeffffffffac5a7c14bf3a63944b50333e12263213fdf59eacfc1a60af27298324db736aea020000009200004730440220701fde05a432c78f541aaa4e10a714e79a60daa3d764d0830ca2195d7d85f325022053cf5f59dad9e3322df53990eea4a6fe4aa2455d63e72bbabe07cc3189ca69fe0147522103d2586577e2f4460b1c299f4f74b719982c1982f31e48ca1ef406347472fa611221038097033cb34a88b8bfc052adbbfefa8e92c33b7635e30c7a79d90ff4917c6c0b52aeffffffff034a140000000000001976a914a2bebc3bbc138a248296ad96e6aaf71d83f69c3688ac0000000000000000166a146f6d6e69000000000000008900000000058b114022020000000000001976a914c18bb19ca8f23be298fd305f06f4e039cb10dca088ac00000000"
	privkey := "cRvLERMVjEND2XGi1YEgPjQT6KkshQadJjtmBkbUgcQvJ5ZXNY6P"
	transaction := DecodeRawTransaction(redeemhex, tool.GetCoreNet())
	log.Println("redeemhex 部分签", transaction)

	redeemhexA := "0200000002ac5a7c14bf3a63944b50333e12263213fdf59eacfc1a60af27298324db736aea00000000da00473044022052864a0e9a3ba7175506b6aaf21229fed03ca05c42846de050e53603c55ff37302204b3f484f91f3e4ac799a98afca351309921c857c38298e351ca5a0328b37291101483045022100c4acce9704328d6ffd85e1c4cd358b999c9590adff9db3ddf24a388f282cc8cf022024660b7321d7502e4ca71b989cda642d676714c400a6bbc4b5f2dd205f512d0c0147522103d2586577e2f4460b1c299f4f74b719982c1982f31e48ca1ef406347472fa611221038097033cb34a88b8bfc052adbbfefa8e92c33b7635e30c7a79d90ff4917c6c0b52aeffffffffac5a7c14bf3a63944b50333e12263213fdf59eacfc1a60af27298324db736aea02000000d9004730440220410c2c64d9cf4a5b9ffd390a3f2b9cb493c34990da20d00b3dbbbec0195a4df202205dea90808f168b2f5dc61f4bb3e5293688cb5cdfa8b851e4500e615e749486a0014730440220701fde05a432c78f541aaa4e10a714e79a60daa3d764d0830ca2195d7d85f325022053cf5f59dad9e3322df53990eea4a6fe4aa2455d63e72bbabe07cc3189ca69fe0147522103d2586577e2f4460b1c299f4f74b719982c1982f31e48ca1ef406347472fa611221038097033cb34a88b8bfc052adbbfefa8e92c33b7635e30c7a79d90ff4917c6c0b52aeffffffff034a140000000000001976a914a2bebc3bbc138a248296ad96e6aaf71d83f69c3688ac0000000000000000166a146f6d6e69000000000000008900000000058b114022020000000000001976a914c18bb19ca8f23be298fd305f06f4e039cb10dca088ac00000000"
	transaction = DecodeRawTransaction(redeemhexA, tool.GetCoreNet())
	log.Println("redeemhexA", transaction)

	inputs := []bean.RawTxInputItem{}
	item := bean.RawTxInputItem{}
	item.ScriptPubKey = "a914a1617398d1a34529bbd35eeb5c30a4ce20a73d2b87"
	redeemScript := "522103d2586577e2f4460b1c299f4f74b719982c1982f31e48ca1ef406347472fa611221038097033cb34a88b8bfc052adbbfefa8e92c33b7635e30c7a79d90ff4917c6c0b52ae"
	item.RedeemScript = redeemScript
	inputs = append(inputs, item)
	item = bean.RawTxInputItem{}
	item.ScriptPubKey = "a914a1617398d1a34529bbd35eeb5c30a4ce20a73d2b87"
	redeemScript = "522103d2586577e2f4460b1c299f4f74b719982c1982f31e48ca1ef406347472fa611221038097033cb34a88b8bfc052adbbfefa8e92c33b7635e30c7a79d90ff4917c6c0b52ae"
	item.RedeemScript = redeemScript
	inputs = append(inputs, item)

	sign, err := SignRawHex(inputs, redeemhex, privkey)
	log.Println(err)
	log.Println(sign)
	transaction = DecodeRawTransaction(sign, tool.GetCoreNet())
	log.Println("sign 完成", transaction)
}

func TestSign4(t *testing.T) {
	sourcehex := "02000000021733ee13399d1e21c7e7fdb54ca1592bacd93ad691174486af55e8e335f9edbc00000000d900473044022066c6c4061564a00b5bb4823ca47d6ccee35be9e18e6aff0cafe776b1835bf6b202201d92f1c318bf33c662b0f9fc51b438976f48d3ee0cbab9352aab153047329f810147304402203f9c4cbf91cb0686ef89a5d1a2ef0d6fac057156a9b0f2ef0eb0feb15737d55b0220303f8be69010ee078ec2680e0952b5eb0edb39600de2a2c11f339a0097c0808a0147522102ab22188dd37966ab8f56fc36559c02fe0c498e492c90fc20f740ec8f45aff30021038097033cb34a88b8bfc052adbbfefa8e92c33b7635e30c7a79d90ff4917c6c0b52aeffffffff1733ee13399d1e21c7e7fdb54ca1592bacd93ad691174486af55e8e335f9edbc02000000da00473044022057b988b41809b514a15e1785ffddbd15a949f2a7c99695c65402cf01e2a8a92a022045a4e6226bc7ce7b58982179f380e57762373d5f4b9b9bda96c93b7b7e1b082801483045022100f1847c4b201063cfa12b07cf2255df862e08026cd5e935ad210922941431551e02200708f4a2328ecf91cdcdd9b249a1adf33b9ea9cda21f6382bc58ca88d7c1a4790147522102ab22188dd37966ab8f56fc36559c02fe0c498e492c90fc20f740ec8f45aff30021038097033cb34a88b8bfc052adbbfefa8e92c33b7635e30c7a79d90ff4917c6c0b52aeffffffff034a140000000000001976a914a2bebc3bbc138a248296ad96e6aaf71d83f69c3688ac0000000000000000166a146f6d6e6900000000000000890000000005a995c022020000000000001976a914c18bb19ca8f23be298fd305f06f4e039cb10dca088ac00000000"
	err := DecodeRawTransaction(sourcehex, tool.GetCoreNet())
	log.Println(err)
}
