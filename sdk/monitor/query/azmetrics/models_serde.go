//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azmetrics

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type LocalizableString.
func (l LocalizableString) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "localizedValue", l.LocalizedValue)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LocalizableString.
func (l *LocalizableString) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "localizedValue":
			err = unpopulate(val, "LocalizedValue", &l.LocalizedValue)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &l.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetadataValue.
func (m MetadataValue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetadataValue.
func (m *MetadataValue) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &m.Name)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &m.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Metric.
func (m Metric) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "displayDescription", m.DisplayDescription)
	populate(objectMap, "errorCode", m.ErrorCode)
	populate(objectMap, "errorMessage", m.ErrorMessage)
	populate(objectMap, "id", m.ID)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "timeseries", m.TimeSeries)
	populate(objectMap, "type", m.Type)
	populate(objectMap, "unit", m.Unit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Metric.
func (m *Metric) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "displayDescription":
			err = unpopulate(val, "DisplayDescription", &m.DisplayDescription)
			delete(rawMsg, key)
		case "errorCode":
			err = unpopulate(val, "ErrorCode", &m.ErrorCode)
			delete(rawMsg, key)
		case "errorMessage":
			err = unpopulate(val, "ErrorMessage", &m.ErrorMessage)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &m.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &m.Name)
			delete(rawMsg, key)
		case "timeseries":
			err = unpopulate(val, "TimeSeries", &m.TimeSeries)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &m.Type)
			delete(rawMsg, key)
		case "unit":
			err = unpopulate(val, "Unit", &m.Unit)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetricResults.
func (m MetricResults) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "values", m.Values)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetricResults.
func (m *MetricResults) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "values":
			err = unpopulate(val, "Values", &m.Values)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetricValues.
func (m MetricValues) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "endtime", m.EndTime)
	populate(objectMap, "interval", m.Interval)
	populate(objectMap, "namespace", m.Namespace)
	populate(objectMap, "resourceid", m.ResourceID)
	populate(objectMap, "resourceregion", m.ResourceRegion)
	populate(objectMap, "starttime", m.StartTime)
	populate(objectMap, "value", m.Values)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetricValues.
func (m *MetricValues) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endtime":
			err = unpopulate(val, "EndTime", &m.EndTime)
			delete(rawMsg, key)
		case "interval":
			err = unpopulate(val, "Interval", &m.Interval)
			delete(rawMsg, key)
		case "namespace":
			err = unpopulate(val, "Namespace", &m.Namespace)
			delete(rawMsg, key)
		case "resourceid":
			err = unpopulate(val, "ResourceID", &m.ResourceID)
			delete(rawMsg, key)
		case "resourceregion":
			err = unpopulate(val, "ResourceRegion", &m.ResourceRegion)
			delete(rawMsg, key)
		case "starttime":
			err = unpopulate(val, "StartTime", &m.StartTime)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Values", &m.Values)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetricValue.
func (m MetricValue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "average", m.Average)
	populate(objectMap, "count", m.Count)
	populate(objectMap, "maximum", m.Maximum)
	populate(objectMap, "minimum", m.Minimum)
	populateDateTimeRFC3339(objectMap, "timeStamp", m.TimeStamp)
	populate(objectMap, "total", m.Total)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetricValue.
func (m *MetricValue) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "average":
			err = unpopulate(val, "Average", &m.Average)
			delete(rawMsg, key)
		case "count":
			err = unpopulate(val, "Count", &m.Count)
			delete(rawMsg, key)
		case "maximum":
			err = unpopulate(val, "Maximum", &m.Maximum)
			delete(rawMsg, key)
		case "minimum":
			err = unpopulate(val, "Minimum", &m.Minimum)
			delete(rawMsg, key)
		case "timeStamp":
			err = unpopulateDateTimeRFC3339(val, "TimeStamp", &m.TimeStamp)
			delete(rawMsg, key)
		case "total":
			err = unpopulate(val, "Total", &m.Total)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResourceIDList.
func (r ResourceIDList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "resourceids", r.ResourceIDs)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ResourceIDList.
func (r *ResourceIDList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceids":
			err = unpopulate(val, "ResourceIDs", &r.ResourceIDs)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TimeSeriesElement.
func (t TimeSeriesElement) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "data", t.Data)
	populate(objectMap, "metadatavalues", t.MetadataValues)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TimeSeriesElement.
func (t *TimeSeriesElement) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "data":
			err = unpopulate(val, "Data", &t.Data)
			delete(rawMsg, key)
		case "metadatavalues":
			err = unpopulate(val, "MetadataValues", &t.MetadataValues)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}