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

const FastLineInventoryAndStatusDataClassId uint16 = 435

// FastLineInventoryAndStatusData (class ID #435) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type FastLineInventoryAndStatusData struct {
	BaseManagedEntityDefinition
}

// NewFastLineInventoryAndStatusData (class ID 435 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewFastLineInventoryAndStatusData(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "FastLineInventoryAndStatusData",
		ClassID:  435,
		EntityID: eid,
		MessageTypes: []MsgType{
			Get,
		},
		AllowedAttributeMask: 0,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, Read),
			1: ByteField("ItuTG9701ProfileProfile", 0, Read),
			2: MultiByteField("GammaDataRAteGdr", 0.0, nil, Read),
			3: MultiByteField("AttainableGammaDataRaTeAttgdr", 0.0, nil, Read),
			4: Uint64Field("DpuSystemVendorIdDpuSystemVendor", 0, Read),
			5: Uint64Field("NtSystemVendorIdNtSystemVendor", 0, Read),
			6: MultiByteField("DpuSerialNumberDpuSystemSerialnr", 32, nil, Read),
			7: MultiByteField("NtSerialNumberNtSystemSerialnr", 32, nil, Read),
		},
	}
	entity.computeAttributeMask()
	return &FastLineInventoryAndStatusData{entity}, nil
}
