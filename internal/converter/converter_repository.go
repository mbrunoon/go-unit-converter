package converter

type ConverterRequest struct {
	Value float64
	From  string
	To    string
}

type ConverterResponse struct {
	Value float64 `json:"value"`
}
