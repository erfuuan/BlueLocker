package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// Function to execute commands and return the output
func runCommand(name string, args ...string) string {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing %s %v: %s\n", name, args, string(output))
	}
	return string(output)
}

// Function to show a desktop notification
func showAlert(title, message string) {
	cmd := exec.Command("notify-send", title, message, "-i", "error", "-t", "5000")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error showing alert: %v\n", err)
	}
}

// Function to scan for the Bluetooth device and lock the system if not found
func scanAndCheck() {
	for {
		fmt.Println("🔍 Scanning for Bluetooth devices...")
		runCommand("bluetoothctl", "scan", "on")
		time.Sleep(12 * time.Second)

		// Retrieve the list of available Bluetooth devices
		devices := runCommand("bluetoothctl", "devices")
		fmt.Println("📋 Device list:")
		fmt.Println(devices)

		// Check if the target device is in the list
		if strings.Contains(devices, deviceMAC) {
			fmt.Println("✅ Device Found:", deviceMAC)
		} else {
			fmt.Println("❌ Device Not Found! Locking system...")
			showAlert("Device Not Found", "Your device is not connected. Locking system...")

			time.Sleep(10 * time.Second)                  // Delay before locking the system
			runCommand("gnome-screensaver-command", "-l") // Lock the screen
		}

		time.Sleep(5 * time.Second) // Prevent excessive resource usage
	}
}

// The MAC address of the target Bluetooth device
var deviceMAC = "AA:A1:1A:A1:A1:11"

func main() {
	fmt.Println("🚀 Starting BlueLocker...")
	scanAndCheck()
}
