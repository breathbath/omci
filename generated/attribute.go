/*
 * Copyright (c) 2018 - present.  Boling Consulting Solutions (bcsw.net)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/*
 * NOTE: This file was generated, manual edits will be overwritten!
 *
 * Generated by 'goCodeGenerator.py':
 *              https://github.com/cboling/OMCI-parser/README.md
 */
package generated

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/google/gopacket"
	"sort"
	"strings"
)

type AttributeDefinitionMap map[uint]*AttributeDefinition

// AttributeDefinition defines a single specific Managed Entity attribute
type AttributeDefinition struct {
	Name       string
	DefValue   interface{}
	Size       int
	Access     AttributeAccess
	Constraint func(interface{}) error
	Avc        bool // If true, an AVC notification can occur for the attribute
	Tca        bool // If true, a threshold crossing alert alarm notification can occur for the attribute
	Counter    bool // If true, this attribute is a PM counter
	Optional   bool // If true, attribute is option, else mandatory
	Deprecated bool //  If true, this attribute is deprecated and only 'read' operations (if-any) performed
}

func (attr *AttributeDefinition) String() string {
	return fmt.Sprintf("Definition: %v: Size: %v, Default: %v, Access: %v",
		attr.GetName(), attr.GetSize(), attr.GetDefault(), attr.GetAccess())
}
func (attr *AttributeDefinition) GetName() string             { return attr.Name }
func (attr *AttributeDefinition) GetDefault() interface{}     { return attr.DefValue }
func (attr *AttributeDefinition) GetSize() int                { return attr.Size }
func (attr *AttributeDefinition) GetAccess() AttributeAccess  { return attr.Access }
func (attr *AttributeDefinition) GetConstraints() func(interface{}) error {
	return attr.Constraint
}

func (attr *AttributeDefinition) Decode(data []byte, df gopacket.DecodeFeedback) (interface{}, error) {
	// Use negative numbers to indicate signed values
	size := attr.GetSize()
	if size < 0 {
		size = -size
	}
	if len(data) < size {
		df.SetTruncated()
		return nil, errors.New("packet too small for field")
	}
	var err error
	switch attr.GetSize() {
	default:
		value := make([]byte, size)
		if attr.GetConstraints() != nil {
			err = attr.GetConstraints()(value)
			if err != nil {
				return nil, err
			}
		}
		return value, err
	case 1:
		value := data[0]
		if attr.GetConstraints() != nil {
			err = attr.GetConstraints()(value)
			if err != nil {
				return nil, err
			}
		}
		return value, err
	case 2:
		value := binary.BigEndian.Uint16(data[0:2])
		if attr.GetConstraints() != nil {
			err = attr.GetConstraints()(value)
			if err != nil {
				return nil, err
			}
		}
		return value, err
	case 4:
		value := binary.BigEndian.Uint32(data[0:4])
		if attr.GetConstraints() != nil {
			err = attr.GetConstraints()(value)
			if err != nil {
				return nil, err
			}
		}
		return value, err
	case 8:
		value := binary.BigEndian.Uint64(data[0:8])
		if attr.GetConstraints() != nil {
			err = attr.GetConstraints()(value)
			if err != nil {
				return nil, err
			}
		}
		return value, err
	}
}

func (attr *AttributeDefinition) SerializeTo(value interface{}, b gopacket.SerializeBuffer) error {
	// TODO: Check to see if space in buffer here !!!!
	bytes, err := b.AppendBytes(attr.GetSize())
	if err != nil {
		return err
	}
	switch attr.GetSize() {
	default:
		copy(bytes, value.([]byte))
	case 1:
		bytes[0] = value.(byte)
	case 2:
		binary.BigEndian.PutUint16(bytes, value.(uint16))
	case 4:
		binary.BigEndian.PutUint32(bytes, value.(uint32))
	case 8:
		binary.BigEndian.PutUint64(bytes, value.(uint64))
	}
	return nil
}

// GetAttributeDefinitionByName searches the attribute definition map for the
// attribute with the specified name (case insensitive)
func GetAttributeDefinitionByName(attrMap AttributeDefinitionMap, name string) (*AttributeDefinition, error) {
	nameLower := strings.ToLower(name)
	for _, attrVal := range attrMap {
		if nameLower == strings.ToLower(attrVal.GetName()) {
			return attrVal, nil
		}
	}
	return nil, errors.New("attribute not found")
}

// GetAttributeDefinitionMapKeys is a convenience functions since we may need to
// iterate a map in key index order. Maps in Go since v1.0 the iteration order
// of maps have been randomized.
func GetAttributeDefinitionMapKeys(attrMap AttributeDefinitionMap) []uint {
	var keys []uint
	for k:= range attrMap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return i < j })
	return keys
}

///////////////////////////////////////////////////////////////////////
// Packet definitions for attributes of various types/sizes

func ByteField(name string, defVal uint16, access AttributeAccess) *AttributeDefinition {
	return &AttributeDefinition{Name: name, DefValue: defVal, Size: 1, Access: access}
}

func Uint16Field(name string, defVal uint16, access AttributeAccess) *AttributeDefinition {
	return &AttributeDefinition{Name: name, DefValue: defVal, Size: 2, Access: access}
}

func Uint32Field(name string, defVal uint16, access AttributeAccess) *AttributeDefinition {
	return &AttributeDefinition{Name: name, DefValue: defVal, Size: 4, Access: access}
}

func Uint64Field(name string, defVal uint16, access AttributeAccess) *AttributeDefinition {
	return &AttributeDefinition{Name: name, DefValue: defVal, Size: 8, Access: access}
}

func MultiByteField(name string, size uint, defVal []byte, access AttributeAccess) *AttributeDefinition {
	return &AttributeDefinition{Name: name, DefValue: defVal, Size: int(size), Access: access}
}

// TODO: Need more fields...
func UnknownField(name string, defVal uint16, access AttributeAccess) *AttributeDefinition {
	return &AttributeDefinition{Name: name, DefValue: defVal, Size: 99999999, Access: access}
}

///////////////////////////////////////////////////////////////////////
// Attribute Value   (Interfaced defined in generated subdirectory)

type AttributeValueMap map[uint]*AttributeValue

// AttributeValue provides the value for a single specific Managed Entity attribute
type AttributeValue struct {
	Name   string
	Value  interface{}
}

func (attr *AttributeValue) String() string {
	val, err := attr.GetValue()
	return fmt.Sprintf("Value: %v, Value: %v, Error: %v",
		attr.GetName(), val, err)
}
func (attr *AttributeValue) GetName() string  { return attr.Name }
func (attr *AttributeValue) GetValue() (interface{}, error) {
	// TODO: Better way to detect not-initialized and no default available?
	return attr.Value, nil
}

func (attr *AttributeValue) SetValue(value interface{}) error {
	return nil
}

// GetAttributeValueByName searches the attribute value map for the
// attribute with the specified name (case insensitive)
func GetAttributeValueByName(attrMap AttributeValueMap, name string) (*AttributeValue, error) {
	nameLower := strings.ToLower(name)
	for _, attrVal := range attrMap {
		if nameLower == strings.ToLower(attrVal.GetName()) {
			return attrVal, nil
		}
	}
	return nil, errors.New("attribute not found")
}

// GetAttributeValueMapKeys is a convenience functions since we may need to
// iterate a map in key index order. Maps in Go since v1.0 the iteration order
// of maps have been randomized.
func GetAttributeValueMapKeys(attrMap AttributeValueMap) []uint {
	var keys []uint
	for k:= range attrMap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return i < j })
	return keys
}