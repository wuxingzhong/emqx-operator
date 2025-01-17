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

package v1beta1

import (
	"fmt"
	"reflect"

	v1beta2 "github.com/emqx/emqx-operator/apis/apps/v1beta2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EmqxEnterpriseSpec defines the desired state of EmqxEnterprise
type EmqxEnterpriseSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The fields of Broker.
	//The replicas of emqx broker
	//+kubebuilder:default:=3
	Replicas *int32 `json:"replicas,omitempty"`

	//+kubebuilder:validation:Required
	Image            string                        `json:"image,omitempty"`
	ImagePullPolicy  corev1.PullPolicy             `json:"imagePullPolicy,omitempty"`
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`

	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// The service account name which is being bind with the service
	// account of the crd instance.
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	Storage *Storage `json:"storage,omitempty"`

	NodeName     string            `json:"nodeName,omitempty"`
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// TODO: waiting to be deleted, should use meta.labels
	Labels map[string]string `json:"labels,omitempty"`
	// TODO: waiting to be deleted, should use meta.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	Listener v1beta2.Listener `json:"listener,omitempty"`
	License  string           `json:"license,omitempty"`

	Affinity    *corev1.Affinity    `json:"affinity,omitempty"`
	ToleRations []corev1.Toleration `json:"toleRations,omitempty"`

	ExtraVolumes      []corev1.Volume      `json:"extraVolumes,omitempty"`
	ExtraVolumeMounts []corev1.VolumeMount `json:"extraVolumeMounts,omitempty"`

	Env []corev1.EnvVar `json:"env,omitempty"`

	ACL []v1beta2.ACL `json:"acl,omitempty"`

	Plugins []v1beta2.Plugin `json:"plugins,omitempty"`

	Modules []v1beta2.EmqxEnterpriseModules `json:"modules,omitempty"`

	SecurityContext *corev1.PodSecurityContext `json:"securityContext,omitempty"`

	TelegrafTemplate *v1beta2.TelegrafTemplate `json:"telegrafTemplate,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:resource:shortName=emqx-ee
//+kubebuilder:subresource:status
//+kubebuilder:subresource:scale:specpath=.spec.replicas,statuspath=.status.replicas
//+kubebuilder:unservedversion

// EmqxEnterprise is the Schema for the emqxenterprises API
type EmqxEnterprise struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec           EmqxEnterpriseSpec `json:"spec,omitempty"`
	v1beta2.Status `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
// EmqxEnterpriseList contains a list of EmqxEnterprise
type EmqxEnterpriseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmqxEnterprise `json:"items"`
}

func (emqx *EmqxEnterprise) String() string {
	return fmt.Sprintf("EmqxEnterprise instance [%s],Image [%s]",
		emqx.ObjectMeta.Name,
		emqx.Spec.Image,
	)
}

func init() {
	SchemeBuilder.Register(&EmqxEnterprise{}, &EmqxEnterpriseList{})
}

func (emqx *EmqxEnterprise) GetAPIVersion() string        { return emqx.APIVersion }
func (emqx *EmqxEnterprise) SetAPIVersion(version string) { emqx.APIVersion = version }
func (emqx *EmqxEnterprise) GetKind() string              { return emqx.Kind }
func (emqx *EmqxEnterprise) SetKind(kind string)          { emqx.Kind = kind }

func (emqx *EmqxEnterprise) GetReplicas() *int32 {
	if reflect.ValueOf(emqx.Spec.Replicas).IsZero() {
		defaultReplicas := int32(3)
		emqx.SetReplicas(&defaultReplicas)
	}
	return emqx.Spec.Replicas
}
func (emqx *EmqxEnterprise) SetReplicas(replicas *int32) { emqx.Spec.Replicas = replicas }

func (emqx *EmqxEnterprise) GetImage() string      { return emqx.Spec.Image }
func (emqx *EmqxEnterprise) SetImage(image string) { emqx.Spec.Image = image }

func (emqx *EmqxEnterprise) GetImagePullPolicy() corev1.PullPolicy { return emqx.Spec.ImagePullPolicy }
func (emqx *EmqxEnterprise) SetImagePullPolicy(pullPolicy corev1.PullPolicy) {
	emqx.Spec.ImagePullPolicy = pullPolicy
}

func (emqx *EmqxEnterprise) GetImagePullSecrets() []corev1.LocalObjectReference {
	return emqx.Spec.ImagePullSecrets
}
func (emqx *EmqxEnterprise) SetImagePullSecrets(imagePullSecrets []corev1.LocalObjectReference) {
	emqx.Spec.ImagePullSecrets = imagePullSecrets
}

func (emqx *EmqxEnterprise) GetServiceAccountName() string {
	if emqx.Spec.ServiceAccountName == "" {
		emqx.SetServiceAccountName(emqx.Name)
	}
	return emqx.Spec.ServiceAccountName
}
func (emqx *EmqxEnterprise) SetServiceAccountName(serviceAccountName string) {
	emqx.Spec.ServiceAccountName = serviceAccountName
}

func (emqx *EmqxEnterprise) GetResource() corev1.ResourceRequirements { return emqx.Spec.Resources }
func (emqx *EmqxEnterprise) SetResource(resource corev1.ResourceRequirements) {
	emqx.Spec.Resources = resource
}

func (emqx *EmqxEnterprise) GetLicense() string        { return emqx.Spec.License }
func (emqx *EmqxEnterprise) SetLicense(license string) { emqx.Spec.License = license }

func (emqx *EmqxEnterprise) GetStorage() *Storage        { return emqx.Spec.Storage }
func (emqx *EmqxEnterprise) SetStorage(storage *Storage) { emqx.Spec.Storage = storage }

func (emqx *EmqxEnterprise) GetNodeName() string { return emqx.Spec.NodeName }
func (emqx *EmqxEnterprise) SetNodeName(nodeName string) {
	emqx.Spec.NodeName = nodeName
}

func (emqx *EmqxEnterprise) GetNodeSelector() map[string]string { return emqx.Spec.NodeSelector }
func (emqx *EmqxEnterprise) SetNodeSelector(nodeSelector map[string]string) {
	emqx.Spec.NodeSelector = nodeSelector
}

func (emqx *EmqxEnterprise) GetAnnotations() map[string]string { return emqx.Spec.Annotations }
func (emqx *EmqxEnterprise) SetAnnotations(annotations map[string]string) {
	emqx.Spec.Annotations = annotations
}

func (emqx *EmqxEnterprise) GetListener() v1beta2.Listener { return emqx.Spec.Listener }
func (emqx *EmqxEnterprise) SetListener(listener v1beta2.Listener) {
	emqx.Spec.Listener = listener
}

func (emqx *EmqxEnterprise) GetAffinity() *corev1.Affinity         { return emqx.Spec.Affinity }
func (emqx *EmqxEnterprise) SetAffinity(affinity *corev1.Affinity) { emqx.Spec.Affinity = affinity }

func (emqx *EmqxEnterprise) GetToleRations() []corev1.Toleration { return emqx.Spec.ToleRations }
func (emqx *EmqxEnterprise) SetToleRations(tolerations []corev1.Toleration) {
	emqx.Spec.ToleRations = tolerations
}

func (emqx *EmqxEnterprise) GetExtraVolumes() []corev1.Volume { return emqx.Spec.ExtraVolumes }
func (emqx *EmqxEnterprise) GetExtraVolumeMounts() []corev1.VolumeMount {
	return emqx.Spec.ExtraVolumeMounts
}

func (emqx *EmqxEnterprise) GetACL() []v1beta2.ACL { return emqx.Spec.ACL }
func (emqx *EmqxEnterprise) SetACL(acl []v1beta2.ACL) {
	emqx.Spec.ACL = acl
}

func (emqx *EmqxEnterprise) GetEnv() []corev1.EnvVar { return emqx.Spec.Env }
func (emqx *EmqxEnterprise) SetEnv(env []corev1.EnvVar) {
	emqx.Spec.Env = env
}

func (emqx *EmqxEnterprise) GetPlugins() []v1beta2.Plugin { return emqx.Spec.Plugins }
func (emqx *EmqxEnterprise) SetPlugins(plugins []v1beta2.Plugin) {
	emqx.Spec.Plugins = plugins
}

func (emqx *EmqxEnterprise) GetModules() []v1beta2.EmqxEnterpriseModules { return emqx.Spec.Modules }
func (emqx *EmqxEnterprise) SetModules(modules []v1beta2.EmqxEnterpriseModules) {
	emqx.Spec.Modules = modules
}

func (emqx *EmqxEnterprise) GetHeadlessServiceName() string {
	return fmt.Sprintf("%s-%s", emqx.Name, "headless")
}

func (emqx *EmqxEnterprise) GetSecurityContext() *corev1.PodSecurityContext {
	return emqx.Spec.SecurityContext
}
func (emqx *EmqxEnterprise) SetSecurityContext(securityContext *corev1.PodSecurityContext) {
	emqx.Spec.SecurityContext = securityContext
}

func (emqx *EmqxEnterprise) GetTelegrafTemplate() *v1beta2.TelegrafTemplate {
	return emqx.Spec.TelegrafTemplate
}
func (emqx *EmqxEnterprise) SetTelegrafTemplate(telegrafTemplate *v1beta2.TelegrafTemplate) {
	emqx.Spec.TelegrafTemplate = telegrafTemplate
}
