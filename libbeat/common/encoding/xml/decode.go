// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package xml

import (
	"bytes"
	"encoding/xml"
	"io"
	"strings"
	"github.com/sheng855174/elastic/libbeat/logp"
)

// Debug logging functions for this package.
var (
	debugf = logp.MakeDebug("winevent")
)

// A Decoder reads and decodes XML from an input stream.
type Decoder struct {
	prependHyphenToAttr bool
	lowercaseKeys       bool
	xmlDec              *xml.Decoder
}

// NewDecoder returns a new decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
	dec := xml.NewDecoder(r)
	// Assume all text has already been converted to UTF-8, so ignore encoding declarations.
	dec.CharsetReader = func(_ string, input io.Reader) (io.Reader, error) { return input, nil }
	return &Decoder{xmlDec: dec}
}

// PrependHyphenToAttr causes the Decoder to prepend a hyphen ('-') to to all
// XML attribute names.
func (d *Decoder) PrependHyphenToAttr() { d.prependHyphenToAttr = true }

// LowercaseKeys causes the Decoder to transform all key name to lowercase.
func (d *Decoder) LowercaseKeys() { d.lowercaseKeys = true }

// Decode reads XML from the input stream and return a map containing the data.
func (d *Decoder) Decode() (map[string]interface{}, error) {
	_, m, err := d.decode(nil)
	return m, err
}

func (d *Decoder) decode(attrs []xml.Attr) (string, map[string]interface{}, error) {
	elements := map[string]interface{}{}
	var cdata string

	logp.Info("sheng decode(attrs []xml.Attr) (string, map[string]interface{}, error) ...")

	for {
		logp.Info("sheng d.xmlDec.Token() ...")
		t, err := d.xmlDec.Token()
		if err != nil {
			if err == io.EOF { //nolint:errorlint // io.EOF should never be wrapped and is not by xml.Decoder.
				return "", elements, nil
			}
			return "", nil, err
		}

		switch elem := t.(type) {
		case xml.StartElement:
			cdata, subElements, err := d.decode(elem.Attr)
			if err != nil {
				return "", nil, err
			}

			// Combine sub-elements and cdata.
			var add interface{} = subElements
			if len(subElements) == 0 {
				add = cdata
			} else if len(cdata) > 0 {
				subElements["#text"] = cdata
			}

			// Add the data to the current object while taking into account
			// if the current key already exists (in the case of lists).
			key := d.key(elem.Name.Local)
			value := elements[key]
			switch v := value.(type) {
			case nil:
				elements[key] = add
			case []interface{}:
				elements[key] = append(v, add)
			default:
				elements[key] = []interface{}{v, add}
			}
		case xml.CharData:
			cdata = string(bytes.TrimSpace(elem.Copy()))
		case xml.EndElement:
			d.addAttributes(attrs, elements)
			return cdata, elements, nil
		}
	}
}

func (d *Decoder) addAttributes(attrs []xml.Attr, m map[string]interface{}) {
	for _, attr := range attrs {
		key := d.attrKey(attr.Name.Local)
		m[key] = attr.Value
	}
}

func (d *Decoder) key(in string) string {
	if d.lowercaseKeys {
		return strings.ToLower(in)
	}
	return in
}

func (d *Decoder) attrKey(in string) string {
	if d.prependHyphenToAttr {
		return d.key("-" + in)
	}
	return d.key(in)
}
