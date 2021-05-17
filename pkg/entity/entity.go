package entity

type ApplicationsDoc []ApplicationDoc

type ApplicationDoc struct {
	Title     string `json:"title,omitempty"`
	Endpoints []struct {
		EndpointName   string   `json:"endpoint_name,omitempty"`
		Description    string   `json:"description,omitempty"`
		Method         string   `json:"method,omitempty"`
		URL            string   `json:"url,omitempty"`
		RequestBody    string   `json:"request_body,omitempty"`
		RequestFields  []Field  `json:"request_fields,omitempty"`
		Params         []string `json:"params,omitempty"`
		Response       string   `json:"response,omitempty"`
		ResponseFields []Field  `json:"response_fields,omitempty"`
	} `json:"endpoints,omitempty"`
}

type Field struct {
	Field       string `json:"field,omitempty"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
}
