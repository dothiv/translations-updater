package model

type Sys struct {
	Type      string `json:"type"`
	Id        string `json:"id"`
	Revision  int    `json:"revision"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Locale    string `json:"locale"`
}

type Error struct {
	Name string   `json:"name"`
	Path []string `json:"path"`
}

type Details struct {
	Errors []Error `json:"errors"`
}

type Code struct {
	En string `json:"en"`
}

type Value struct {
	En string `json:"en"`
	De string `json:"de"`
}

type Fields struct {
	Code  Code  `json:"code"`
	Value Value `json:"value"`
}

type LocaleItem struct {
	Sys    Sys               `json:"sys"`
	Fields map[string]string `json:"fields"`
}

type Item struct {
	Sys    Sys    `json:"sys"`
	Fields Fields `json:"fields"`
}

type SearchResponse struct {
	Sys       Sys          `json:"sys"`
	Total     int          `json:"total"`
	Skip      int          `json:"skip"`
	Limit     int          `json:"limit"`
	Items     []LocaleItem `json:"items"`
	Message   string       `json:"message"`
	Details   Details      `json:"details"`
	RequestId string       `json:"requestId"`
}
