// Code generated by zanzibar
// @generated
// Checksum : bpbC6E/fg9UWEqM0jkRsZQ==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package googlenow

import (
	json "encoding/json"
	fmt "fmt"
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

func easyjson25b7f9c9DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials(in *jlexer.Lexer, out *GoogleNowService_CheckCredentials_Result) {
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
func easyjson25b7f9c9EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials(out *jwriter.Writer, in GoogleNowService_CheckCredentials_Result) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GoogleNowService_CheckCredentials_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson25b7f9c9EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GoogleNowService_CheckCredentials_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson25b7f9c9EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GoogleNowService_CheckCredentials_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson25b7f9c9DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GoogleNowService_CheckCredentials_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson25b7f9c9DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials(l, v)
}
func easyjson25b7f9c9DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials1(in *jlexer.Lexer, out *GoogleNowService_CheckCredentials_Args) {
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
func easyjson25b7f9c9EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials1(out *jwriter.Writer, in GoogleNowService_CheckCredentials_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GoogleNowService_CheckCredentials_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson25b7f9c9EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GoogleNowService_CheckCredentials_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson25b7f9c9EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GoogleNowService_CheckCredentials_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson25b7f9c9DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GoogleNowService_CheckCredentials_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson25b7f9c9DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials1(l, v)
}
func easyjson25b7f9c9DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceAddCredentials(in *jlexer.Lexer, out *GoogleNowService_AddCredentials_Result) {
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
func easyjson25b7f9c9EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceAddCredentials(out *jwriter.Writer, in GoogleNowService_AddCredentials_Result) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GoogleNowService_AddCredentials_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson25b7f9c9EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceAddCredentials(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GoogleNowService_AddCredentials_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson25b7f9c9EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceAddCredentials(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GoogleNowService_AddCredentials_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson25b7f9c9DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceAddCredentials(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GoogleNowService_AddCredentials_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson25b7f9c9DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceAddCredentials(l, v)
}
func easyjson25b7f9c9DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceAddCredentials1(in *jlexer.Lexer, out *GoogleNowService_AddCredentials_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var AuthCodeSet bool
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
		case "authCode":
			out.AuthCode = string(in.String())
			AuthCodeSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !AuthCodeSet {
		in.AddError(fmt.Errorf("key 'authCode' is required"))
	}
}
func easyjson25b7f9c9EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceAddCredentials1(out *jwriter.Writer, in GoogleNowService_AddCredentials_Args) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"authCode\":"
		out.RawString(prefix[1:])
		out.String(string(in.AuthCode))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GoogleNowService_AddCredentials_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson25b7f9c9EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceAddCredentials1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GoogleNowService_AddCredentials_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson25b7f9c9EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceAddCredentials1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GoogleNowService_AddCredentials_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson25b7f9c9DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceAddCredentials1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GoogleNowService_AddCredentials_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson25b7f9c9DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceAddCredentials1(l, v)
}
