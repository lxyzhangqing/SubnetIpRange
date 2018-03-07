package SubnetIpRange

import "testing"

func TestRangeIp(t *testing.T) {
	beg_ip, end_ip, err := SubnetRange("172.16.1.19/20")
	if err != nil {
		t.Errorf("Error = %v\n", err)
	}

	if beg_ip != "172.16.1.19" || end_ip != "172.16.15.255" {
		t.Errorf("Ip range ivalid")
	}
}
