package request

type BookCondition struct {
	Num       string `json:"num"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	Isbn      string `json:"isbn"`
	Publisher string `json:"publisher"`
}
