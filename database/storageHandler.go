package database

import (
	"encoding/json"
	"os"
)

type DnsConfig struct {
	Name         string `json:"name"`
	PrimaryDNS   string `json:"primary_dns"`
	SecondaryDNS string `json:"secondary_dns"`
}

type InterfaceConfig struct {
	Name string `json:"name"`
}

func SaveDNSConfigs(configs DnsConfig) error {

	ldata, err := LoadMultipleDNSConfigs()
	if err != nil {
		return err
	}
	ldata = append(ldata, configs)
	data, err := json.MarshalIndent(ldata, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("C:/Users/MOUOOD/Documents/Go Projects/interface_changer/dns_configs.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadMultipleDNSConfigs() ([]DnsConfig, error) {
	var configs []DnsConfig
	data, err := os.ReadFile("C:/Users/MOUOOD/Documents/Go Projects/interface_changer/dns_configs.json")
	if err != nil {
		return configs, err
	}

	err = json.Unmarshal(data, &configs)
	if err != nil {
		return configs, err
	}

	return configs, nil
}

func SaveInterfaceConfigs(configs InterfaceConfig) error {

	data, err := json.MarshalIndent(configs, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("C:/Users/MOUOOD/Documents/Go Projects/interface_changer/interface_configs.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadInterfaceConfigs() (InterfaceConfig, error) {
	var configs InterfaceConfig
	data, err := os.ReadFile("C:/Users/MOUOOD/Documents/Go Projects/interface_changer/interface_configs.json")
	if err != nil {
		return configs, err
	}

	err = json.Unmarshal(data, &configs)
	if err != nil {
		return configs, err
	}

	return configs, nil
}
