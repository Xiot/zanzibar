// Code generated by zanzibar
// @generated
// Checksum : wCzBREXeazUMAaApk2fuFg==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package contacts

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

func easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts(in *jlexer.Lexer, out *SaveContactsResponse) {
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
func easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts(out *jwriter.Writer, in SaveContactsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SaveContactsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SaveContactsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SaveContactsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SaveContactsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts(l, v)
}
func easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts1(in *jlexer.Lexer, out *SaveContactsRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var UserUUIDSet bool
	var ContactsSet bool
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
		case "userUUID":
			out.UserUUID = string(in.String())
			UserUUIDSet = true
		case "contacts":
			if in.IsNull() {
				in.Skip()
				out.Contacts = nil
			} else {
				in.Delim('[')
				if out.Contacts == nil {
					if !in.IsDelim(']') {
						out.Contacts = make([]*Contact, 0, 8)
					} else {
						out.Contacts = []*Contact{}
					}
				} else {
					out.Contacts = (out.Contacts)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Contact
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Contact)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Contacts = append(out.Contacts, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
			ContactsSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !UserUUIDSet {
		in.AddError(fmt.Errorf("key 'userUUID' is required"))
	}
	if !ContactsSet {
		in.AddError(fmt.Errorf("key 'contacts' is required"))
	}
}
func easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts1(out *jwriter.Writer, in SaveContactsRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"userUUID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UserUUID))
	}
	{
		const prefix string = ",\"contacts\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Contacts == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Contacts {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SaveContactsRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SaveContactsRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SaveContactsRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SaveContactsRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts1(l, v)
}
func easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts2(in *jlexer.Lexer, out *ContactFragment) {
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
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				if out.Type == nil {
					out.Type = new(ContactFragmentType)
				}
				*out.Type = ContactFragmentType(in.String())
			}
		case "text":
			if in.IsNull() {
				in.Skip()
				out.Text = nil
			} else {
				if out.Text == nil {
					out.Text = new(string)
				}
				*out.Text = string(in.String())
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
func easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts2(out *jwriter.Writer, in ContactFragment) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Type != nil {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Type))
	}
	if in.Text != nil {
		const prefix string = ",\"text\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Text))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ContactFragment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ContactFragment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ContactFragment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ContactFragment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts2(l, v)
}
func easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts3(in *jlexer.Lexer, out *ContactAttributes) {
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
		case "firstName":
			if in.IsNull() {
				in.Skip()
				out.FirstName = nil
			} else {
				if out.FirstName == nil {
					out.FirstName = new(string)
				}
				*out.FirstName = string(in.String())
			}
		case "lastName":
			if in.IsNull() {
				in.Skip()
				out.LastName = nil
			} else {
				if out.LastName == nil {
					out.LastName = new(string)
				}
				*out.LastName = string(in.String())
			}
		case "nickname":
			if in.IsNull() {
				in.Skip()
				out.Nickname = nil
			} else {
				if out.Nickname == nil {
					out.Nickname = new(string)
				}
				*out.Nickname = string(in.String())
			}
		case "hasPhoto":
			if in.IsNull() {
				in.Skip()
				out.HasPhoto = nil
			} else {
				if out.HasPhoto == nil {
					out.HasPhoto = new(bool)
				}
				*out.HasPhoto = bool(in.Bool())
			}
		case "numFields":
			if in.IsNull() {
				in.Skip()
				out.NumFields = nil
			} else {
				if out.NumFields == nil {
					out.NumFields = new(int32)
				}
				*out.NumFields = int32(in.Int32())
			}
		case "timesContacted":
			if in.IsNull() {
				in.Skip()
				out.TimesContacted = nil
			} else {
				if out.TimesContacted == nil {
					out.TimesContacted = new(int32)
				}
				*out.TimesContacted = int32(in.Int32())
			}
		case "lastTimeContacted":
			if in.IsNull() {
				in.Skip()
				out.LastTimeContacted = nil
			} else {
				if out.LastTimeContacted == nil {
					out.LastTimeContacted = new(int32)
				}
				*out.LastTimeContacted = int32(in.Int32())
			}
		case "isStarred":
			if in.IsNull() {
				in.Skip()
				out.IsStarred = nil
			} else {
				if out.IsStarred == nil {
					out.IsStarred = new(bool)
				}
				*out.IsStarred = bool(in.Bool())
			}
		case "hasCustomRingtone":
			if in.IsNull() {
				in.Skip()
				out.HasCustomRingtone = nil
			} else {
				if out.HasCustomRingtone == nil {
					out.HasCustomRingtone = new(bool)
				}
				*out.HasCustomRingtone = bool(in.Bool())
			}
		case "isSendToVoicemail":
			if in.IsNull() {
				in.Skip()
				out.IsSendToVoicemail = nil
			} else {
				if out.IsSendToVoicemail == nil {
					out.IsSendToVoicemail = new(bool)
				}
				*out.IsSendToVoicemail = bool(in.Bool())
			}
		case "hasThumbnail":
			if in.IsNull() {
				in.Skip()
				out.HasThumbnail = nil
			} else {
				if out.HasThumbnail == nil {
					out.HasThumbnail = new(bool)
				}
				*out.HasThumbnail = bool(in.Bool())
			}
		case "namePrefix":
			if in.IsNull() {
				in.Skip()
				out.NamePrefix = nil
			} else {
				if out.NamePrefix == nil {
					out.NamePrefix = new(string)
				}
				*out.NamePrefix = string(in.String())
			}
		case "nameSuffix":
			if in.IsNull() {
				in.Skip()
				out.NameSuffix = nil
			} else {
				if out.NameSuffix == nil {
					out.NameSuffix = new(string)
				}
				*out.NameSuffix = string(in.String())
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
func easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts3(out *jwriter.Writer, in ContactAttributes) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FirstName != nil {
		const prefix string = ",\"firstName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.FirstName))
	}
	if in.LastName != nil {
		const prefix string = ",\"lastName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.LastName))
	}
	if in.Nickname != nil {
		const prefix string = ",\"nickname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Nickname))
	}
	if in.HasPhoto != nil {
		const prefix string = ",\"hasPhoto\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(*in.HasPhoto))
	}
	if in.NumFields != nil {
		const prefix string = ",\"numFields\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(*in.NumFields))
	}
	if in.TimesContacted != nil {
		const prefix string = ",\"timesContacted\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(*in.TimesContacted))
	}
	if in.LastTimeContacted != nil {
		const prefix string = ",\"lastTimeContacted\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(*in.LastTimeContacted))
	}
	if in.IsStarred != nil {
		const prefix string = ",\"isStarred\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(*in.IsStarred))
	}
	if in.HasCustomRingtone != nil {
		const prefix string = ",\"hasCustomRingtone\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(*in.HasCustomRingtone))
	}
	if in.IsSendToVoicemail != nil {
		const prefix string = ",\"isSendToVoicemail\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(*in.IsSendToVoicemail))
	}
	if in.HasThumbnail != nil {
		const prefix string = ",\"hasThumbnail\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(*in.HasThumbnail))
	}
	if in.NamePrefix != nil {
		const prefix string = ",\"namePrefix\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.NamePrefix))
	}
	if in.NameSuffix != nil {
		const prefix string = ",\"nameSuffix\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.NameSuffix))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ContactAttributes) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ContactAttributes) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ContactAttributes) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ContactAttributes) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts3(l, v)
}
func easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts4(in *jlexer.Lexer, out *Contact) {
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
		case "fragments":
			if in.IsNull() {
				in.Skip()
				out.Fragments = nil
			} else {
				in.Delim('[')
				if out.Fragments == nil {
					if !in.IsDelim(']') {
						out.Fragments = make([]*ContactFragment, 0, 8)
					} else {
						out.Fragments = []*ContactFragment{}
					}
				} else {
					out.Fragments = (out.Fragments)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *ContactFragment
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(ContactFragment)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Fragments = append(out.Fragments, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "attributes":
			if in.IsNull() {
				in.Skip()
				out.Attributes = nil
			} else {
				if out.Attributes == nil {
					out.Attributes = new(ContactAttributes)
				}
				(*out.Attributes).UnmarshalEasyJSON(in)
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
func easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts4(out *jwriter.Writer, in Contact) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Fragments) != 0 {
		const prefix string = ",\"fragments\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Fragments {
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
	if in.Attributes != nil {
		const prefix string = ",\"attributes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Attributes).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Contact) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Contact) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Contact) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Contact) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContacts4(l, v)
}
