package lxml

import (
	"testing"
)

func TestXml(t *testing.T) {
	ret, err := ParseXml([]byte(`
	<xml><appid><![CDATA[wx3fd74ebab76b843a]]></appid>
    <attach><![CDATA[attach]]></attach>
    <bank_type><![CDATA[CFT]]></bank_type>
    <cash_fee><![CDATA[28]]></cash_fee>
    <coupon_count><![CDATA[1]]></coupon_count>
    <coupon_fee>12</coupon_fee>
    <coupon_fee_0><![CDATA[12]]></coupon_fee_0>
    <coupon_id_0><![CDATA[2000000081712314105]]></coupon_id_0>
    <fee_type><![CDATA[CNY]]></fee_type>
    <is_subscribe><![CDATA[N]]></is_subscribe>
    <mch_id><![CDATA[1525958141]]></mch_id>
    <nonce_str><![CDATA[2718669443746]]></nonce_str>
    <openid><![CDATA[oheXT5OLQ56tquSaWrct8os4nQgw]]></openid>
    <out_trade_no><![CDATA[201908081152524c18bd007e2c]]></out_trade_no>
    <result_code><![CDATA[SUCCESS]]></result_code>
    <return_code><![CDATA[SUCCESS]]></return_code>
    <sign><![CDATA[19F14BEC21FA95832FA4CC375E63FEA4]]></sign>
    <time_end><![CDATA[20190808115319]]></time_end>
    <total_fee>40</total_fee>
    <trade_type><![CDATA[JSAPI]]></trade_type>
    <transaction_id><![CDATA[4200000339201908083414090578]]></transaction_id>
</xml>`))
	t.Logf("result %v, err: %v", ret, err)
}
