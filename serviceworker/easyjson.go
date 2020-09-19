// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package serviceworker

import (
	json "encoding/json"
	target "github.com/goconnectx/cdproto/target"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker(in *jlexer.Lexer, out *Version) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "versionId":
			out.VersionID = string(in.String())
		case "registrationId":
			out.RegistrationID = RegistrationID(in.String())
		case "scriptURL":
			out.ScriptURL = string(in.String())
		case "runningStatus":
			(out.RunningStatus).UnmarshalEasyJSON(in)
		case "status":
			(out.Status).UnmarshalEasyJSON(in)
		case "scriptLastModified":
			out.ScriptLastModified = float64(in.Float64())
		case "scriptResponseTime":
			out.ScriptResponseTime = float64(in.Float64())
		case "controlledClients":
			if in.IsNull() {
				in.Skip()
				out.ControlledClients = nil
			} else {
				in.Delim('[')
				if out.ControlledClients == nil {
					if !in.IsDelim(']') {
						out.ControlledClients = make([]target.ID, 0, 4)
					} else {
						out.ControlledClients = []target.ID{}
					}
				} else {
					out.ControlledClients = (out.ControlledClients)[:0]
				}
				for !in.IsDelim(']') {
					var v1 target.ID
					v1 = target.ID(in.String())
					out.ControlledClients = append(out.ControlledClients, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "targetId":
			out.TargetID = target.ID(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker(out *jwriter.Writer, in Version) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"versionId\":"
		out.RawString(prefix[1:])
		out.String(string(in.VersionID))
	}
	{
		const prefix string = ",\"registrationId\":"
		out.RawString(prefix)
		out.String(string(in.RegistrationID))
	}
	{
		const prefix string = ",\"scriptURL\":"
		out.RawString(prefix)
		out.String(string(in.ScriptURL))
	}
	{
		const prefix string = ",\"runningStatus\":"
		out.RawString(prefix)
		(in.RunningStatus).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		(in.Status).MarshalEasyJSON(out)
	}
	if in.ScriptLastModified != 0 {
		const prefix string = ",\"scriptLastModified\":"
		out.RawString(prefix)
		out.Float64(float64(in.ScriptLastModified))
	}
	if in.ScriptResponseTime != 0 {
		const prefix string = ",\"scriptResponseTime\":"
		out.RawString(prefix)
		out.Float64(float64(in.ScriptResponseTime))
	}
	if len(in.ControlledClients) != 0 {
		const prefix string = ",\"controlledClients\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v2, v3 := range in.ControlledClients {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	if in.TargetID != "" {
		const prefix string = ",\"targetId\":"
		out.RawString(prefix)
		out.String(string(in.TargetID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Version) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Version) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Version) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Version) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker1(in *jlexer.Lexer, out *UpdateRegistrationParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "scopeURL":
			out.ScopeURL = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker1(out *jwriter.Writer, in UpdateRegistrationParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"scopeURL\":"
		out.RawString(prefix[1:])
		out.String(string(in.ScopeURL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdateRegistrationParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateRegistrationParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdateRegistrationParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateRegistrationParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker1(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker2(in *jlexer.Lexer, out *UnregisterParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "scopeURL":
			out.ScopeURL = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker2(out *jwriter.Writer, in UnregisterParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"scopeURL\":"
		out.RawString(prefix[1:])
		out.String(string(in.ScopeURL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UnregisterParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UnregisterParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UnregisterParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UnregisterParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker2(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker3(in *jlexer.Lexer, out *StopWorkerParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "versionId":
			out.VersionID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker3(out *jwriter.Writer, in StopWorkerParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"versionId\":"
		out.RawString(prefix[1:])
		out.String(string(in.VersionID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StopWorkerParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StopWorkerParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StopWorkerParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StopWorkerParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker3(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker4(in *jlexer.Lexer, out *StopAllWorkersParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker4(out *jwriter.Writer, in StopAllWorkersParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StopAllWorkersParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StopAllWorkersParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StopAllWorkersParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StopAllWorkersParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker4(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker5(in *jlexer.Lexer, out *StartWorkerParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "scopeURL":
			out.ScopeURL = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker5(out *jwriter.Writer, in StartWorkerParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"scopeURL\":"
		out.RawString(prefix[1:])
		out.String(string(in.ScopeURL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StartWorkerParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StartWorkerParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StartWorkerParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StartWorkerParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker5(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker6(in *jlexer.Lexer, out *SkipWaitingParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "scopeURL":
			out.ScopeURL = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker6(out *jwriter.Writer, in SkipWaitingParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"scopeURL\":"
		out.RawString(prefix[1:])
		out.String(string(in.ScopeURL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SkipWaitingParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SkipWaitingParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SkipWaitingParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SkipWaitingParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker6(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker7(in *jlexer.Lexer, out *SetForceUpdateOnPageLoadParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "forceUpdateOnPageLoad":
			out.ForceUpdateOnPageLoad = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker7(out *jwriter.Writer, in SetForceUpdateOnPageLoadParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"forceUpdateOnPageLoad\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.ForceUpdateOnPageLoad))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetForceUpdateOnPageLoadParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetForceUpdateOnPageLoadParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetForceUpdateOnPageLoadParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetForceUpdateOnPageLoadParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker7(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker8(in *jlexer.Lexer, out *Registration) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "registrationId":
			out.RegistrationID = RegistrationID(in.String())
		case "scopeURL":
			out.ScopeURL = string(in.String())
		case "isDeleted":
			out.IsDeleted = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker8(out *jwriter.Writer, in Registration) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"registrationId\":"
		out.RawString(prefix[1:])
		out.String(string(in.RegistrationID))
	}
	{
		const prefix string = ",\"scopeURL\":"
		out.RawString(prefix)
		out.String(string(in.ScopeURL))
	}
	{
		const prefix string = ",\"isDeleted\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsDeleted))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Registration) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Registration) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Registration) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Registration) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker8(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker9(in *jlexer.Lexer, out *InspectWorkerParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "versionId":
			out.VersionID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker9(out *jwriter.Writer, in InspectWorkerParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"versionId\":"
		out.RawString(prefix[1:])
		out.String(string(in.VersionID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v InspectWorkerParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v InspectWorkerParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *InspectWorkerParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *InspectWorkerParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker9(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker10(in *jlexer.Lexer, out *EventWorkerVersionUpdated) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "versions":
			if in.IsNull() {
				in.Skip()
				out.Versions = nil
			} else {
				in.Delim('[')
				if out.Versions == nil {
					if !in.IsDelim(']') {
						out.Versions = make([]*Version, 0, 8)
					} else {
						out.Versions = []*Version{}
					}
				} else {
					out.Versions = (out.Versions)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *Version
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(Version)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Versions = append(out.Versions, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker10(out *jwriter.Writer, in EventWorkerVersionUpdated) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"versions\":"
		out.RawString(prefix[1:])
		if in.Versions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Versions {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventWorkerVersionUpdated) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventWorkerVersionUpdated) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventWorkerVersionUpdated) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventWorkerVersionUpdated) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker10(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker11(in *jlexer.Lexer, out *EventWorkerRegistrationUpdated) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "registrations":
			if in.IsNull() {
				in.Skip()
				out.Registrations = nil
			} else {
				in.Delim('[')
				if out.Registrations == nil {
					if !in.IsDelim(']') {
						out.Registrations = make([]*Registration, 0, 8)
					} else {
						out.Registrations = []*Registration{}
					}
				} else {
					out.Registrations = (out.Registrations)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *Registration
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(Registration)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					out.Registrations = append(out.Registrations, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker11(out *jwriter.Writer, in EventWorkerRegistrationUpdated) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"registrations\":"
		out.RawString(prefix[1:])
		if in.Registrations == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Registrations {
				if v8 > 0 {
					out.RawByte(',')
				}
				if v9 == nil {
					out.RawString("null")
				} else {
					(*v9).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventWorkerRegistrationUpdated) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventWorkerRegistrationUpdated) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventWorkerRegistrationUpdated) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventWorkerRegistrationUpdated) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker11(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker12(in *jlexer.Lexer, out *EventWorkerErrorReported) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "errorMessage":
			if in.IsNull() {
				in.Skip()
				out.ErrorMessage = nil
			} else {
				if out.ErrorMessage == nil {
					out.ErrorMessage = new(ErrorMessage)
				}
				(*out.ErrorMessage).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker12(out *jwriter.Writer, in EventWorkerErrorReported) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"errorMessage\":"
		out.RawString(prefix[1:])
		if in.ErrorMessage == nil {
			out.RawString("null")
		} else {
			(*in.ErrorMessage).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventWorkerErrorReported) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventWorkerErrorReported) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventWorkerErrorReported) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventWorkerErrorReported) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker12(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker13(in *jlexer.Lexer, out *ErrorMessage) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "errorMessage":
			out.ErrorMessage = string(in.String())
		case "registrationId":
			out.RegistrationID = RegistrationID(in.String())
		case "versionId":
			out.VersionID = string(in.String())
		case "sourceURL":
			out.SourceURL = string(in.String())
		case "lineNumber":
			out.LineNumber = int64(in.Int64())
		case "columnNumber":
			out.ColumnNumber = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker13(out *jwriter.Writer, in ErrorMessage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"errorMessage\":"
		out.RawString(prefix[1:])
		out.String(string(in.ErrorMessage))
	}
	{
		const prefix string = ",\"registrationId\":"
		out.RawString(prefix)
		out.String(string(in.RegistrationID))
	}
	{
		const prefix string = ",\"versionId\":"
		out.RawString(prefix)
		out.String(string(in.VersionID))
	}
	{
		const prefix string = ",\"sourceURL\":"
		out.RawString(prefix)
		out.String(string(in.SourceURL))
	}
	{
		const prefix string = ",\"lineNumber\":"
		out.RawString(prefix)
		out.Int64(int64(in.LineNumber))
	}
	{
		const prefix string = ",\"columnNumber\":"
		out.RawString(prefix)
		out.Int64(int64(in.ColumnNumber))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ErrorMessage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ErrorMessage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ErrorMessage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ErrorMessage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker13(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker14(in *jlexer.Lexer, out *EnableParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker14(out *jwriter.Writer, in EnableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EnableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker14(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EnableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker14(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EnableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker14(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EnableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker14(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker15(in *jlexer.Lexer, out *DispatchSyncEventParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "origin":
			out.Origin = string(in.String())
		case "registrationId":
			out.RegistrationID = RegistrationID(in.String())
		case "tag":
			out.Tag = string(in.String())
		case "lastChance":
			out.LastChance = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker15(out *jwriter.Writer, in DispatchSyncEventParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"origin\":"
		out.RawString(prefix[1:])
		out.String(string(in.Origin))
	}
	{
		const prefix string = ",\"registrationId\":"
		out.RawString(prefix)
		out.String(string(in.RegistrationID))
	}
	{
		const prefix string = ",\"tag\":"
		out.RawString(prefix)
		out.String(string(in.Tag))
	}
	{
		const prefix string = ",\"lastChance\":"
		out.RawString(prefix)
		out.Bool(bool(in.LastChance))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DispatchSyncEventParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker15(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DispatchSyncEventParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker15(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DispatchSyncEventParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker15(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DispatchSyncEventParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker15(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker16(in *jlexer.Lexer, out *DispatchPeriodicSyncEventParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "origin":
			out.Origin = string(in.String())
		case "registrationId":
			out.RegistrationID = RegistrationID(in.String())
		case "tag":
			out.Tag = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker16(out *jwriter.Writer, in DispatchPeriodicSyncEventParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"origin\":"
		out.RawString(prefix[1:])
		out.String(string(in.Origin))
	}
	{
		const prefix string = ",\"registrationId\":"
		out.RawString(prefix)
		out.String(string(in.RegistrationID))
	}
	{
		const prefix string = ",\"tag\":"
		out.RawString(prefix)
		out.String(string(in.Tag))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DispatchPeriodicSyncEventParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker16(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DispatchPeriodicSyncEventParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker16(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DispatchPeriodicSyncEventParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker16(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DispatchPeriodicSyncEventParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker16(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker17(in *jlexer.Lexer, out *DisableParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker17(out *jwriter.Writer, in DisableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DisableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker17(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DisableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker17(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DisableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker17(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DisableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker17(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker18(in *jlexer.Lexer, out *DeliverPushMessageParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "origin":
			out.Origin = string(in.String())
		case "registrationId":
			out.RegistrationID = RegistrationID(in.String())
		case "data":
			out.Data = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker18(out *jwriter.Writer, in DeliverPushMessageParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"origin\":"
		out.RawString(prefix[1:])
		out.String(string(in.Origin))
	}
	{
		const prefix string = ",\"registrationId\":"
		out.RawString(prefix)
		out.String(string(in.RegistrationID))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		out.String(string(in.Data))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DeliverPushMessageParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker18(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeliverPushMessageParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoServiceworker18(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeliverPushMessageParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker18(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeliverPushMessageParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoServiceworker18(l, v)
}
