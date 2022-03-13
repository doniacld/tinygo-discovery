# TinyGo Discovery

A repository to explore [TinyGo](https://tinygo.org/) features and have fun.  
All the examples have been tested on an Arduino Nano IoT 33.

## Prerequisites

Install [TinyGo](https://tinygo.org/getting-started/install/)

If you do not have an arduino or a compatible microcontroller; you can use to begin the [TinyGo playground](https://play.tinygo.org/) to begin.

## Examples

| Example     | Description                                                                                      |
|-------------|--------------------------------------------------------------------------------------------------|
| Blink       | The HelloWorld of the microcontroller. Blink a led.                                              |
| Serial      | Write on a serial port and use serial reader in utils to read values.                            |
| Thermometer | Read temperature and humidity from a DHT22 sensor and use serial reader in utils to read values. |

### Quickstart

Flash your arduino-nano33 with the blink example:
```bash
$ make led
```
You should observe a LED blinking on the board.
