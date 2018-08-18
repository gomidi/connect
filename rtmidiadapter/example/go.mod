module github.com/gomidi/mid/adapters/rtmidiadapter/example

replace (
	github.com/gomidi/connect/rtmidiadapter => ../
	github.com/gomidi/connect/imported/rtmidi => ../../../imported/rtmidi
)

require (
	github.com/gomidi/mid v0.0.0-20180818170814-7d37ca6b4142
	github.com/gomidi/midi v1.1.4
)
