// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package utils

import (
	json "encoding/json"
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

func easyjsonAe596ad1DecodeGithubComGolineUtils(in *jlexer.Lexer, out *FactoryRescuer) {
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
func easyjsonAe596ad1EncodeGithubComGolineUtils(out *jwriter.Writer, in FactoryRescuer) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FactoryRescuer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAe596ad1EncodeGithubComGolineUtils(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FactoryRescuer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAe596ad1EncodeGithubComGolineUtils(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FactoryRescuer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAe596ad1DecodeGithubComGolineUtils(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FactoryRescuer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAe596ad1DecodeGithubComGolineUtils(l, v)
}
func easyjsonAe596ad1DecodeGithubComGolineUtils1(in *jlexer.Lexer, out *ErrorResponse) {
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
		case "code":
			out.Code = string(in.String())
		case "message":
			out.Message = string(in.String())
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
func easyjsonAe596ad1EncodeGithubComGolineUtils1(out *jwriter.Writer, in ErrorResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"code\":")
	out.String(string(in.Code))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"message\":")
	out.String(string(in.Message))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ErrorResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAe596ad1EncodeGithubComGolineUtils1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ErrorResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAe596ad1EncodeGithubComGolineUtils1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ErrorResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAe596ad1DecodeGithubComGolineUtils1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ErrorResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAe596ad1DecodeGithubComGolineUtils1(l, v)
}