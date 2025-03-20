# BlueLocker 🔵🔒

**BlueLocker** is a simple and efficient tool that automatically locks your laptop when your **Galaxy Watch** (or any Bluetooth device) disconnects. It enhances security by ensuring your device is locked when you step away.

## 🚀 Features
- 🔍 **Scans for Bluetooth devices** every few seconds.
- 🛡️ **Locks the laptop** when the target device is not found.
- ⚡ Lightweight and fast, built with **Golang**.
- 🏆 Works on **Linux** (requires `gnome-screensaver` or a custom lock command).

## 🛠️ Installation
### 1. Clone the repository:
```sh
git clone https://github.com/yourusername/BlueLocker.git
cd BlueLocker
```

### 2. Install dependencies:
Make sure **Go** is installed on your system. If not, install it first: [Go Installation Guide](https://go.dev/doc/install)

```sh
go mod tidy
```

### 3. Run the program:
```sh
sudo go run monitor.go
```
**(Requires sudo to access Bluetooth adapter)**

## 🔧 Configuration
### Set Your Device's MAC Address
Find your Bluetooth device’s MAC address using:
```sh
hcitool scan
```
Edit `monitor.go` and set `targetMAC` to your device’s MAC address:
```go
targetMAC = "DC:C4:9C:B3:D4:17" // Example MAC
```

## 📌 Usage
- Keep your Bluetooth device near your laptop.
- If the device disconnects, **BlueLocker** locks your screen.
- Reconnect your device to unlock manually.

## 🏗️ Future Improvements
- ✅ Support for multiple devices.
- ✅ Custom lock commands for different desktop environments.
- ✅ Windows & macOS support.

## 💡 Contributing
Pull requests are welcome! If you have suggestions, create an issue or contribute directly. 🙌

## 📜 License
This project is licensed under the **MIT License**.

---
Made with ❤️ by [Your Name]
