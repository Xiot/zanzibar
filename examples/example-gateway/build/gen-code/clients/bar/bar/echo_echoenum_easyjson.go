// Code generated by zanzibar
// @generated
// Checksum : vBNSXy2D1+M4xkygfOWQtQ==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bar

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

func easyjsonDf3c892cDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoEnum(in *jlexer.Lexer, out *Echo_EchoEnum_Result) {
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
					out.Success = new(Fruit)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Success).UnmarshalJSON(data))
				}
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
func easyjsonDf3c892cEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoEnum(out *jwriter.Writer, in Echo_EchoEnum_Result) {
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
		out.Raw((*in.Success).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Echo_EchoEnum_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDf3c892cEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoEnum(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Echo_EchoEnum_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDf3c892cEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoEnum(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Echo_EchoEnum_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDf3c892cDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoEnum(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Echo_EchoEnum_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDf3c892cDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoEnum(l, v)
}
func easyjsonDf3c892cDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoEnum1(in *jlexer.Lexer, out *Echo_EchoEnum_Args) {
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
		case "arg":
			if in.IsNull() {
				in.Skip()
				out.Arg = nil
			} else {
				if out.Arg == nil {
					out.Arg = new(Fruit)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Arg).UnmarshalJSON(data))
				}
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
func easyjsonDf3c892cEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoEnum1(out *jwriter.Writer, in Echo_EchoEnum_Args) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Arg != nil {
		const prefix string = ",\"arg\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.Arg).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Echo_EchoEnum_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDf3c892cEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoEnum1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Echo_EchoEnum_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDf3c892cEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoEnum1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Echo_EchoEnum_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDf3c892cDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoEnum1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Echo_EchoEnum_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDf3c892cDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoEnum1(l, v)
}
