package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ObjectBucketSource struct {
	// Host the host URL of the object store with
	Host string `json:"host"`
	// Region the region of the bucket within an object store
	Region string `json:"region"`
	// Port the insecure port number of the object store, if it exists
	Port int `json:"port"`
	// SecrePort the secure port number of the object store, if it exists
	SecurePort int `json:"securePort"`
	// SSL true if the connection is secured with SSL, false if it is not.
	SSL bool `json:"ssl"`
	// SupportsTenants true if the object store provider supports the use of Tenants
	Tenant string `json:"tenant,omitempty"`
	// SuppotsNamespace true if the object store provider supports Namespaces
	Namespace string `json:"namespace,omitempty"`
	// Versioned true if the object store support versioned buckets, false if not
	Versioned bool `json:"versioned,omitempty"`
}

// ObjectBucketSpec defines the desired state of ObjectBucket
type ObjectBucketSpec struct {
	Provisioner string
	BucketName  string

	ObjectBucketSource ObjectBucketSource
}

// ObjectBucketStatus defines the observed state of ObjectBucket
type ObjectBucketStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ObjectBucket is the Schema for the objectbuckets API
// +k8s:openapi-gen=true
type ObjectBucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ObjectBucketSpec   `json:"spec,omitempty"`
	Status ObjectBucketStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ObjectBucketList contains a list of ObjectBucket
type ObjectBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectBucket `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ObjectBucket{}, &ObjectBucketList{})
}
