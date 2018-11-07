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

type PhysicalPathTerminationPointEthernetUni struct {
	omci.BaseManagedEntity
}

func NewPhysicalPathTerminationPointEthernetUni(params ...ParamData) (omci.IManagedEntity, error) {
	eid := decodeEntityID(params...)
	entity := omci.BaseManagedEntity{
		Name:     "PhysicalPathTerminationPointEthernetUni",
		ClassID:  11,
		EntityID: eid,
		MessageTypes: []omci.MsgType{
			omci.Set,
			omci.Get,
		},
		AttributeMask: 0,
		Attributes: []omci.IAttribute{
			omci.NewUint16Field("ManagedEntityId:", 0, omci.Read),
			omci.NewByteField("ExpectedType", 0, omci.Read|omci.Write),
			omci.NewByteField("SensedType", 0, omci.Read),
			omci.NewByteField("AutoDetectionConfiguration", 0, omci.Read|omci.Write),
			omci.NewByteField("EthernetLoopbackConfiguration", 0, omci.Read|omci.Write),
			omci.NewByteField("AdministrativeState", 0, omci.Read|omci.Write),
			omci.NewByteField("OperationalState", 0, omci.Read),
			omci.NewByteField("ConfigurationInd", 0, omci.Read),
			omci.NewUint16Field("MaxFrameSize", 0, omci.Read|omci.Write),
			omci.NewByteField("DteOrDceInd", 0, omci.Read|omci.Write),
			omci.NewUint16Field("PauseTime", 0, omci.Read|omci.Write),
			omci.NewByteField("BridgedOrIpInd", 0, omci.Read|omci.Write),
			omci.NewByteField("Arc", 0, omci.Read|omci.Write),
			omci.NewByteField("ArcInterval", 0, omci.Read|omci.Write),
			omci.NewByteField("PppoeFilter", 0, omci.Read|omci.Write),
			omci.NewByteField("PowerControl", 0, omci.Read|omci.Write),
		},
	}
	entity.ComputeAttributeMask()
	return &PhysicalPathTerminationPointEthernetUni{entity}, nil
}
