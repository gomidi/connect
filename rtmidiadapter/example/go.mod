module github.com/gomidi/mid/adapters/rtmidiadapter/example

replace (
	github.com/gomidi/connect/imported/rtmidi => ../../../imported/rtmidi
	github.com/gomidi/connect/rtmidiadapter => ../
)

require github.com/gomidi/mid v0.8.1 // indirect
