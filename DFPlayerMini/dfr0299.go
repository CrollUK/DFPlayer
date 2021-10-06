/* Code for setup and code values taken from https://wiki.dfrobot.com/DFPlayer_Mini_SKU_DFR0299
The DFPlayer Mini MP3 Player For Arduino is a small and low price MP3 module with a simplified output directly to the speaker.
The module can be used as a stand-alone module with attached battery, speaker and push buttons or used in
combination with an Arduino UNO or any other microcontroller with RX/TX capabilities.

@file DFPlayerMini.go
@brief DFPlayer - A Mini MP3 Player From DFRobot
@driver file for DFRobot's DFPlayer

@copyright	[DFRobot]( http://www.dfrobot.com ), 2016
@copyright	GNU Lesser General Public License
@author [Angelo](Angelo.qiao@dfrobot.com)
* @version  V1.0.3
* @date  2016-12-0

TinyGo Conversion ATTEMPT
@author [nigel](croll.uk@gmail.com)
@version  V0.0.1
@date  2021-10-06
*/

package DFPlayerMini

import (
	"machine"
	"time"
)

var (
	uart = machine.UART1
	tx   = machine.UART_TX_PIN
	rx   = machine.UART_RX_PIN
	baud = uint32(9600)
)

var (
	StartByte     = startByte
	VersionByte   = versionByte
	CommandLength = commandLength
	EndByte       = endByte
	Acknowledge   = acknowledge // Returns info with command 0x41 [0x01: info, 0x00: no info]
	ACTIVATED     = false
	isPlaying     = false
)

func lowByte(x int16) byte {
	x = x >> 8
	return byte(x)
}

func highByte(x byte) byte {
	x = x & 0xFF
	return x
}

func executeCmd(cmd, par1, par2 byte) {
	checksum := VersionByte + CommandLength + cmd + Acknowledge + par1 + par2
	// Build the command line
	commandLine := []byte{StartByte, VersionByte, CommandLength, cmd, Acknowledge, par1, par2,
		highByte(checksum), lowByte(int16(checksum)), EndByte}
	//example: commanddata := []byte{0x7E, 0xFF, 0x06, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0xEF}
	_, err := uart.Write(commandLine)
	if err != nil {
		return
	}
}

func Setup() {
	err := uart.Configure(machine.UARTConfig{BaudRate: baud, TX: tx, RX: rx})
	if err != nil {
		panic(err)
	}
	// Set serial comms time-out 500ms
	// Set Volume value (0~30)
	// Set EQ
	// Set file storage device
}

func Play() {
	executeCmd(play, 0, 0)
	time.Sleep(time.Second / 2)
}

func PlayFirst() {
	executeCmd(sendInit, 0, 0)
	time.Sleep(time.Second / 2)
	executeCmd(repeatPlay, 0, 1)
	time.Sleep(time.Second / 2)
}

func PlayNext() {
	executeCmd(next, 0, 1)
	time.Sleep(time.Second / 2)
}

func Pause() {
	executeCmd(pause, 0, 0)
	time.Sleep(time.Second / 2)
}

func SetStorageDevice() {
	executeCmd(playbackSource, byte(deviceSD), 0)
}

func SetVolume(i int) {
	if i > 48 {
		i = 48
	}
	executeCmd(volume, 0, byte(i)) // Set the volume 0x00~0x30 | 0-48
	time.Sleep(2 * time.Second)
}
