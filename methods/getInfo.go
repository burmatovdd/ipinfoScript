package methods

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (service *Service) GetInfo(ip string) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://ipinfo.io/widget/demo/"+ip, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("referer", "https://ipinfo.io/")
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	body := Response{}

	err = json.Unmarshal(bodyText, &body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("IP: ", body.Data.RespIp)
	fmt.Println("Subnet: ", body.Data.Asn.Route)
	fmt.Println("Company: ", body.Data.Company.Name)
	fmt.Println("Country: ", body.Data.Country)
	fmt.Println("Type: ", body.Data.Asn.Type)
	fmt.Println("ASN: ", body.Data.Asn.Asn)
	fmt.Println("--------------------------------------------------------------------------------------------")
}
