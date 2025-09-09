# raspberrykeyboard
A keyboard intercept for Raspberry PI, so I can create macros and other cool recording systems.

# Raspberry Pi Keyboard/Mouse Intercept Device

This Go application is designed to run on a Raspberry Pi and turn it into a USB HID device (keyboard/mouse) for a host computer. It will:

- Intercept keyboard and mouse events on the Pi
- Emulate a USB HID device using Linux gadget drivers
- Record and replay input events based on key combinations

## Getting Started

1. **Run as root:** The application must be run as root to access input devices and configure USB gadget mode.
2. **Dependencies:**
	- Go 1.21+
	- Linux kernel with configfs and USB gadget support
	- `golang.org/x/sys` for input event handling

## Usage

```bash
sudo go run ./cmd/main.go
```

## Next Steps
- Implement event capture from `/dev/input/event*`
- Add USB HID emulation logic
- Add recording and replay features
