module github.com/gomidi/mid/adapters/rtmidiadapter/example

replace (
	github.com/gomidi/connect/imported/rtmidi => ../../imported/rtmidi
	github.com/gomidi/connect/rtmidiadapter => ../
)

require (
	github.com/gomidi/connect/imported/rtmidi v0.0.0-20180901202738-6d273e81f890
	github.com/gomidi/connect/rtmidiadapter v0.0.0-20180901202738-6d273e81f890
	github.com/gomidi/mid v0.15.0
)
