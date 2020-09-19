// Package audits provides the Chrome DevTools Protocol
// commands, types, and events for the Audits domain.
//
// Audits domain allows investigation of page violations and possible
// improvements.
//
// Generated by the cdproto-gen command.
package audits

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"
	"encoding/base64"

	"github.com/goconnectx/cdproto/cdp"
	"github.com/goconnectx/cdproto/network"
)

// GetEncodedResponseParams returns the response body and size if it were
// re-encoded with the specified settings. Only applies to images.
type GetEncodedResponseParams struct {
	RequestID network.RequestID          `json:"requestId"`          // Identifier of the network request to get content for.
	Encoding  GetEncodedResponseEncoding `json:"encoding"`           // The encoding to use.
	Quality   float64                    `json:"quality,omitempty"`  // The quality of the encoding (0-1). (defaults to 1)
	SizeOnly  bool                       `json:"sizeOnly,omitempty"` // Whether to only return the size information (defaults to false).
}

// GetEncodedResponse returns the response body and size if it were
// re-encoded with the specified settings. Only applies to images.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#method-getEncodedResponse
//
// parameters:
//   requestID - Identifier of the network request to get content for.
//   encoding - The encoding to use.
func GetEncodedResponse(requestID network.RequestID, encoding GetEncodedResponseEncoding) *GetEncodedResponseParams {
	return &GetEncodedResponseParams{
		RequestID: requestID,
		Encoding:  encoding,
	}
}

// WithQuality the quality of the encoding (0-1). (defaults to 1).
func (p GetEncodedResponseParams) WithQuality(quality float64) *GetEncodedResponseParams {
	p.Quality = quality
	return &p
}

// WithSizeOnly whether to only return the size information (defaults to
// false).
func (p GetEncodedResponseParams) WithSizeOnly(sizeOnly bool) *GetEncodedResponseParams {
	p.SizeOnly = sizeOnly
	return &p
}

// GetEncodedResponseReturns return values.
type GetEncodedResponseReturns struct {
	Body         string `json:"body,omitempty"`         // The encoded body as a base64 string. Omitted if sizeOnly is true.
	OriginalSize int64  `json:"originalSize,omitempty"` // Size before re-encoding.
	EncodedSize  int64  `json:"encodedSize,omitempty"`  // Size after re-encoding.
}

// Do executes Audits.getEncodedResponse against the provided context.
//
// returns:
//   body - The encoded body as a base64 string. Omitted if sizeOnly is true.
//   originalSize - Size before re-encoding.
//   encodedSize - Size after re-encoding.
func (p *GetEncodedResponseParams) Do(ctx context.Context) (body []byte, originalSize int64, encodedSize int64, err error) {
	// execute
	var res GetEncodedResponseReturns
	err = cdp.Execute(ctx, CommandGetEncodedResponse, p, &res)
	if err != nil {
		return nil, 0, 0, err
	}

	// decode
	var dec []byte
	dec, err = base64.StdEncoding.DecodeString(res.Body)
	if err != nil {
		return nil, 0, 0, err
	}
	return dec, res.OriginalSize, res.EncodedSize, nil
}

// DisableParams disables issues domain, prevents further issues from being
// reported to the client.
type DisableParams struct{}

// Disable disables issues domain, prevents further issues from being
// reported to the client.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#method-disable
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Audits.disable against the provided context.
func (p *DisableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDisable, nil, nil)
}

// EnableParams enables issues domain, sends the issues collected so far to
// the client by means of the issueAdded event.
type EnableParams struct{}

// Enable enables issues domain, sends the issues collected so far to the
// client by means of the issueAdded event.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#method-enable
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes Audits.enable against the provided context.
func (p *EnableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandEnable, nil, nil)
}

// Command names.
const (
	CommandGetEncodedResponse = "Audits.getEncodedResponse"
	CommandDisable            = "Audits.disable"
	CommandEnable             = "Audits.enable"
)
