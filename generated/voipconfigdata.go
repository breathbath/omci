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

import (
	"../../omci"
)

type VoipConfigData struct {
	omci.BaseManagedEntity
}

func NewVoipConfigData(params ...ParamData) (omci.IManagedEntity, error) {
	eid := decodeEntityID(params...)
	entity := omci.BaseManagedEntity{
		Name:     "VoipConfigData",
		ClassID:  138,
		EntityID: eid,
		MessageTypes: []omci.MsgType{
			omci.Set,
			omci.Get,
		},
		AttributeMask: 0,
		Attributes: []omci.IAttribute{
			omci.NewUint16Field("ManagedEntityId", 0, omci.Read),
			omci.NewByteField("AvailableSignallingProtocols", 0, omci.Read),
			omci.NewByteField("SignallingProtocolUsed", 0, omci.Read|omci.Write),
			omci.NewUint32Field("AvailableVoipConfigurationMethods", 0, omci.Read),
			omci.NewByteField("VoipConfigurationMethodUsed", 0, omci.Read|omci.Write),
			omci.NewUint16Field("VoipConfigurationAddressPointer", 0, omci.Read|omci.Write),
			omci.NewByteField("VoipConfigurationState", 0, omci.Read),
			omci.NewByteField("RetrieveProfile", 0, omci.Write),
			omci.NewUnknownField("ProfileVersion", 0, omci.Read),
		},
	}
	entity.ComputeAttributeMask()
	return &VoipConfigData{entity}, nil
}
