package themo

type Me struct {
	DevicesCount int    `json:"Devices_Count"`
	Id           int    `json:"ID"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	EmailAddress string `json:"EmailAddress"`
	Phone        string `json:"Phone"`
}

func (t *Themo) Me() (*Me, error) {
	var me *Me
	if _, err := t.client.R().
		SetResult(&me).
		Get("/api/clients/me"); err != nil {
		return nil, err
	} else {
		return me, nil
	}
}
