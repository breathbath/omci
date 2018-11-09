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

// MacBridgePortConfigurationData (class ID 47 defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type MacBridgePortConfigurationData struct {
	BaseManagedEntityDefinition
}

// NewMacBridgePortConfigurationData (class ID 47 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewMacBridgePortConfigurationData(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "MacBridgePortConfigurationData",
		ClassID:  47,
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
			Uint16Field("BridgeIdPointer", 0, Read|Write|SetByCreate),
			ByteField("PortNum", 0, Read|Write|SetByCreate),
			ByteField("TpType", 0, Read|Write|SetByCreate),
			Uint16Field("TpPointer", 0, Read|Write|SetByCreate),
			Uint16Field("PortPriority", 0, Read|Write|SetByCreate),
			Uint16Field("PortPathCost", 0, Read|Write|SetByCreate),
			ByteField("PortSpanningTreeInd", 0, Read|Write|SetByCreate),
			ByteField("Deprecated1", 0, Read|Write|SetByCreate),
			ByteField("Deprecated2", 0, Read|Write|SetByCreate),
			UnknownField("PortMacAddress", 0, Read),
			Uint16Field("OutboundTdPointer", 0, Read|Write),
			Uint16Field("InboundTdPointer", 0, Read|Write),
			ByteField("MacLearningDepth", 0, Read|Write|SetByCreate),
		},
	}
	entity.computeAttributeMask()
	return &MacBridgePortConfigurationData{entity}, nil
}
