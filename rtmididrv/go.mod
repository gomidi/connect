module github.com/gomidi/connect/rtmididrv

replace (
	github.com/gomidi/connect => ../
	github.com/gomidi/mid/imported/rtmidi => ../imported/rtmidi
)

require (
	github.com/gomidi/connect v0.7.0
	github.com/gomidi/connect/imported/rtmidi v0.0.0-20180901202738-6d273e81f890
)
