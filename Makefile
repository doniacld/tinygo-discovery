# Go parameters
GOCMD=go
GORUN=$(GOCMD) run

# TinyGo parameters
TINYGOCMD=tinygo
TINYGOFLASH=$(TINYGOCMD) flash

TARGET=arduino-nano33

led:
	$(TINYGOFLASH) -target=$(TARGET) blink/blink.go

hello:
	$(TINYGOFLASH) -target=$(TARGET) serial/serial.go

therm:
	$(TINYGOFLASH) -target=$(TARGET) thermometer/thermometer.go

readserial:
	$(GORUN) utils/read_serial.go
