/*
Copyright 2022-2023 The Nephio Authors.

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

package nfdeploy

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// UpfDeploy is the Schema for the upfdeploys API
type UpfDeploy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UpfDeploySpec   `json:"spec,omitempty"`
	Status UpfDeployStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// UpfDeployList contains a list of UpfDeploy
type UpfDeployList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UpfDeploy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UpfDeploy{}, &UpfDeployList{})
}
