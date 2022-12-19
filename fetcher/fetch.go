package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// 原生的方式
func Fetch(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	req.Header.Set("Cookie", "bid=AmBItSsYb_M; gr_user_id=3c04c011-bad5-41f5-a49a-7424085f9aea; douban-fav-remind=1; _vwo_uuid_v2=D5CBE7681A52FE69D36480CEBCC9DDED4|8f178091c26bbe5559804e930aadf43d; __utmz=30149280.1671352287.6.5.utmcsr=cn.bing.com|utmccn=(referral)|utmcmd=referral|utmcct=/; __utmz=81379588.1671352287.3.3.utmcsr=cn.bing.com|utmccn=(referral)|utmcmd=referral|utmcct=/; __gads=ID=ce67fea25bc84ba9-224a05f8ecd80002:T=1671352288:RT=1671352288:S=ALNI_Mbe_hAEyEu_XVOGY-PECIKt5vfGkw; __utmc=30149280; __utmc=81379588; viewed=\"36104107_30329536_1842426\"; __yadk_uid=fQuFUHKdW77GudKbdU1ODFwQhqHmR5FE; _pk_ref.100001.3ac3=%5B%22%22%2C%22%22%2C1671427861%2C%22https%3A%2F%2Fcn.bing.com%2F%22%5D; _pk_ses.100001.3ac3=*; __utma=30149280.38393203.1660027411.1671378815.1671427861.11; __utma=81379588.41032280.1660280940.1671378815.1671427861.8; ap_v=0,6.0; __gpi=UID=00000b92fe419763:T=1671352288:RT=1671427862:S=ALNI_Mb0qPqjl8y_-37gt7a01zzPQcgngw; gr_cs1_c91d473b-1f28-4eb8-84d6-153c961e1bbc=user_id%3A0; ll=\"118208\"; __utmt=1; dbcl2=\"175522715:OudDw1yhUf4\"; ck=Aj8P; _ga=GA1.1.38393203.1660027411; _ga_RXNMP372GL=GS1.1.1671428494.1.0.1671428505.49.0.0; __utmt_douban=1; push_noty_num=0; push_doumail_num=0; gr_session_id_22c937bbd8ebd703f2d8e9445f7dfd03=ce452105-d118-40ce-9bde-10e8269ec38d; gr_cs1_ce452105-d118-40ce-9bde-10e8269ec38d=user_id%3A1; gr_session_id_22c937bbd8ebd703f2d8e9445f7dfd03_ce452105-d118-40ce-9bde-10e8269ec38d=true; __utmb=30149280.8.9.1671428422903; __utmb=81379588.6.10.1671427861; _pk_id.100001.3ac3=7fd7cf10ed95de44.1657903034.9.1671428530.1671378815.; frodotk_db=\"2bd64a54c977b7be8382defe5c67273a\"")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bodyReader := bufio.NewReader(resp.Body)
	// 探测编码
	ec := DetectEncoding(bodyReader)
	// 生成对应编码的转换阅读器
	utf8Reader := transform.NewReader(bodyReader, ec.NewDecoder())
	// 阅读器提供内容
	return ioutil.ReadAll(utf8Reader)

}

// 编码探测函数，返回探测出的编码
func DetectEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		fmt.Println("error: ", err)
		return unicode.UTF8
	}
	// 这个包也要安装
	// go get golang.org/x/net/html/charset
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e

}
