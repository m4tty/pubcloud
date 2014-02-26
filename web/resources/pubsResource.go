package resources

type Pubs struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	LicName     string `json:"licName"`
	XCoord      int    `json:"xCoord"`
	YCoord      int    `json:"yCoord"`
}
