package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="Desired",type="string",JSONPath=`.spec.replicas`
// +kubebuilder:printcolumn:name="Current",type="string",JSONPath=`.status.replicas`

// PodSet is the Schema for the podsets API
type PodSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PodSetSpec   `json:"spec,omitempty"`
	Status            PodSetStatus `json:"status,omitempty"`
}

// PodSetSpec defines the desired state of PodSet
type PodSetSpec struct {
	Replicas int32 `json:"replicas"`
}

// +k8s:openapi-gen=true

// PodSetStatus defines the observed state of PodSet
type PodSetStatus struct {
	Replicas int32    `json:"replicas"`
	PodNames []string `json:"podNames"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodSetList contains a list of PodSet
type PodSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PodSet{}, &PodSetList{})
}
