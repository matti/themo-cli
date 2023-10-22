package themo

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type themoDeviceData struct {
	Id int `json:"ID"`
}
type themoDevice struct {
	Devices []*themoDeviceData `json:"Devices"`
}

type Device struct {
	client *resty.Client
	Id     int
}

func (t *Themo) Devices(clientId int) ([]*Device, error) {
	var themoDevice *themoDevice

	req := t.client.R().
		SetQueryParam("clientID", fmt.Sprint(clientId)).
		SetQueryParam("state", "false").
		SetQueryParam("page", "0").
		SetQueryParam("pageSize", "-1").
		SetResult(&themoDevice)

	if _, err := req.Get("/api/devices"); err != nil {
		return nil, err
	} else {
		var devices []*Device
		for _, td := range themoDevice.Devices {
			devices = append(devices, &Device{
				client: t.client,
				Id:     td.Id,
			})
		}
		return devices, nil
	}
}

func (d *Device) SetColor(low *Color, high *Color) error {
	req := d.client.R().
		SetBody(fmt.Sprintf(`
{
	"Colors": [
		{
			"RGB": "rgb(%d, %d, %d)"
		},
		{
			"RGB": "rgb(%d, %d, %d)"
		}
	],
}
`, low.R, low.G, low.B, high.R, high.G, high.B))

	req.Header.Set("Content-Type", "application/json")

	if _, err := req.Put(fmt.Sprintf("/api/devices/color/%d", d.Id)); err != nil {
		return err
	} else {
		return nil
	}
}

// {
//   "Colors": [
//     {
//       "ID": 18337,
//       "Name": "Lowest price",
//       "RGB": "rgb(255, 0, 0)"
//     },
//     {
//       "ID": 18338,
//       "Name": "Highest price",
//       "RGB": "rgb(0, 255, 0)"
//     }
//   ],
//   "ID": 9170
// }

type themoColor struct {
	RGB string
}

type themoColors struct {
	Colors []*themoColor
}

func (d *Device) Colors() ([]*Color, error) {
	var themoColors *themoColors

	if _, err := d.client.R().SetResult(&themoColors).Get(fmt.Sprintf("/api/devices/%d/color", d.Id)); err != nil {
		return nil, err
	} else {
		var colors []*Color
		for _, rs := range themoColors.Colors {
			if c, err := ParseColor(rs.RGB); err != nil {
				log.Fatalln(err)
			} else {
				colors = append(colors, c)
			}
		}
		return colors, nil
	}
}
