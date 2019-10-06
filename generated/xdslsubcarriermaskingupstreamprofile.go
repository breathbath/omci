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

const XdslSubcarrierMaskingUpstreamProfileClassId ClassID = ClassID(109)

var xdslsubcarriermaskingupstreamprofileBME *ManagedEntityDefinition

// XdslSubcarrierMaskingUpstreamProfile (class ID #109)
//	This ME contains the subcarrier masking upstream profile for an xDSL UNI. An instance of this ME
//	is created and deleted by the OLT.
//
//	Relationships
//		An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI part
//		1.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. The value 0 is
//			reserved. (R, setbycreate) (mandatory) (2-bytes)
//
//		Upstream Subcarrier Mask
//			Subcarrier number 1 is the lowest, and the number of xDSL subcarriers, upstream (NSCus) is the
//			highest subcarrier that can be transmitted in the upstream direction. For [ITUT-G.992.3],
//			[ITUT-G.992.4] and [ITUT-G.992.5], it is defined in the corresponding Recommendation. For Annex
//			A of [ITUT-G.992.1] and [ITUT G.992.2], NSCus-= 32 and for Annex B of [ITUT-G.992.1], NSCus-=
//			64. (R, W, setbycreate) (mandatory) (8-bytes)
//
type XdslSubcarrierMaskingUpstreamProfile struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	xdslsubcarriermaskingupstreamprofileBME = &ManagedEntityDefinition{
		Name:    "XdslSubcarrierMaskingUpstreamProfile",
		ClassID: 109,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0X8000,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: Uint64Field("UpstreamSubcarrierMask", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 1),
		},
	}
}

// NewXdslSubcarrierMaskingUpstreamProfile (class ID 109 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewXdslSubcarrierMaskingUpstreamProfile(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*xdslsubcarriermaskingupstreamprofileBME, params...)
}
