package netbox

type VLANGroup struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Site *Site  `json:"site"`
}
