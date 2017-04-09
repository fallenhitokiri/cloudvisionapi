package cloudvisionapi

type APIResponse struct {
	Name   string
	Status int

	Labels []*Label
	Colors []*Color
	Logos  []*Logo
	Texts  []*Text
}

type Label struct {
	Description string
	Score       float64
	Confidence  float64
}

type Color struct {
	Alpha float64
	Blue  float64
	Green float64
	Red   float64
	Score float64
}

type Logo struct {
	Description string
	Locale      string
	Score       float64
	Confidence  float64
}

type Text struct {
	Description string
	Locale      string
	Score       float64
	Confidence  float64
}
