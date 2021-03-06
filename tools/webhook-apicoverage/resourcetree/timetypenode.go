/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resourcetree

import "reflect"

// TimeTypeNode is a node type that encapsulates fields that are internally time based. E.g metav1.ObjectMeta.CreationTimestamp or metav1.ObjectMeta.DeletionTimestamp.
// These are internally of type metav1.Time which use standard time type, but their values are specified as timestamp strings with parsing logic to create time objects. For
// use-case we only care if the value is set, so we create TimeTypeNodes and mark them as leafnodes.
type TimeTypeNode struct {
	nodeData
}

func (ti *TimeTypeNode) getData() nodeData {
	return ti.nodeData
}

func (ti *TimeTypeNode) initialize(field string, parent NodeInterface, t reflect.Type, rt *ResourceTree) {
	ti.nodeData.initialize(field, parent, t, rt)
	ti.leafNode = true
}

func (ti *TimeTypeNode) buildChildNodes(t reflect.Type) {}