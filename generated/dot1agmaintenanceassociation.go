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

const Dot1AgMaintenanceAssociationClassId ClassID = ClassID(300)

var dot1agmaintenanceassociationBME *ManagedEntityDefinition

// Dot1AgMaintenanceAssociation (class ID #300)
//	This ME models an [IEEE 802.1ag] service defined on a bridge port. An MA is a set of endpoints
//	on opposite sides of a network, all existing at a defined maintenance level. One of the
//	endpoints resides on the local ONU; the others are understood to be configured in a consistent
//	way on external equipment. [ITUT Y.1731] refers to the MA as a maintenance entity group (MEG).
//
//	An MA is created and deleted by the OLT.
//
//	Relationships
//		Any number of MAs may be associated with a given MD, or may stand on their own without an MD.
//		One or more MAs may be associated with a MAC bridge or an IEEE 802.1p mapper. An MA exists at
//		one of eight possible maintenance levels.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies an instance of this ME. The values 0 and
//			0xFFFF are reserved. (R, setbycreate) (mandatory) (2-bytes)
//
//		Md Pointer
//			MD pointer:	This pointer specifies the dot1ag maintenance domain with which this MA is
//			associated. A null pointer specifies that the MA is not associated with an MD. (R,-W,
//			setbycreate) (mandatory) (2-bytes)
//
//		Short Ma Name Format
//			Short MA name format: This attribute specifies one of several possible formats for the short MA
//			name attribute. Value 1, the primary VLAN ID, is recommended to be the default. (R,-W,
//			setbycreate) (mandatory) (1-byte)
//
//		Short Ma Name 1, Short Ma Name 2
//			Short MA name 1, Short MA name 2: These two attributes may be regarded as an octet string whose
//			value is the left-justified MA name. Because the MA name may or may not be a printable character
//			string, an octet string is the appropriate representation. If the short MA name format specifies
//			a character string, the string is null-terminated; otherwise, its length is determined by the
//			short MA name format. Note that binary comparisons of the short MA name are made in other CFM
//			state machines, so blanks, alphabetic case, etc., are significant. Also, note that the MD name
//			and the MA short name must be packed (with additional bytes) into 48-byte CFM message headers.
//			(R,-W) (mandatory) (25-bytes * 2 attributes)
//
//		Continuity Check Message Ccm Interval
//			Short intervals should be used judiciously, as they can interfere with the network's ability to
//			handle subscriber traffic. The recommended value is 1-s. (R,-W, setbycreate) (mandatory)
//			(1-byte)
//
//		Associated Vlans
//			Associated VLANs: This attribute is a list of up to 12 VLAN IDs with which this MA is
//			associated. Once a set of VLANs is defined, the ONU should deny operations to other dot1ag MAs
//			or dot1ag default MD level entries that conflict with the set membership. The all-zeros value
//			indicates that this MA is not associated with any VLANs. Assuming that the attribute is not 0,
//			the first entry is understood to be the primary VLAN. Except forwarded linktrace messages
//			(LTMs), CFM messages emitted by MPs in this MA are tagged with the primary VLAN ID. (R,-W)
//			(mandatory) (2-bytes/entry * 12-entries-=-24-bytes)
//
//		Mhf Creation
//			(R,-W, setbycreate) (mandatory) (1-byte)
//
//		Sender Id Permission
//			(R,-W, setbycreate) (mandatory) (1-byte)
//
type Dot1AgMaintenanceAssociation struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	dot1agmaintenanceassociationBME = &ManagedEntityDefinition{
		Name:    "Dot1AgMaintenanceAssociation",
		ClassID: 300,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFE00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: Uint16Field("MdPointer", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 1),
			2: ByteField("ShortMaNameFormat", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3: MultiByteField("ShortMaName1,ShortMaName2", 25, nil, mapset.NewSetWith(Read, Write), false, false, false, false, 3),
			4: ByteField("ContinuityCheckMessageCcmInterval", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 4),
			5: MultiByteField("AssociatedVlans", 24, nil, mapset.NewSetWith(Read, Write), false, false, false, false, 5),
			6: ByteField("MhfCreation", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 6),
			7: ByteField("SenderIdPermission", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 7),
		},
	}
}

// NewDot1AgMaintenanceAssociation (class ID 300 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewDot1AgMaintenanceAssociation(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*dot1agmaintenanceassociationBME, params...)
}
