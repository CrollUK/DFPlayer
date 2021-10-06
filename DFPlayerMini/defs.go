package DFPlayerMini

const (
	// DFPlayer command string sentence
	startByte     byte = 0x7E
	commandLength byte = 0x06
	versionByte   byte = 0xFF
	endByte       byte = 0xEF
	acknowledge   byte = 0x00 // Returns info with command 0x41 [0x01: info, 0x00: no info]
)

const (
	// PLAYER Serial Control Commands
	_              byte = iota
	next                // Play next track
	previous            // Play previous track
	track               // Specify trackk number 0-2999
	volumeUp            // Increase volume
	volumeDown          // Decrease volume
	volume              // Specify volume level 0-30
	eq                  // Specify EQ 0-5 Normal/Pop/Rock/Jazz/Classic/Bass
	playbackMode        // 0-3 Repeat/Folder Repeat/Single Repeat/Random
	playbackSource      // 0-4 U/TF/AUX/SLEEP/FLASH
	standby             // Enter into standby - low power loss
	normal              // Normal working
	reset               // Reset module
	play                // Start playback
	pause               // Pause playback
	folder              // Specify folder to playback 1~10(needs set by user)
	volumeAdjust        // Volume adjust set - DH=1: Open Vol adjust|DL: set vol gain 0-31
	repeatPlay          // Repeat play - 1:start repeat play 0:stop play
)

const (
	// Serial Query Commands
	stay1               byte = 0x3C // STAY - I've no idea!
	stay2               byte = 0x3D // STAY - I've no idea!
	stay3               byte = 0x3E // STAY - I've no idea!
	sendInit            byte = 0x3F // Send initialisation parameters|0-0x0F(each bit represents one device of the low-four bits)
	returnErr           byte = 0x40 // Returns an error, request retransmission
	reply               byte = 0x41
	status              byte = 0x42 // Query the current status
	currentVolume       byte = 0x43 // Query the current volume level
	currentEq           byte = 0x44 // Query the current EQ setting
	currentPlaybackMode byte = 0x45 // Query the current playback mode
	firmwareVer         byte = 0x46 // Query the current firmware version
	totalFilesOnTfCard  byte = 0x47
	totalFilesOnUDisk   byte = 0x48
	totalFilesOnFlash   byte = 0x49
	keepOn              byte = 0x4A // Not sure!
	tfCardTrack         byte = 0x4B // Query current track of TF card
	uDiskTrack          byte = 0x4C // Query current track of U-Disk
	flashTrack          byte = 0x4D // Query current track of Flash
)

const (
	// Built-in Equaliser
	eqNormal int8 = iota
	eqPop
	eqRocK
	eqJazz
	eqClassic
	eqBass
)

const (
	// DFPlayer Device
	_ int8 = iota
	deviceUDisk
	deviceSD
	deviceAux
	deviceSleep
	deviceFlash
)

const (
	timeOut int8 = iota
	wrongStack
	dfpCardInserted
)