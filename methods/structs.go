package methods

type Data struct {
	Data []Ip `json:"data"`
}

type Ip struct {
	Ip string `json:"ip"`
}

type Response struct {
	Input string   `json:"input"`
	Data  RespData `json:"data"`
}

type RespData struct {
	RespIp   string  `json:"ip"`
	City     string  `json:"city"`
	Region   string  `json:"region"`
	Country  string  `json:"country"`
	Loc      string  `json:"loc"`
	Org      string  `json:"org"`
	Postal   string  `json:"postal"`
	Timezone string  `json:"timezone"`
	Asn      Asn     `json:"asn"`
	Company  Company `json:"company"`
	Privacy  Privacy `json:"privacy"`
	Abuse    Abuse   `json:"abuse"`
}

type Asn struct {
	Asn    string `json:"asn"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Route  string `json:"route"`
	Type   string `json:"type"`
}

type Company struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Type   string `json:"type"`
}

type Privacy struct {
	Vpn     bool   `json:"vpn"`
	Proxy   bool   `json:"proxy"`
	Tor     bool   `json:"tor"`
	Relay   bool   `json:"relay"`
	Hosting bool   `json:"hosting"`
	Service string `json:"service"`
}

type Abuse struct {
	Address string `json:"address"`
	Country string `json:"country"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Network string `json:"network"`
	Phone   string `json:"phone"`
}

type Service struct {
	method *Methods
}

type Methods interface {
	GetInfo()
}
