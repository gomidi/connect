module github.com/gomidi/connect/rtmidiadapter

replace github.com/gomidi/mid/imported/rtmidi => ../imported/rtmidi

require (
	github.com/gomidi/connect/imported/rtmidi v0.0.0-20180901202738-6d273e81f890
	github.com/gomidi/mid v0.15.0
)
