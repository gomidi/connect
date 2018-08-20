module github.com/gomidi/mid/adapters/rtmidiadapter/example

replace (
	github.com/gomidi/connect/imported/rtmidi => ../../imported/rtmidi
	github.com/gomidi/connect/rtmidiadapter => ../
)

require (
	github.com/gomidi/connect/imported/rtmidi v0.0.0-20180818202137-868136c52a76
	github.com/gomidi/connect/rtmidiadapter v0.0.0-20180818202137-868136c52a76
	github.com/gomidi/mid v0.9.5
)
