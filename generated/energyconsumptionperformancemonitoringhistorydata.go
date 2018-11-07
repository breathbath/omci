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

type EnergyConsumptionPerformanceMonitoringHistoryData struct {
	omci.BaseManagedEntity
}

func NewEnergyConsumptionPerformanceMonitoringHistoryData(params ...ParamData) (omci.IManagedEntity, error) {
	eid := decodeEntityID(params...)
	entity := omci.BaseManagedEntity{
		Name:     "EnergyConsumptionPerformanceMonitoringHistoryData",
		ClassID:  343,
		EntityID: eid,
		MessageTypes: []omci.MsgType{
			omci.Set,
			omci.Get,
			omci.Create,
			omci.Delete,
		},
		AttributeMask: 0,
		Attributes: []omci.IAttribute{
			omci.NewUint16Field("ManagedEntityId", 0, omci.Read|omci.SetByCreate),
			omci.NewByteField("IntervalEndTime", 0, omci.Read),
			omci.NewUint16Field("ThresholdData12Id", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUint32Field("DozeTime", 0, omci.Read),
			omci.NewUint32Field("CyclicSleepTime", 0, omci.Read),
			omci.NewUint32Field("WatchfulSleepTime", 0, omci.Read),
			omci.NewUint32Field("EnergyConsumed", 0, omci.Read),
		},
	}
	entity.ComputeAttributeMask()
	return &EnergyConsumptionPerformanceMonitoringHistoryData{entity}, nil
}
