package lxml

import (
	"testing"
)

func TestXml(t *testing.T) {
	ret, err := ParseXml([]byte(`
	<xml><attach><![CDATA[attach]]></attach>
    <bank_type><![CDATA[CFT]]></bank_type>
    <cash_fee><![CDATA[28]]></cash_fee>
    <coupon_count><![CDATA[1]]></coupon_count>
    <coupon_fee>12</coupon_fee>
    <coupon_fee_0><![CDATA[12]]></coupon_fee_0>
    <fee_type><![CDATA[CNY]]></fee_type>
    <is_subscribe><![CDATA[N]]></is_subscribe>
    <nonce_str><![CDATA[2718669443746]]></nonce_str>
    <result_code><![CDATA[SUCCESS]]></result_code>
    <return_code><![CDATA[SUCCESS]]></return_code>
    <time_end><![CDATA[20190808115319]]></time_end>
    <total_fee>40</total_fee>
    <trade_type><![CDATA[JSAPI]]></trade_type>
</xml>`))
	t.Logf("result %v, err: %v", ret, err)
}
