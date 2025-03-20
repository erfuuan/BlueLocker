package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"tinygo.org/x/bluetooth"
)

const targetMACAddress = "DC:C4:9C:B3:D4:17"

// const targetMACAddress = "38:8A:06:67:65:A5"

func lockLaptop() {
	fmt.Println("Locking laptop...")

	// Try `loginctl lock-session` (Recommended)
	cmd := exec.Command("loginctl", "lock-session")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to lock using loginctl, trying gnome-screensaver-command...")

		// Fallback to `gnome-screensaver-command -l`
		cmd = exec.Command("gnome-screensaver-command", "-l")
		err = cmd.Run()
		if err != nil {
			log.Fatal("Failed to lock laptop:", err)
		}
	}
}

func main() {

	fmt.Println(" start scanning")

	adapter := bluetooth.DefaultAdapter
	err := adapter.Enable()
	if err != nil {
		log.Fatal("Failed to enable Bluetooth:", err)
	}

	var targetDevice bluetooth.ScanResult
	deviceFound := false

	// Start scanning
	err = adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		fmt.Println("Found device:", device.Address.String(), "- RSSI:", device.RSSI, "- LocalName:", device.LocalName())

		// time.Sleep(50 * time.Second)
		if device.Address.String() == targetMACAddress {
			fmt.Println("Target device found!")
			targetDevice = device
			deviceFound = true
			adapter.StopScan() // Stop scanning once the device is found
		} else {
			fmt.Println("device not found , now laptop is lock")
			// lockLaptop() // Lock laptop if connection fails

		}
	})
	if err != nil {
		log.Fatal("Error during scanning:", err)
	}

	time.Sleep(2 * time.Second)

	if deviceFound {
		fmt.Println(targetDevice)
		peripheral, err := adapter.Connect(targetDevice.Address, bluetooth.ConnectionParams{})

		if err != nil {
			fmt.Println("Failed to connect to device:", err)
			lockLaptop() // Lock laptop if connection fails
			return
		}
		defer peripheral.Disconnect()
		fmt.Println("Successfully connected to the device!")

		go func() {
			for {
				time.Sleep(1 * time.Second)

				// Check if the device is still connected using RSSI
				err := adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
					if device.Address.String() == targetMACAddress {
						fmt.Println("Device still connected, RSSI:", device.RSSI)
					}
				})
				if err != nil {
					lockLaptop()

					fmt.Println("Error scanning:", err)
				}

				// If scan does not find the device, assume disconnected
				if err != nil || targetDevice.Address.String() != targetMACAddress {
					fmt.Println("Device disconnected! Locking laptop...")
					lockLaptop()
					return
				}
			}
		}()

		// Keep the program running
		select {}

	} else {
		fmt.Println("Target device not found, locking laptop.")
		lockLaptop()
	}

}
