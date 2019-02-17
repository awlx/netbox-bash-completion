package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	netbox         = flag.String("netbox", "https://netbox.local", "Netbox BaseURL")
	netboxAPIToken = flag.String("netbox-api-token", "", "Mandatory: Netbox API Token")
	netboxDevice   = flag.String("netbox-device", "", "Device String to search for")
	tld            = flag.String("tld", ".local", "Default TLD for devices")
)

// NetboxResult the whole Json Reply
type NetboxResult struct {
	Count    int                `json:"count"`
	Next     interface{}        `json:"next"`
	Previous interface{}        `json:"previous"`
	Results  []NetboxResultList `json:"results"`
}

// NetboxResultList all Results of the json call
type NetboxResultList struct {
	ID      int    `json:"id"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

func main() {
	flag.Parse()
	if *netboxAPIToken == "" {
		panic("Please provide a netbox-api-token")
	}

	netboxClient := http.Client{
		Timeout: time.Second * 2,
	}

	devices := make(map[string]bool)

	devices = getAllDevices(netboxClient, *netboxDevice)
	for device := range devices {
		fmt.Println(fmt.Sprintf("%s%s", device, *tld))
	}
}

func getAllDevices(netboxClient http.Client, searchString string) (netboxDevices map[string]bool) {
	deviceMap := make(map[string]bool)

	deviceURL := fmt.Sprintf("%s/api/dcim/devices/?q=%s", *netbox, searchString)
	vmURL := fmt.Sprintf("%s/api/virtualization/virtual-machines/?q=%s", *netbox, searchString)

	// Query devices
	req, err := http.NewRequest(http.MethodGet, deviceURL, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Token %s", *netboxAPIToken))

	res, err := netboxClient.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}

	var netboxDeviceList NetboxResult

	bodyErr := json.Unmarshal(body, &netboxDeviceList)

	if bodyErr != nil {
		fmt.Println("Json Unmarshal failed with: ", err)
	}

	for i := 0; i < len(netboxDeviceList.Results); i++ {
		if _, device := deviceMap[netboxDeviceList.Results[i].Name]; !device {
			deviceMap[netboxDeviceList.Results[i].Name] = true
		}
	}

	// Query VMs
	req, err = http.NewRequest(http.MethodGet, vmURL, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Token %s", *netboxAPIToken))

	res, err = netboxClient.Do(req)
	if err != nil {
		panic(err)
	}

	body, err = ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}

	var netboxVMList NetboxResult

	bodyErr = json.Unmarshal(body, &netboxVMList)

	if bodyErr != nil {
		fmt.Println("Json Unmarshal failed with: ", err)
	}

	for i := 0; i < len(netboxVMList.Results); i++ {
		if _, device := deviceMap[netboxVMList.Results[i].Name]; !device {
			deviceMap[netboxVMList.Results[i].Name] = true
		}
	}

	return (deviceMap)
}
