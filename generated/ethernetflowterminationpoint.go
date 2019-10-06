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

const EthernetFlowTerminationPointClassId ClassID = ClassID(286)

var ethernetflowterminationpointBME *ManagedEntityDefinition

// EthernetFlowTerminationPoint (class ID #286)
//	The Ethernet flow TP contains the attributes necessary to originate and terminate Ethernet
//	frames in the ONU. It is appropriate when transporting pseudowire services via layer-2.
//	Instances of this ME are created and deleted by the OLT.
//
//	Relationships
//		One Ethernet flow TP ME exists for each distinct pseudowire service that is transported via
//		layer 2.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to a pseudowire TP ME. (R, setbycreate) (mandatory)
//			(2-bytes)
//
//		Destination Mac
//			Destination MAC: This attribute specifies the destination MAC address of upstream Ethernet
//			frames. (R,-W, setbycreate) (mandatory) (6-bytes)
//
//		Source Mac
//			Source MAC: This attribute specifies the near-end MAC address. It is established by nonOMCI
//			means (e.g., factory programmed into ONU flash memory) and is included here for information
//			only. (R) (mandatory) (6-bytes)
//
//		Tag Policy
//			(R,-W, setbycreate) (mandatory) (1-byte)
//
//		Tci
//			TCI:	If the tag policy calls for tagging of upstream Ethernet frames, this attribute specifies
//			the tag control information, which includes the VLAN tag, P bits and CFI bit. (R,-W) (optional)
//			(2-bytes)
//
//		Loopback
//			(R,-W) (mandatory) (1-byte)
//
type EthernetFlowTerminationPoint struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	ethernetflowterminationpointBME = &ManagedEntityDefinition{
		Name:    "EthernetFlowTerminationPoint",
		ClassID: 286,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XF800,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: MultiByteField("DestinationMac", 6, nil, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 1),
			2: MultiByteField("SourceMac", 6, nil, mapset.NewSetWith(Read), false, false, false, false, 2),
			3: ByteField("TagPolicy", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 3),
			4: Uint16Field("Tci", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 4),
			5: ByteField("Loopback", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 5),
		},
	}
}

// NewEthernetFlowTerminationPoint (class ID 286 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewEthernetFlowTerminationPoint(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*ethernetflowterminationpointBME, params...)
}
