package gosolar

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)

// IPAddress is a struct for basic unmarshalling of an IPAddress object
type IPAddress struct {
	Address string `json:"DisplayName"`
	Status  string `json:"Status"`
	//TransientCount string `json"Transient"`
}

func (c *Client) GetIP(subnetAddress string, subnetCIDR string) IPAddress {

	body := fmt.Sprintf("{\"%s\", \"%s\"}", subnetAddress, subnetCIDR)
	res, err := c.Invoke("IPAM.SubnetManagement", "GetFirstAvailableIp", body)

	// run the query without with the parameters map above
	bodyString := string(res)

	if err != nil {
		log.Infof("ResponseString %s", bodyString)
		log.Fatal(err)
	}

	var ip IPAddress

	// This should catch an empty ip.
	if err := json.Unmarshal(res, &ip); err != nil {
		log.Infof("ResponseString %s", bodyString)
		log.Fatal(err)
	}

	return ip
}

// Set-SwisObject $swis -Uri 'swis://localhost/Orion/IPAM.IPNode/IPAddress=1.1.1.1' -Properties @{ Alias = 'test1' }
//Invoke-SwisVerb $swis IPAM.SubnetManagement GetFirstAvailableIp @("199.10.1.0", "24")
//Invoke-SwisVerb $swis IPAM.SubnetManagement ChangeIPStatus  @("199.10.1.1", "Used")
//Update: Set-SwisObject $swis -Uri 'swis://localhost/Orion/IPAM.IPNode/IpNodeId=2' -Properties @{ Status = 'Used', Comment = "Reserved by terraform." }

// ReserveIP will set the IP Status to "Used"
func (c *Client) ReserveIP(ipAddress string) string {
	body := fmt.Sprintf("{\"%s\", \"Used\"}", ipAddress)
	result, err := c.Invoke("IPAM.SubnetManagement", "ChangeIPStatus", body)

	resultString := string(result)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(resultString)
	return resultString
}

// ReleaseIP will set the IP Status to "Unused"
func (c *Client) ReleaseIP(ipAddress string) string {
	body := fmt.Sprintf("{\"%s\", \"Unused\"}", ipAddress)
	result, err := c.Invoke("IPAM.SubnetManagement", "ChangeIPStatus", body)
	resultString := string(result)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(resultString)
	return resultString
}
