// Code generated by zanzibar
// @generated
// Checksum : TPvZ1L/B5BCI9weBUDZhaw==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package googlenow

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

func easyjson950a5fb3DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials(in *jlexer.Lexer, out *GoogleNowService_CheckCredentials_Result) {
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
func easyjson950a5fb3EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials(out *jwriter.Writer, in GoogleNowService_CheckCredentials_Result) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GoogleNowService_CheckCredentials_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson950a5fb3EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GoogleNowService_CheckCredentials_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson950a5fb3EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GoogleNowService_CheckCredentials_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson950a5fb3DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GoogleNowService_CheckCredentials_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson950a5fb3DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials(l, v)
}
func easyjson950a5fb3DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials1(in *jlexer.Lexer, out *GoogleNowService_CheckCredentials_Args) {
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
func easyjson950a5fb3EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials1(out *jwriter.Writer, in GoogleNowService_CheckCredentials_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GoogleNowService_CheckCredentials_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson950a5fb3EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GoogleNowService_CheckCredentials_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson950a5fb3EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GoogleNowService_CheckCredentials_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson950a5fb3DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GoogleNowService_CheckCredentials_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson950a5fb3DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsGooglenowGooglenowGoogleNowServiceCheckCredentials1(l, v)
}
