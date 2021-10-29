/*
Copyright 2021.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MyresouceSpec defines the desired state of Myresouce
type MyresouceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Myresouce. Edit myresouce_types.go to remove/update

	//+kubebuilder:validation:Required
	SecondNum *int `json:"secondNum"`

	//+kubebuilder:validation:Required
	FirstNum *int `json:"firstNum"`

	//+kubebuilder:validation:Required
	Operation string `json:"operation"`
}

// MyresouceStatus defines the observed state of Myresouce
type MyresouceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	State   string `json:"state,omitempty"`
	Message string `json:"message,omitempty"`

	//+kubebuilder:validation:Schemaless
	LastUpdatedc *int `json:"lastUpdated,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Myresouce is the Schema for the myresouces API
type Myresouce struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MyresouceSpec `json:"spec,omitempty"`

	Status MyresouceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MyresouceList contains a list of Myresouce
type MyresouceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Myresouce `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Myresouce{}, &MyresouceList{})
}
