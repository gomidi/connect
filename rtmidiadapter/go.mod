module github.com/gomidi/connect/rtmidiadapter

replace github.com/gomidi/mid/imported/rtmidi => ../imported/rtmidi

require (
	github.com/gomidi/connect/imported/rtmidi v0.0.0-20180827213430-b1136e6d4610
	github.com/gomidi/mid v0.10.0
)
