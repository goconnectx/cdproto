// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package accessibility

import (
	json "encoding/json"
	runtime "github.com/goconnectx/cdproto/runtime"
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

func easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility(in *jlexer.Lexer, out *ValueSource) {
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
		case "type":
			(out.Type).UnmarshalEasyJSON(in)
		case "value":
			if in.IsNull() {
				in.Skip()
				out.Value = nil
			} else {
				if out.Value == nil {
					out.Value = new(Value)
				}
				(*out.Value).UnmarshalEasyJSON(in)
			}
		case "attribute":
			out.Attribute = string(in.String())
		case "attributeValue":
			if in.IsNull() {
				in.Skip()
				out.AttributeValue = nil
			} else {
				if out.AttributeValue == nil {
					out.AttributeValue = new(Value)
				}
				(*out.AttributeValue).UnmarshalEasyJSON(in)
			}
		case "superseded":
			out.Superseded = bool(in.Bool())
		case "nativeSource":
			(out.NativeSource).UnmarshalEasyJSON(in)
		case "nativeSourceValue":
			if in.IsNull() {
				in.Skip()
				out.NativeSourceValue = nil
			} else {
				if out.NativeSourceValue == nil {
					out.NativeSourceValue = new(Value)
				}
				(*out.NativeSourceValue).UnmarshalEasyJSON(in)
			}
		case "invalid":
			out.Invalid = bool(in.Bool())
		case "invalidReason":
			out.InvalidReason = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility(out *jwriter.Writer, in ValueSource) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix[1:])
		(in.Type).MarshalEasyJSON(out)
	}
	if in.Value != nil {
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		(*in.Value).MarshalEasyJSON(out)
	}
	if in.Attribute != "" {
		const prefix string = ",\"attribute\":"
		out.RawString(prefix)
		out.String(string(in.Attribute))
	}
	if in.AttributeValue != nil {
		const prefix string = ",\"attributeValue\":"
		out.RawString(prefix)
		(*in.AttributeValue).MarshalEasyJSON(out)
	}
	if in.Superseded {
		const prefix string = ",\"superseded\":"
		out.RawString(prefix)
		out.Bool(bool(in.Superseded))
	}
	if in.NativeSource != "" {
		const prefix string = ",\"nativeSource\":"
		out.RawString(prefix)
		(in.NativeSource).MarshalEasyJSON(out)
	}
	if in.NativeSourceValue != nil {
		const prefix string = ",\"nativeSourceValue\":"
		out.RawString(prefix)
		(*in.NativeSourceValue).MarshalEasyJSON(out)
	}
	if in.Invalid {
		const prefix string = ",\"invalid\":"
		out.RawString(prefix)
		out.Bool(bool(in.Invalid))
	}
	if in.InvalidReason != "" {
		const prefix string = ",\"invalidReason\":"
		out.RawString(prefix)
		out.String(string(in.InvalidReason))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ValueSource) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ValueSource) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ValueSource) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ValueSource) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility1(in *jlexer.Lexer, out *Value) {
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
		case "type":
			(out.Type).UnmarshalEasyJSON(in)
		case "value":
			(out.Value).UnmarshalEasyJSON(in)
		case "relatedNodes":
			if in.IsNull() {
				in.Skip()
				out.RelatedNodes = nil
			} else {
				in.Delim('[')
				if out.RelatedNodes == nil {
					if !in.IsDelim(']') {
						out.RelatedNodes = make([]*RelatedNode, 0, 8)
					} else {
						out.RelatedNodes = []*RelatedNode{}
					}
				} else {
					out.RelatedNodes = (out.RelatedNodes)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *RelatedNode
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(RelatedNode)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.RelatedNodes = append(out.RelatedNodes, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "sources":
			if in.IsNull() {
				in.Skip()
				out.Sources = nil
			} else {
				in.Delim('[')
				if out.Sources == nil {
					if !in.IsDelim(']') {
						out.Sources = make([]*ValueSource, 0, 8)
					} else {
						out.Sources = []*ValueSource{}
					}
				} else {
					out.Sources = (out.Sources)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *ValueSource
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(ValueSource)
						}
						(*v2).UnmarshalEasyJSON(in)
					}
					out.Sources = append(out.Sources, v2)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility1(out *jwriter.Writer, in Value) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix[1:])
		(in.Type).MarshalEasyJSON(out)
	}
	if (in.Value).IsDefined() {
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		(in.Value).MarshalEasyJSON(out)
	}
	if len(in.RelatedNodes) != 0 {
		const prefix string = ",\"relatedNodes\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v3, v4 := range in.RelatedNodes {
				if v3 > 0 {
					out.RawByte(',')
				}
				if v4 == nil {
					out.RawString("null")
				} else {
					(*v4).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.Sources) != 0 {
		const prefix string = ",\"sources\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v5, v6 := range in.Sources {
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
func (v Value) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Value) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Value) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Value) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility1(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility2(in *jlexer.Lexer, out *RelatedNode) {
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
		case "backendDOMNodeId":
			(out.BackendDOMNodeID).UnmarshalEasyJSON(in)
		case "idref":
			out.Idref = string(in.String())
		case "text":
			out.Text = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility2(out *jwriter.Writer, in RelatedNode) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"backendDOMNodeId\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.BackendDOMNodeID))
	}
	if in.Idref != "" {
		const prefix string = ",\"idref\":"
		out.RawString(prefix)
		out.String(string(in.Idref))
	}
	if in.Text != "" {
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RelatedNode) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RelatedNode) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RelatedNode) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RelatedNode) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility2(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility3(in *jlexer.Lexer, out *Property) {
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
		case "name":
			(out.Name).UnmarshalEasyJSON(in)
		case "value":
			if in.IsNull() {
				in.Skip()
				out.Value = nil
			} else {
				if out.Value == nil {
					out.Value = new(Value)
				}
				(*out.Value).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility3(out *jwriter.Writer, in Property) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		(in.Name).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		if in.Value == nil {
			out.RawString("null")
		} else {
			(*in.Value).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Property) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Property) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Property) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Property) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility3(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility4(in *jlexer.Lexer, out *Node) {
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
		case "nodeId":
			out.NodeID = NodeID(in.String())
		case "ignored":
			out.Ignored = bool(in.Bool())
		case "ignoredReasons":
			if in.IsNull() {
				in.Skip()
				out.IgnoredReasons = nil
			} else {
				in.Delim('[')
				if out.IgnoredReasons == nil {
					if !in.IsDelim(']') {
						out.IgnoredReasons = make([]*Property, 0, 8)
					} else {
						out.IgnoredReasons = []*Property{}
					}
				} else {
					out.IgnoredReasons = (out.IgnoredReasons)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *Property
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(Property)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					out.IgnoredReasons = append(out.IgnoredReasons, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "role":
			if in.IsNull() {
				in.Skip()
				out.Role = nil
			} else {
				if out.Role == nil {
					out.Role = new(Value)
				}
				(*out.Role).UnmarshalEasyJSON(in)
			}
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(Value)
				}
				(*out.Name).UnmarshalEasyJSON(in)
			}
		case "description":
			if in.IsNull() {
				in.Skip()
				out.Description = nil
			} else {
				if out.Description == nil {
					out.Description = new(Value)
				}
				(*out.Description).UnmarshalEasyJSON(in)
			}
		case "value":
			if in.IsNull() {
				in.Skip()
				out.Value = nil
			} else {
				if out.Value == nil {
					out.Value = new(Value)
				}
				(*out.Value).UnmarshalEasyJSON(in)
			}
		case "properties":
			if in.IsNull() {
				in.Skip()
				out.Properties = nil
			} else {
				in.Delim('[')
				if out.Properties == nil {
					if !in.IsDelim(']') {
						out.Properties = make([]*Property, 0, 8)
					} else {
						out.Properties = []*Property{}
					}
				} else {
					out.Properties = (out.Properties)[:0]
				}
				for !in.IsDelim(']') {
					var v8 *Property
					if in.IsNull() {
						in.Skip()
						v8 = nil
					} else {
						if v8 == nil {
							v8 = new(Property)
						}
						(*v8).UnmarshalEasyJSON(in)
					}
					out.Properties = append(out.Properties, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "childIds":
			if in.IsNull() {
				in.Skip()
				out.ChildIds = nil
			} else {
				in.Delim('[')
				if out.ChildIds == nil {
					if !in.IsDelim(']') {
						out.ChildIds = make([]NodeID, 0, 4)
					} else {
						out.ChildIds = []NodeID{}
					}
				} else {
					out.ChildIds = (out.ChildIds)[:0]
				}
				for !in.IsDelim(']') {
					var v9 NodeID
					v9 = NodeID(in.String())
					out.ChildIds = append(out.ChildIds, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "backendDOMNodeId":
			(out.BackendDOMNodeID).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility4(out *jwriter.Writer, in Node) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nodeId\":"
		out.RawString(prefix[1:])
		out.String(string(in.NodeID))
	}
	{
		const prefix string = ",\"ignored\":"
		out.RawString(prefix)
		out.Bool(bool(in.Ignored))
	}
	if len(in.IgnoredReasons) != 0 {
		const prefix string = ",\"ignoredReasons\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v10, v11 := range in.IgnoredReasons {
				if v10 > 0 {
					out.RawByte(',')
				}
				if v11 == nil {
					out.RawString("null")
				} else {
					(*v11).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Role != nil {
		const prefix string = ",\"role\":"
		out.RawString(prefix)
		(*in.Role).MarshalEasyJSON(out)
	}
	if in.Name != nil {
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		(*in.Name).MarshalEasyJSON(out)
	}
	if in.Description != nil {
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		(*in.Description).MarshalEasyJSON(out)
	}
	if in.Value != nil {
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		(*in.Value).MarshalEasyJSON(out)
	}
	if len(in.Properties) != 0 {
		const prefix string = ",\"properties\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v12, v13 := range in.Properties {
				if v12 > 0 {
					out.RawByte(',')
				}
				if v13 == nil {
					out.RawString("null")
				} else {
					(*v13).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.ChildIds) != 0 {
		const prefix string = ",\"childIds\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v14, v15 := range in.ChildIds {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.String(string(v15))
			}
			out.RawByte(']')
		}
	}
	if in.BackendDOMNodeID != 0 {
		const prefix string = ",\"backendDOMNodeId\":"
		out.RawString(prefix)
		out.Int64(int64(in.BackendDOMNodeID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Node) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Node) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Node) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Node) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility4(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility5(in *jlexer.Lexer, out *GetPartialAXTreeReturns) {
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
		case "nodes":
			if in.IsNull() {
				in.Skip()
				out.Nodes = nil
			} else {
				in.Delim('[')
				if out.Nodes == nil {
					if !in.IsDelim(']') {
						out.Nodes = make([]*Node, 0, 8)
					} else {
						out.Nodes = []*Node{}
					}
				} else {
					out.Nodes = (out.Nodes)[:0]
				}
				for !in.IsDelim(']') {
					var v16 *Node
					if in.IsNull() {
						in.Skip()
						v16 = nil
					} else {
						if v16 == nil {
							v16 = new(Node)
						}
						(*v16).UnmarshalEasyJSON(in)
					}
					out.Nodes = append(out.Nodes, v16)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility5(out *jwriter.Writer, in GetPartialAXTreeReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Nodes) != 0 {
		const prefix string = ",\"nodes\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v17, v18 := range in.Nodes {
				if v17 > 0 {
					out.RawByte(',')
				}
				if v18 == nil {
					out.RawString("null")
				} else {
					(*v18).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetPartialAXTreeReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetPartialAXTreeReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetPartialAXTreeReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetPartialAXTreeReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility5(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility6(in *jlexer.Lexer, out *GetPartialAXTreeParams) {
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
		case "nodeId":
			(out.NodeID).UnmarshalEasyJSON(in)
		case "backendNodeId":
			(out.BackendNodeID).UnmarshalEasyJSON(in)
		case "objectId":
			out.ObjectID = runtime.RemoteObjectID(in.String())
		case "fetchRelatives":
			out.FetchRelatives = bool(in.Bool())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility6(out *jwriter.Writer, in GetPartialAXTreeParams) {
	out.RawByte('{')
	first := true
	_ = first
	if in.NodeID != 0 {
		const prefix string = ",\"nodeId\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.NodeID))
	}
	if in.BackendNodeID != 0 {
		const prefix string = ",\"backendNodeId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.BackendNodeID))
	}
	if in.ObjectID != "" {
		const prefix string = ",\"objectId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ObjectID))
	}
	if in.FetchRelatives {
		const prefix string = ",\"fetchRelatives\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.FetchRelatives))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetPartialAXTreeParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetPartialAXTreeParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetPartialAXTreeParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetPartialAXTreeParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility6(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility7(in *jlexer.Lexer, out *GetFullAXTreeReturns) {
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
		case "nodes":
			if in.IsNull() {
				in.Skip()
				out.Nodes = nil
			} else {
				in.Delim('[')
				if out.Nodes == nil {
					if !in.IsDelim(']') {
						out.Nodes = make([]*Node, 0, 8)
					} else {
						out.Nodes = []*Node{}
					}
				} else {
					out.Nodes = (out.Nodes)[:0]
				}
				for !in.IsDelim(']') {
					var v19 *Node
					if in.IsNull() {
						in.Skip()
						v19 = nil
					} else {
						if v19 == nil {
							v19 = new(Node)
						}
						(*v19).UnmarshalEasyJSON(in)
					}
					out.Nodes = append(out.Nodes, v19)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility7(out *jwriter.Writer, in GetFullAXTreeReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Nodes) != 0 {
		const prefix string = ",\"nodes\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v20, v21 := range in.Nodes {
				if v20 > 0 {
					out.RawByte(',')
				}
				if v21 == nil {
					out.RawString("null")
				} else {
					(*v21).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFullAXTreeReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFullAXTreeReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFullAXTreeReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFullAXTreeReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility7(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility8(in *jlexer.Lexer, out *GetFullAXTreeParams) {
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility8(out *jwriter.Writer, in GetFullAXTreeParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFullAXTreeParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFullAXTreeParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFullAXTreeParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFullAXTreeParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility8(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility9(in *jlexer.Lexer, out *EnableParams) {
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility9(out *jwriter.Writer, in EnableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EnableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EnableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EnableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EnableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility9(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility10(in *jlexer.Lexer, out *DisableParams) {
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility10(out *jwriter.Writer, in DisableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DisableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DisableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAccessibility10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DisableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DisableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAccessibility10(l, v)
}
