package main

import (
	"fmt"
	storageHandler "interface_changer/database"
	"log"
	"os/exec"
	"strings"

	"github.com/rivo/tview"
)

func setDNS(adapterName, primaryDns string, secondaryDns string) error {
	cmd := exec.Command("netsh", "interface", "ip", "set", "dns", fmt.Sprintf("name=%s", adapterName), "static", primaryDns)

	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("netsh", "interface", "ip", "add", "dns", fmt.Sprintf("name=%s", adapterName), secondaryDns, "index=2")
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DNS set successfully:")
	return nil
}

var adapter storageHandler.InterfaceConfig
var dnss []storageHandler.DnsConfig
var dnsNames []string

func main() {
	var err error
	adapter, err = storageHandler.LoadInterfaceConfigs()
	if err != nil {
		log.Fatal(err)
	}
	dnss, err = storageHandler.LoadMultipleDNSConfigs()
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range dnss {
		dnsNames = append(dnsNames, s.Name)
	}

	app := tview.NewApplication()
	list := createMainMenu(app)

	if err := app.SetRoot(list, true).Run(); err != nil {
		panic(err)
	}

}

func createMainMenu(app *tview.Application) *tview.List {
	return tview.NewList().
		AddItem("Change DNS", "Select dns", '1', func() {
			switchDropdown(app)
		}).
		AddItem("Add Dns", "add dns to database", '2', func() {
			switchAddDns(app)
		}).
		AddItem("Change Interface", "default is wifi", '3', func() {
			switchChangeInterface(app)
		}).
		AddItem("Quit", "Exit the program", 'q', func() {
			app.Stop()
		})
}
func switchToMainMenu(app *tview.Application, mainMenu *tview.List) {
	app.SetRoot(mainMenu, true)
}

func switchDropdown(app *tview.Application) {

	dropdown := tview.NewDropDown().
		SetLabel("Select DNS option (hit Enter): ").
		SetOptions(dnsNames, func(text string, index int) {

			selectedDNS := dnss[index]
			fmt.Printf("Selected DNS: %s\n", selectedDNS.Name)
			setDNS(adapter.Name, selectedDNS.PrimaryDNS, selectedDNS.SecondaryDNS)

			app.Stop()
		})

	form := tview.NewForm().
		AddFormItem(dropdown).
		AddButton("Back", func() {

			switchToMainMenu(app, createMainMenu(app))
		})

	app.SetRoot(form, true)
}

func getInterfaces() []string {
	cmd := exec.Command("netsh", "interface", "show", "interface")
	output, error := cmd.CombinedOutput()
	if error != nil {
		log.Fatal(error)
	}
	var details []string

	netshInterfaces := strings.Split(string(output), "\n")
	interfaces := netshInterfaces[3:]
	for _, s := range interfaces {
		if len(s) < 5 {
			continue
		}
		interfacesDetail := strings.Split(s, "        ")
		details = append(details, interfacesDetail[2])
	}
	return details
}

func switchAddDns(app *tview.Application) {
	inputName := tview.NewInputField().SetLabel("Name: ")
	inputPrimaryDns := tview.NewInputField().SetLabel("Primary Dns: ")
	inputSecondaryDns := tview.NewInputField().SetLabel("Secondary Dns: ")

	form := tview.NewForm().
		AddFormItem(inputName).
		AddFormItem(inputPrimaryDns).
		AddFormItem(inputSecondaryDns).
		AddButton("Save", func() {
			primaryDns := inputPrimaryDns.GetText()
			name := inputName.GetText()
			secondaryDns := inputSecondaryDns.GetText()
			storageHandler.SaveDNSConfigs(storageHandler.DnsConfig{
				PrimaryDNS:   primaryDns,
				SecondaryDNS: secondaryDns,
				Name:         name,
			})
			var err error
			dnss, err = storageHandler.LoadMultipleDNSConfigs()
			if err != nil {
				log.Fatal(err)
				app.Stop()
			}
			dnsNames = nil
			for _, s := range dnss {
				dnsNames = append(dnsNames, s.Name)
			}
			switchToMainMenu(app, createMainMenu(app))
		}).
		AddButton("Back", func() {

			switchToMainMenu(app, createMainMenu(app))
		})

	app.SetRoot(form, true)
}

func switchChangeInterface(app *tview.Application) {
	interfaces := getInterfaces()
	dropdown := tview.NewDropDown().
		SetLabel("Select Interface option (hit Enter): ").
		SetOptions(interfaces, func(text string, index int) {

			storageHandler.SaveInterfaceConfigs(storageHandler.InterfaceConfig{
				Name: strings.TrimSpace(text),
			})

			switchToMainMenu(app, createMainMenu(app))
		})

	form := tview.NewForm().
		AddFormItem(dropdown).
		AddButton("Back", func() {

			switchToMainMenu(app, createMainMenu(app))
		})

	app.SetRoot(form, true)
}
