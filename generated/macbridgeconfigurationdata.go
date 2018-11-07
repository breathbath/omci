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

type MacBridgeConfigurationData struct {
	omci.BaseManagedEntity
}

func NewMacBridgeConfigurationData(params ...ParamData) (omci.IManagedEntity, error) {
	eid := decodeEntityID(params...)
	entity := omci.BaseManagedEntity{
		Name:     "MacBridgeConfigurationData",
		ClassID:  46,
		EntityID: eid,
		MessageTypes: []omci.MsgType{
			omci.Get,
		},
		AttributeMask: 0,
		Attributes: []omci.IAttribute{
			omci.NewUint16Field("ManagedEntityId", 0, omci.Read),
			omci.NewUnknownField("BridgeMacAddress", 0, omci.Read),
			omci.NewUint16Field("BridgePriority", 0, omci.Read),
			omci.NewUint64Field("DesignatedRoot", 0, omci.Read),
			omci.NewUint32Field("RootPathCost", 0, omci.Read),
			omci.NewByteField("BridgePortCount", 0, omci.Read),
			omci.NewUint16Field("RootPortNum", 0, omci.Read),
			omci.NewUint16Field("HelloTime", 0, omci.Read),
			omci.NewUint16Field("ForwardDelay", 0, omci.Read),
		},
	}
	entity.ComputeAttributeMask()
	return &MacBridgeConfigurationData{entity}, nil
}
