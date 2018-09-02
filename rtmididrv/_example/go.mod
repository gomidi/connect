module github.com/gomidi/connect/rtmididrv/_example

replace (
	github.com/gomidi/connect/imported/rtmidi => ../../imported/rtmidi
	github.com/gomidi/connect/rtmididrv => ../
)

require (
	github.com/gomidi/connect/imported/rtmidi v0.0.0-20180901203434-17b8e81ae4ad
	github.com/gomidi/mid v0.15.0
)
