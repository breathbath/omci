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

// EthernetFrameExtendedPm (class ID 334 defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type EthernetFrameExtendedPm struct {
	BaseManagedEntityDefinition
}

// NewEthernetFrameExtendedPm (class ID 334 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewEthernetFrameExtendedPm(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "EthernetFrameExtendedPm",
		ClassID:  334,
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
			ByteField("IntervalEndTime", 0, Read),
			UnknownField("ControlBlock", 0, Read|Write|SetByCreate),
			Uint32Field("DropEvents", 0, Read),
			Uint32Field("Octets", 0, Read),
			Uint32Field("Frames", 0, Read),
			Uint32Field("BroadcastFrames", 0, Read),
			Uint32Field("MulticastFrames", 0, Read),
			Uint32Field("CrcErroredFrames", 0, Read),
			Uint32Field("UndersizeFrames", 0, Read),
			Uint32Field("OversizeFrames", 0, Read),
			Uint32Field("Frames64Octets", 0, Read),
			Uint32Field("Frames65To127Octets", 0, Read),
			Uint32Field("Frames128To255Octets", 0, Read),
			Uint32Field("Frames256To511Octets", 0, Read),
			Uint32Field("Frames512To1023Octets", 0, Read),
			Uint32Field("Frames1024To1518Octets", 0, Read),
		},
	}
	entity.computeAttributeMask()
	return &EthernetFrameExtendedPm{entity}, nil
}
