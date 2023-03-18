package main

import (
	"testing"
)

func Test_generateSasToken(t *testing.T) {
	hostname := "sbox-iot.azure-devices.net"
	deviceId := "wsl2"
	key := "YIklI05JZpMpoZ6hN9OPCjFOZmlxrUgaqHWlMtNpgvI=" // 再生成して無効化したもの
	var expires int64 = 1679107024

	expect := "SharedAccessSignature sr=sbox-iot.azure-devices.net%2Fdevices%2Fwsl2&sig=%2FLEYFLyV5VU7ad404%2BQXU4wtA25KYtshdndkeD30UDU%3D&se=1679107024"
	actual := generateSasToken(hostname, deviceId, key, expires)
	if actual != expect {
		t.Errorf("Expect \n%s, but got \n%s", expect, actual)
	}
}
