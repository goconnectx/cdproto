package applicationcache

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"github.com/goconnectx/cdproto/cdp"
)

// EventApplicationCacheStatusUpdated [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache#event-applicationCacheStatusUpdated
type EventApplicationCacheStatusUpdated struct {
	FrameID     cdp.FrameID `json:"frameId"`     // Identifier of the frame containing document whose application cache updated status.
	ManifestURL string      `json:"manifestURL"` // Manifest URL.
	Status      int64       `json:"status"`      // Updated application cache status.
}

// EventNetworkStateUpdated [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache#event-networkStateUpdated
type EventNetworkStateUpdated struct {
	IsNowOnline bool `json:"isNowOnline"`
}
