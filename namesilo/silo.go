package namesilo

import (
	"fmt"
	"strings"
)

type Request struct {
	Operation string `xml:"request>operation"`
	IP        string `xml:"request>ip"`
}

type Reply struct {
	Code   uint   `xml:"reply>code"`
	Detail string `xml:"reply>detail"`
}

type SiloClient struct {
	Version uint
	Type    string
	Key     string
}

func (s *SiloClient) Url(operation string, params map[string]string) string {
	params_arr := make([]string, len(params)+3)

	idx := 0

	params_arr[idx] = fmt.Sprintf("version=%d", s.Version)
	idx++

	params_arr[idx] = "type=" + s.Type
	idx++

	params_arr[idx] = "key=" + s.Key
	idx++

	for k, v := range params {
		params_arr[idx] = k + "=" + v
		idx++
	}

	return fmt.Sprintf("https://www.namesilo.com/api/%s?%s", operation, strings.Join(params_arr, "&"))
}
