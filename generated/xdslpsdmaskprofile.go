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

// XdslPsdMaskProfileClassID is the 16-bit ID for the OMCI
// Managed entity xDSL PSD mask profile
const XdslPsdMaskProfileClassID ClassID = ClassID(110)

var xdslpsdmaskprofileBME *ManagedEntityDefinition

// XdslPsdMaskProfile (class ID #110)
//	This ME contains a PSD mask profile for an xDSL UNI. An instance of this ME is created and
//	deleted by the OLT.
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
//		Psd Mask Table
//			(R,-W) (mandatory) (4 * N bytes where N is the number of breakpoints)
//
//		Mask Valid
//			(R,-W) (mandatory) (1-byte)
//
type XdslPsdMaskProfile struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	xdslpsdmaskprofileBME = &ManagedEntityDefinition{
		Name:    "XdslPsdMaskProfile",
		ClassID: 110,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			GetNext,
			Set,
		),
		AllowedAttributeMask: 0xc000,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: TableField("PsdMaskTable", TableInfo{nil, 4}, mapset.NewSetWith(Read, Write), false, false, false, 1),
			2: ByteField("MaskValid", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 2),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewXdslPsdMaskProfile (class ID 110) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewXdslPsdMaskProfile(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*xdslpsdmaskprofileBME, params...)
}
