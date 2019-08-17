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

import "github.com/deckarep/golang-set"

const RtpPseudowireParametersClassId ClassID = ClassID(283)

var rtppseudowireparametersBME *ManagedEntityDefinition

// RtpPseudowireParameters (class ID #283) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type RtpPseudowireParameters struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	rtppseudowireparametersBME = &ManagedEntityDefinition{
		Name:    "RtpPseudowireParameters",
		ClassID: 283,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFC00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: Uint16Field("ClockReference", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 1),
			2: ByteField("RtpTimestampMode", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3: Uint16Field("Ptype", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 3),
			4: Uint64Field("Ssrc", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 4),
			5: Uint16Field("ExpectedPtype", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, false, 5),
			6: Uint64Field("ExpectedSsrc", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, false, 6),
		},
	}
}

// NewRtpPseudowireParameters (class ID 283 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewRtpPseudowireParameters(params ...ParamData) (*ManagedEntity, error) {
	return NewManagedEntity(rtppseudowireparametersBME, params...)
}
