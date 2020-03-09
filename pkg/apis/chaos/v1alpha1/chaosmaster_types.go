package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ChaosmasterSpec defines the desired state of Chaosmaster
type ChaosmasterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	BlastRadius int32 `json:"BlastRadius"`
	ChaosType string `json:"ChaosType"`
	ServiceName string `json:"ServiceName"`
	Platform string `json:"Platform"`
	AppId string `json:"AppId"`

}

// ChaosmasterStatus defines the observed state of Chaosmaster
type ChaosmasterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	ChaosStatus `json:",inline"`

}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Chaosmaster is the Schema for the chaosmasters API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=chaosmasters,scope=Namespaced
type Chaosmaster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ChaosmasterSpec   `json:"spec,omitempty"`
	Status ChaosmasterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChaosmasterList contains a list of Chaosmaster
type ChaosmasterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Chaosmaster `json:"items"`
}



func init() {
	SchemeBuilder.Register(&Chaosmaster{}, &ChaosmasterList{})
}
