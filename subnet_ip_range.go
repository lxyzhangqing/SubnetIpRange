package SubnetIpRange

import (
	"fmt"
	"regexp"
	"strings"
)

func is_subnet(subnet string) bool {
	pattern := `^(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[1-9])\.`
	pattern += `((1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)\.){2}`
	pattern += `(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)`
	pattern += `(/(\d|[1-2]\d|3[0-2]))$`
	match, err := regexp.MatchString(pattern, subnet)
	return err == nil && match
}

func SubnetRange(subnet string) (string, string, error) {
	if is_subnet(subnet) == false {
		return "","", fmt.Errorf("subnet format invalid")
	}

	subnet_info := [5]uint32{}
	_, err := fmt.Sscanf(subnet,"%d.%d.%d.%d/%d",&subnet_info[0],
		&subnet_info[1], &subnet_info[2], &subnet_info[3], &subnet_info[4])

	if err != nil {
		return "","", err
	}

	var beg_ip_n uint32 = subnet_info[0] << 24 + subnet_info[1] << 16 + subnet_info[2] << 8 + subnet_info[3]
	var end_ip_n uint32 = (0xFFFFFFFF >> subnet_info[4]) | beg_ip_n
	return net_ntoa(beg_ip_n),net_ntoa(end_ip_n),nil
}

func net_ntoa(net uint32) string {
	ip := ""
	for i := 0; i < 4; i++ {
		ip = fmt.Sprintf(".%v", net % 256) + ip
		net = net >> 8
	}
	return strings.TrimLeft(ip,".")
}
