module github.com/gomidi/connect/rtmidiadapter

replace (
	github.com/gomidi/mid/imported/rtmidi => ../imported/rtmidi
)

require (
	github.com/gomidi/mid v0.0.0-20180818170814-7d37ca6b4142
	github.com/gomidi/connect/imported/rtmidi v0.0.0
)
