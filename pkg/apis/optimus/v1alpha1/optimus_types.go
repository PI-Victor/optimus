/*
Copyright 2018 Cloudflavor Org contributors.
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

package v1alpha1

import (
	"time"
	
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// OptimusSpec defines the desired state of Optimus
// +k8s:openapi-gen=true
type OptimusSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Projects []Project `json:"projects"`
}

// OptimusStatus defines the observed state of Optimus
// +k8s:openapi-gen=true
type OptimusStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Optimus is the Schema for the optimus API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type Optimus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OptimusSpec   `json:"spec,omitempty"`
	Status OptimusStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OptimusList contains a list of Optimus
type OptimusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Optimus `json:"items"`
}

// Project holds the specifics of a pipeline project.
type Project struct {
	Name             string            `json:"name"`
	Namespace        string            `json:"namespace"`
	Username         string            `json:"username"`
	ArchiveArtifacts bool              `json:"archive"`
	Registry         ContainerRegistry `json:"registry"`
	Notifiers        []Notifier        `json:"notifiers,omitempty"`
	Repository       string            `json:"repository"`
	RunInterval      RunInterval       `json:"runInterval,omitempty"`

	Stages []Stage `json:"stages"`
}

// Stage represents a stage in the pipeline.
type Stage struct {
	Name     string `json:"name"`
	Parallel bool   `json:"parallel"`
	Notify   bool   `json:"notify"`

	Steps []Step `json:"steps"`

	Stat StageStatus `json:"status"`
}

// Step is a single command inside a stage.
type Step struct {
	Name         string `json:"name"`
	RuntimeImage string `json:"runtimeImage"`
	IgnoreErrors bool   `json:"ignoreError"`

	Cmd []string `json:"cmd"`
}

// ContainerRegistry holds information about a registry where an image will be
// pushed once an image has been built.
type ContainerRegistry struct {
	Username string `json:"username"`
	Secret   string `json:"secret"`
	URI      string `json:"uri"`
}

// Storage is the interface that is used for abstracting away the object storage
// client that will be used to store and archive artifacts. For now this will
// be represented by minio.
type Storage struct {
	URI string `json:"uri"`
}

// Notifier represents a webhook notification that is triggered by an
// action in the pipeline.
type Notifier struct {
	URI   string `json:"uri"`
	Token string `json:"token,omitempty"`
}

// RunInterval holds the time date interval for an automatic pipeline to run
// in cron format.
// TODO: add kubernetes cronjob types to this, avoid reinventing the wheel.
type RunInterval struct{}

// StageStatus represents the status of a stage.
type StageStatus struct {
	Duration  time.Duration
	StageName string
}

func init() {
	SchemeBuilder.Register(&Optimus{}, &OptimusList{})
}
