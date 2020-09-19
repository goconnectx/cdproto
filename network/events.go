package network

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"github.com/goconnectx/cdproto/cdp"
)

// EventDataReceived fired when data chunk was received over the network.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-dataReceived
type EventDataReceived struct {
	RequestID         RequestID          `json:"requestId"`         // Request identifier.
	Timestamp         *cdp.MonotonicTime `json:"timestamp"`         // Timestamp.
	DataLength        int64              `json:"dataLength"`        // Data chunk length.
	EncodedDataLength int64              `json:"encodedDataLength"` // Actual bytes received (might be less than dataLength for compressed encodings).
}

// EventEventSourceMessageReceived fired when EventSource message is
// received.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-eventSourceMessageReceived
type EventEventSourceMessageReceived struct {
	RequestID RequestID          `json:"requestId"` // Request identifier.
	Timestamp *cdp.MonotonicTime `json:"timestamp"` // Timestamp.
	EventName string             `json:"eventName"` // Message type.
	EventID   string             `json:"eventId"`   // Message identifier.
	Data      string             `json:"data"`      // Message content.
}

// EventLoadingFailed fired when HTTP request has failed to load.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-loadingFailed
type EventLoadingFailed struct {
	RequestID     RequestID          `json:"requestId"`               // Request identifier.
	Timestamp     *cdp.MonotonicTime `json:"timestamp"`               // Timestamp.
	Type          ResourceType       `json:"type"`                    // Resource type.
	ErrorText     string             `json:"errorText"`               // User friendly error message.
	Canceled      bool               `json:"canceled,omitempty"`      // True if loading was canceled.
	BlockedReason BlockedReason      `json:"blockedReason,omitempty"` // The reason why loading was blocked, if any.
}

// EventLoadingFinished fired when HTTP request has finished loading.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-loadingFinished
type EventLoadingFinished struct {
	RequestID                RequestID          `json:"requestId"`                          // Request identifier.
	Timestamp                *cdp.MonotonicTime `json:"timestamp"`                          // Timestamp.
	EncodedDataLength        float64            `json:"encodedDataLength"`                  // Total number of bytes received for this request.
	ShouldReportCorbBlocking bool               `json:"shouldReportCorbBlocking,omitempty"` // Set when 1) response was blocked by Cross-Origin Read Blocking and also 2) this needs to be reported to the DevTools console.
}

// EventRequestServedFromCache fired if request ended up loading from cache.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-requestServedFromCache
type EventRequestServedFromCache struct {
	RequestID RequestID `json:"requestId"` // Request identifier.
}

// EventRequestWillBeSent fired when page is about to send HTTP request.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-requestWillBeSent
type EventRequestWillBeSent struct {
	RequestID        RequestID           `json:"requestId"`                  // Request identifier.
	LoaderID         cdp.LoaderID        `json:"loaderId"`                   // Loader identifier. Empty string if the request is fetched from worker.
	DocumentURL      string              `json:"documentURL"`                // URL of the document this request is loaded for.
	Request          *Request            `json:"request"`                    // Request data.
	Timestamp        *cdp.MonotonicTime  `json:"timestamp"`                  // Timestamp.
	WallTime         *cdp.TimeSinceEpoch `json:"wallTime"`                   // Timestamp.
	Initiator        *Initiator          `json:"initiator"`                  // Request initiator.
	RedirectResponse *Response           `json:"redirectResponse,omitempty"` // Redirect response data.
	Type             ResourceType        `json:"type,omitempty"`             // Type of this resource.
	FrameID          cdp.FrameID         `json:"frameId,omitempty"`          // Frame identifier.
	HasUserGesture   bool                `json:"hasUserGesture,omitempty"`   // Whether the request is initiated by a user gesture. Defaults to false.
}

// EventResourceChangedPriority fired when resource loading priority is
// changed.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-resourceChangedPriority
type EventResourceChangedPriority struct {
	RequestID   RequestID          `json:"requestId"`   // Request identifier.
	NewPriority ResourcePriority   `json:"newPriority"` // New priority
	Timestamp   *cdp.MonotonicTime `json:"timestamp"`   // Timestamp.
}

// EventSignedExchangeReceived fired when a signed exchange was received over
// the network.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-signedExchangeReceived
type EventSignedExchangeReceived struct {
	RequestID RequestID           `json:"requestId"` // Request identifier.
	Info      *SignedExchangeInfo `json:"info"`      // Information about the signed exchange response.
}

// EventResponseReceived fired when HTTP response is available.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-responseReceived
type EventResponseReceived struct {
	RequestID RequestID          `json:"requestId"`         // Request identifier.
	LoaderID  cdp.LoaderID       `json:"loaderId"`          // Loader identifier. Empty string if the request is fetched from worker.
	Timestamp *cdp.MonotonicTime `json:"timestamp"`         // Timestamp.
	Type      ResourceType       `json:"type"`              // Resource type.
	Response  *Response          `json:"response"`          // Response data.
	FrameID   cdp.FrameID        `json:"frameId,omitempty"` // Frame identifier.
}

// EventWebSocketClosed fired when WebSocket is closed.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-webSocketClosed
type EventWebSocketClosed struct {
	RequestID RequestID          `json:"requestId"` // Request identifier.
	Timestamp *cdp.MonotonicTime `json:"timestamp"` // Timestamp.
}

// EventWebSocketCreated fired upon WebSocket creation.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-webSocketCreated
type EventWebSocketCreated struct {
	RequestID RequestID  `json:"requestId"`           // Request identifier.
	URL       string     `json:"url"`                 // WebSocket request URL.
	Initiator *Initiator `json:"initiator,omitempty"` // Request initiator.
}

// EventWebSocketFrameError fired when WebSocket message error occurs.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-webSocketFrameError
type EventWebSocketFrameError struct {
	RequestID    RequestID          `json:"requestId"`    // Request identifier.
	Timestamp    *cdp.MonotonicTime `json:"timestamp"`    // Timestamp.
	ErrorMessage string             `json:"errorMessage"` // WebSocket error message.
}

// EventWebSocketFrameReceived fired when WebSocket message is received.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-webSocketFrameReceived
type EventWebSocketFrameReceived struct {
	RequestID RequestID          `json:"requestId"` // Request identifier.
	Timestamp *cdp.MonotonicTime `json:"timestamp"` // Timestamp.
	Response  *WebSocketFrame    `json:"response"`  // WebSocket response data.
}

// EventWebSocketFrameSent fired when WebSocket message is sent.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-webSocketFrameSent
type EventWebSocketFrameSent struct {
	RequestID RequestID          `json:"requestId"` // Request identifier.
	Timestamp *cdp.MonotonicTime `json:"timestamp"` // Timestamp.
	Response  *WebSocketFrame    `json:"response"`  // WebSocket response data.
}

// EventWebSocketHandshakeResponseReceived fired when WebSocket handshake
// response becomes available.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-webSocketHandshakeResponseReceived
type EventWebSocketHandshakeResponseReceived struct {
	RequestID RequestID          `json:"requestId"` // Request identifier.
	Timestamp *cdp.MonotonicTime `json:"timestamp"` // Timestamp.
	Response  *WebSocketResponse `json:"response"`  // WebSocket response data.
}

// EventWebSocketWillSendHandshakeRequest fired when WebSocket is about to
// initiate handshake.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-webSocketWillSendHandshakeRequest
type EventWebSocketWillSendHandshakeRequest struct {
	RequestID RequestID           `json:"requestId"` // Request identifier.
	Timestamp *cdp.MonotonicTime  `json:"timestamp"` // Timestamp.
	WallTime  *cdp.TimeSinceEpoch `json:"wallTime"`  // UTC Timestamp.
	Request   *WebSocketRequest   `json:"request"`   // WebSocket request data.
}

// EventRequestWillBeSentExtraInfo fired when additional information about a
// requestWillBeSent event is available from the network stack. Not every
// requestWillBeSent event will have an additional requestWillBeSentExtraInfo
// fired for it, and there is no guarantee whether requestWillBeSent or
// requestWillBeSentExtraInfo will be fired first for the same request.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-requestWillBeSentExtraInfo
type EventRequestWillBeSentExtraInfo struct {
	RequestID         RequestID                  `json:"requestId"`         // Request identifier. Used to match this information to an existing requestWillBeSent event.
	AssociatedCookies []*BlockedCookieWithReason `json:"associatedCookies"` // A list of cookies potentially associated to the requested URL. This includes both cookies sent with the request and the ones not sent; the latter are distinguished by having blockedReason field set.
	Headers           Headers                    `json:"headers"`           // Raw request headers as they will be sent over the wire.
}

// EventResponseReceivedExtraInfo fired when additional information about a
// responseReceived event is available from the network stack. Not every
// responseReceived event will have an additional responseReceivedExtraInfo for
// it, and responseReceivedExtraInfo may be fired before or after
// responseReceived.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-responseReceivedExtraInfo
type EventResponseReceivedExtraInfo struct {
	RequestID      RequestID                     `json:"requestId"`             // Request identifier. Used to match this information to another responseReceived event.
	BlockedCookies []*BlockedSetCookieWithReason `json:"blockedCookies"`        // A list of cookies which were not stored from the response along with the corresponding reasons for blocking. The cookies here may not be valid due to syntax errors, which are represented by the invalid cookie line string instead of a proper cookie.
	Headers        Headers                       `json:"headers"`               // Raw response headers as they were received over the wire.
	HeadersText    string                        `json:"headersText,omitempty"` // Raw response header text as it was received over the wire. The raw text may not always be available, such as in the case of HTTP/2 or QUIC.
}