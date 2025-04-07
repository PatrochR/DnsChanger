# DNSChangerCLI

A simple and interactive CLI tool built with Go to change your Windows DNS settings using a TUI (Text User Interface). It uses the [`tview`](https://github.com/rivo/tview) library to create a smooth text-based GUI inside the terminal.

> ⚠️ Requires **administrator privileges** to modify network settings.

---

## ✨ Features

- Browse and switch DNS configurations from a list
- Save and load multiple DNS profiles using local JSON storage
- Automatically applies the selected DNS to your active network adapter
- Simple keyboard navigation via terminal interface
- Optional batch script or scheduled task support for admin access

---

## 📦 Installation

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/DNSChangerCLI.git
cd DNSChangerCLI
```
### 2. Build the App
Make sure you have [Go installed](https://golang.org/dl/)

```bash
go build -o dnschanger.exe
```
---
## ⚙️ Usage
Run the application (must be as Administrator):

```bash
dnschanger.exe
```
Or via batch script:
```bash
run_with_admin.bat
```
---
## 🧠 Local DNS Storage
DNS configurations are saved in a JSON file:
```json
[
  {
    "Name": "Google DNS",
    "PrimaryDNS": "8.8.8.8",
    "SecondaryDNS": "8.8.4.4"
  },
  {
    "Name": "Cloudflare",
    "PrimaryDNS": "1.1.1.1",
    "SecondaryDNS": "1.0.0.1"
  }
]
```
You can edit (`interface_configs.json`) manually or add new entries from the UI.
---
## 📋 Dependencies
- [rivo/tview](https://github.com/rivo/tview)
- Standard Go libraries
Install them with:
```bash
go get github.com/rivo/tview
```
[RezaCharsetad](https://github.com/PatrochR)
