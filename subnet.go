package gosolar

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// Subnet this is the json representation of subnet
type Subnet struct {
	SubnetID       int    `json:"SubnetId"`
	Address        string `json:"Address"`
	CIDR           string `json:"CIDR"`
	FriendlyName   string `json:"FriendlyName"`
	DisplayName    string `json:"DisplayName"`
	AvailableCount int    `json:"AvailableCount"`
	ReservedCount  int    `json:"ReservedCount"`
	UsedCount      int    `json:"UsedCount"`
	TotalCount     int    `json:"totalCount"`
	Comments       string `json:"Comments"`
	VLAN           int    `json:"VLAN"`
	AddressMask    string `json:"AddressMask"`
}

// GetSubnet Gets a subnet by display name.
func (c *Client) GetSubnet(subnetName string) Subnet {
	query := `SELECT	Address, 
						CIDR, 
						AddressMask, 
						DisplayName, 
						FriendlyName, 
						Reserved, 
						TotalCount, 
						UsedCount, 
						AvailableCount, 
						ReservedCount, 
						TransientCount, 
							StatusName 
					FROM IPAM.Subnet
					WHERE DisplayName = @name`

	parameters := map[string]interface{}{
		"name": subnetName,
	}

	res, err := c.QueryRow(query, parameters)

	if err != nil {
		log.Fatal(err)
	}

	var subnet Subnet
	bodyString := string(res)
	log.Debugf("ResponseString %s", bodyString)

	if err := json.Unmarshal(res, &subnet); err != nil {
		log.Fatal(err)
	}
	return subnet
}

// ListSubnets Lists subnets.
func (c *Client) ListSubnets() []Subnet {
	query := `SELECT	Address, 
						CIDR, 
						AddressMask, 
						DisplayName, 
						FriendlyName, 
						TotalCount, 
						UsedCount, 
						AvailableCount, 
						ReservedCount, 
						TransientCount, 
							StatusName 
					FROM IPAM.Subnet`

	res, err := c.Query(query, nil)

	if err != nil {
		log.Fatal(err)
	}

	var subnets []Subnet
	bodyString := string(res)

	if err := json.Unmarshal(res, &subnets); err != nil {
		log.Info("Couldnt unmarshal responseString %s", bodyString)
		log.Fatal(err)
	}
	return subnets
}
