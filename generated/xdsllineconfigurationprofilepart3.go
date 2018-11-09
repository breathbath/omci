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

// XdslLineConfigurationProfilePart3 (class ID 106 defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type XdslLineConfigurationProfilePart3 struct {
	BaseManagedEntityDefinition
}

// NewXdslLineConfigurationProfilePart3 (class ID 106 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewXdslLineConfigurationProfilePart3(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "XdslLineConfigurationProfilePart3",
		ClassID:  106,
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
			ByteField("LoopDiagnosticsModeForcedLdsf", 0, Read|Write|SetByCreate),
			ByteField("AutomodeColdStartForced", 0, Read|Write|SetByCreate),
			ByteField("L2Atpr", 0, Read|Write|SetByCreate),
			ByteField("L2Atprt", 0, Read|Write|SetByCreate),
			ByteField("ForceInpDownstream", 0, Read|Write),
			ByteField("ForceInpUpstream", 0, Read|Write),
			ByteField("UpdateRequestFlagForNearEndTestParameters", 0, Read|Write),
			ByteField("UpdateRequestFlagForFarEndTestParameters", 0, Read|Write),
			Uint16Field("InmInterArrivalTimeOffsetUpstream", 0, Read|Write),
			ByteField("InmInterArrivalTimeStepUpstream", 0, Read|Write),
			ByteField("InmClusterContinuationValueUpstream", 0, Read|Write),
			ByteField("InmEquivalentInpModeUpstream", 0, Read|Write),
			Uint16Field("InmInterArrivalTimeOffsetDownstream", 0, Read|Write),
			ByteField("InmInterArrivalTimeStepDownstream", 0, Read|Write),
			ByteField("InmClusterContinuationValueDownstream", 0, Read|Write),
			ByteField("InmEquivalentInpModeDownstream", 0, Read|Write),
		},
	}
	entity.computeAttributeMask()
	return &XdslLineConfigurationProfilePart3{entity}, nil
}
