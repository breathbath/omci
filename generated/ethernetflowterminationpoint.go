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
 *              https://github.com/cboling/OMCI-parser
 */
package generated

// EthernetFlowTerminationPoint (class ID 286 defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type EthernetFlowTerminationPoint struct {
	BaseManagedEntityDefinition
}

// NewEthernetFlowTerminationPoint (class ID 286 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewEthernetFlowTerminationPoint(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "EthernetFlowTerminationPoint",
		ClassID:  286,
		EntityID: eid,
		MessageTypes: []MsgType{
			Set,
			Get,
			Create,
			Delete,
		},
		AttributeMask: 0,
		Attributes: []*AttributeDefinition{
			Uint16Field("ManagedEntityId", 0, Read|SetByCreate),
			UnknownField("DestinationMac", 0, Read|Write|SetByCreate),
			UnknownField("SourceMac", 0, Read),
			ByteField("TagPolicy", 0, Read|Write|SetByCreate),
			Uint16Field("Tci", 0, Read|Write),
			ByteField("Loopback", 0, Read|Write),
		},
	}
	entity.computeAttributeMask()
	return &EthernetFlowTerminationPoint{entity}, nil
}
