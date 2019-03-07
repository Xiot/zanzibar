// Code generated by zanzibar
// @generated
// Checksum : 4XV9GOk2i67VeEOO5LSZhA==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package multi

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

func easyjsonEe937abdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsMultiMultiServiceBBackHello(in *jlexer.Lexer, out *ServiceBBack_Hello_Result) {
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
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(string)
				}
				*out.Success = string(in.String())
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
func easyjsonEe937abdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsMultiMultiServiceBBackHello(out *jwriter.Writer, in ServiceBBack_Hello_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Success))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceBBack_Hello_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEe937abdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsMultiMultiServiceBBackHello(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceBBack_Hello_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEe937abdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsMultiMultiServiceBBackHello(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceBBack_Hello_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEe937abdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsMultiMultiServiceBBackHello(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceBBack_Hello_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEe937abdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsMultiMultiServiceBBackHello(l, v)
}
func easyjsonEe937abdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsMultiMultiServiceBBackHello1(in *jlexer.Lexer, out *ServiceBBack_Hello_Args) {
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
func easyjsonEe937abdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsMultiMultiServiceBBackHello1(out *jwriter.Writer, in ServiceBBack_Hello_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceBBack_Hello_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEe937abdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsMultiMultiServiceBBackHello1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceBBack_Hello_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEe937abdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsMultiMultiServiceBBackHello1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceBBack_Hello_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEe937abdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsMultiMultiServiceBBackHello1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceBBack_Hello_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEe937abdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsMultiMultiServiceBBackHello1(l, v)
}
