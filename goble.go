package goble

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/MarinX/serial"
)

// New func
func New(dev string, baud int) (*Ble, error) {

	c := &serial.Config{Name: dev, Baud: baud, ReadTimeout: 1 * time.Second}
	f, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}

	return &Ble{
		fd: f,
	}, nil
}

// Close func
func (t *Ble) Close() error {
	return t.fd.Close()
}

// TestConnection func
func (t *Ble) TestConnection() *BleResponse {
	return t.writeRead(test)
}

// GetDeviceAddress func
func (t *Ble) GetDeviceAddress() *BleResponse {
	return t.writeRead(getDeviceAddress)
}

// GetAdvertisingInterval func
func (t *Ble) GetAdvertisingInterval() *BleResponse {
	return t.writeRead(getAdvertisingInt)
}

// SetAdvertisingInterval func
func (t *Ble) SetAdvertisingInterval(interval BleAdvertisingInt) *BleResponse {
	return t.writeRead(fmt.Sprintf(setAdvertisingInt, interval))
}

// GetAdvertisingType func
func (t *Ble) GetAdvertisingType() *BleResponse {
	return t.writeRead(getAdvertisingType)
}

// SetAdvertisingType func
func (t *Ble) SetAdvertisingType(typ BleAdvertisingType) *BleResponse {
	return t.writeRead(fmt.Sprintf(setAdvertisingType, typ))
}

// GetBatteryInfo func
func (t *Ble) GetBatteryInfo() *BleResponse {
	return t.writeRead(getBatteryInfo)
}

// GetBaudRate func
func (t *Ble) GetBaudRate() *BleResponse {
	return t.writeRead(getBaudRate)
}

// SetBaudRate func
func (t *Ble) SetBaudRate(baud BleBaudRate) *BleResponse {
	return t.writeRead(fmt.Sprintf(setBaudRate, baud))
}

// GetCharacteristic func
func (t *Ble) GetCharacteristic() *BleResponse {
	return t.writeRead(getChar)
}

// SetCharacteristic func
func (t *Ble) SetCharacteristic(char int64) *BleResponse {
	return t.writeRead(fmt.Sprintf(setChar, t.hexTrick(char, 4)))
}

// ClearLastConnectedDevice func
func (t *Ble) ClearLastConnectedDevice() *BleResponse {
	return t.writeRead(clear)
}

//SetDeployMode func
func (t *Ble) SetDeployMode(mode BleDeployMode) *BleResponse {
	return t.writeRead(fmt.Sprintf(setDeployMode, mode))
}

// GetBeaconMode func
func (t *Ble) GetBeaconMode() *BleResponse {
	return t.writeRead(getBeacon)
}

// SetBeaconMode func
func (t *Ble) SetBeaconMode(mode BleBeaconMode) *BleResponse {
	return t.writeRead(fmt.Sprintf(setBeacon, mode))
}

// GetBeaconUUID func
func (t *Ble) GetBeaconUUID(part BleUUID) *BleResponse {
	return t.writeRead(fmt.Sprintf(getUUID, part))
}

// SetBeaconUUID func
func (t *Ble) SetBeaconUUID(part BleUUID, uuid int64) *BleResponse {
	return t.writeRead(fmt.Sprintf(setUUID, part, t.hexTrick(uuid, 8)))
}

// GetMajor func
func (t *Ble) GetMajor() *BleResponse {
	return t.writeRead(getMajor)
}

// SetMajor func
func (t *Ble) SetMajor(major int64) *BleResponse {
	return t.writeRead(fmt.Sprintf(setMajor, t.hexTrick(major, 4)))
}

// GetMinor func
func (t *Ble) GetMinor() *BleResponse {
	return t.writeRead(getMinor)
}

// SetMinor func
func (t *Ble) SetMinor(minor int64) *BleResponse {
	return t.writeRead(fmt.Sprintf(setMinor, t.hexTrick(minor, 4)))
}

// GetIPower func
func (t *Ble) GetIPower() *BleResponse {
	return t.writeRead(getIPower)
}

// SetIPower func /!\ DANGER /!\
func (t *Ble) SetIPower(power int64) *BleResponse {
	return t.writeRead(fmt.Sprintf(setIPower, t.hexTrick(power, 2)))
}

// GetMode func
func (t *Ble) GetMode() *BleResponse {
	return t.writeRead(getMode)
}

// SetMode func
func (t *Ble) SetMode(m BleMode) *BleResponse {
	return t.writeRead(fmt.Sprintf(setMode, int(m)))
}

// GetDeviceName func
func (t *Ble) GetDeviceName() *BleResponse {
	return t.writeRead(getDeviceName)
}

// SetDeviceName func
func (t *Ble) SetDeviceName(name string) *BleResponse {
	return t.writeRead(fmt.Sprintf(setDeviceName, name))
}

// GetParity func
func (t *Ble) GetParity() *BleResponse {
	return t.writeRead(getParity)
}

// SetParity func
func (t *Ble) SetParity(parity BleParity) *BleResponse {
	return t.writeRead(fmt.Sprintf(setParity, parity))
}

// GetPIO1Status func
func (t *Ble) GetPIO1Status() *BleResponse {
	return t.writeRead(getStatus)
}

// SetPIO1Status func
func (t *Ble) SetPIO1Status(s BlePIO) *BleResponse {
	return t.writeRead(fmt.Sprintf(setStatus, s))
}

// GetPIO func
func (t *Ble) GetPIO(pio BlePIOPin) *BleResponse {
	return t.writeRead(fmt.Sprintf(getPIO, pio))
}

// SetPIO func
func (t *Ble) SetPIO(pio BlePIOPin, value BlePIO) *BleResponse {
	return t.writeRead(fmt.Sprintf(setPIO, pio, int(value)))
}

// GetPin func
func (t *Ble) GetPin() *BleResponse {
	return t.writeRead(getPin)
}

// SetPin func
func (t *Ble) SetPin(pin int) *BleResponse {
	return t.writeRead(fmt.Sprintf(setPin, pin))
}

// GetPower func
func (t *Ble) GetPower() *BleResponse {
	return t.writeRead(getPower)
}

// SetPower func
func (t *Ble) SetPower(power BlePower) *BleResponse {
	return t.writeRead(fmt.Sprintf(setPower, power))
}

// FactoryReset func
func (t *Ble) FactoryReset() *BleResponse {
	return t.writeRead(factoryReset)
}

// Restart func
func (t *Ble) Restart() *BleResponse {
	return t.writeRead(restart)
}

// GetRole func
func (t *Ble) GetRole() *BleResponse {
	return t.writeRead(getRole)
}

// SetRole func
func (t *Ble) SetRole(role BleRole) *BleResponse {
	return t.writeRead(fmt.Sprintf(setRole, role))
}

// GetLastConnectedDeviceAddress func
func (t *Ble) GetLastConnectedDeviceAddress() *BleResponse {
	return t.writeRead(lastDeviceAddress)
}

// GetBondMode func
func (t *Ble) GetBondMode() *BleResponse {
	return t.writeRead(getBondMode)
}

// SetBondMode func
func (t *Ble) SetBondMode(mode BleBondMode) *BleResponse {
	return t.writeRead(fmt.Sprintf(setBondMode, int(mode)))
}

// GetDeviceUUID func
func (t *Ble) GetDeviceUUID() *BleResponse {
	return t.writeRead(getDeviceUUID)
}

// SetDeviceUUID func
func (t *Ble) SetDeviceUUID(uuid int64) *BleResponse {
	return t.writeRead(fmt.Sprintf(setDeviceUUID, t.hexTrick(uuid, 4)))
}

// GetSoftwareVersion func
func (t *Ble) GetSoftwareVersion() *BleResponse {
	return t.writeRead(getSoftwareVersion)
}

// hexTrick func
func (t *Ble) hexTrick(value int64, hexLen int) string {
	if hexLen == 2 {
		if value < 1 {
			value = 1
			fmt.Println("Warning: Value to low! (1-255)")
			fmt.Println("Value reset to: 1")
		} else if value > 255 {
			value = 255
			fmt.Println("Warning: value to high! (1-255)")
			fmt.Println("Value reset to: 255")
		}
	}
	if hexLen == 4 {
		if value < 1 {
			value = 1
			fmt.Println("Warning: Value to low! (1-65534)")
			fmt.Println("Value reset to: 1")
		} else if value > 65534 {
			value = 65534
			fmt.Println("Warning: value to high! (1-65534)")
			fmt.Println("Value reset to: 65534")
		}
	}
	if hexLen == 8 {
		if value < 1 {
			value = 1
			fmt.Println("Warning: Value to low (1-4294967294)")
			fmt.Println("Value reset to: 1")
		} else if value > 4294967294 {
			value = 4294967294
			fmt.Println("Warning: value to high! (1-4294967294)")
			fmt.Println("Value reset to: 4294967294")
		}
	}
	hex := strings.ToUpper(strconv.FormatInt(value, 16))
	for len(hex) < hexLen {
		tmp := hex
		hex = "0" + tmp
	}
	return hex
}

// writeRead func
func (t *Ble) writeRead(cmd string) *BleResponse {

	if _, err := t.fd.Write([]byte(cmd)); err != nil {
		return &BleResponse{
			Error: err,
		}
	}

	buff := make([]byte, 512)

	if _, err := t.fd.Read(buff); err != nil {
		return &BleResponse{
			Error: err,
		}
	}

	sep := strings.Split(string(buff), ":")
	if len(sep) == 2 {
		return &BleResponse{
			Result: string(buff),
			Param:  sep[1],
		}
	}
	return &BleResponse{
		Result: string(buff),
	}

}
