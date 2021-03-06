/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	informers "k8s.io/client-go/informers"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/client-go/listers/rbac/v1beta1"
)

// Interface provides access to all the listers in this group version.
type Interface interface { // ClusterRoles returns a ClusterRoleLister
	ClusterRoles() v1beta1.ClusterRoleLister
	// ClusterRoleBindings returns a ClusterRoleBindingLister
	ClusterRoleBindings() v1beta1.ClusterRoleBindingLister
	// Roles returns a RoleLister
	Roles() v1beta1.RoleLister
	// RoleBindings returns a RoleBindingLister
	RoleBindings() v1beta1.RoleBindingLister
}

type version struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

type infromerVersion struct {
	factory informers.SharedInformerFactory
}

// New returns a new Interface.
func New(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{client: client, tweakListOptions: tweakListOptions}
}

// NewFrom returns a new Interface.
func NewFrom(factory informers.SharedInformerFactory) Interface {
	return &infromerVersion{factory: factory}
}

// ClusterRoles returns a ClusterRoleLister.
func (v *version) ClusterRoles() v1beta1.ClusterRoleLister {
	return &clusterRoleLister{client: v.client, tweakListOptions: v.tweakListOptions}
}

// ClusterRoles returns a ClusterRoleLister.
func (v *infromerVersion) ClusterRoles() v1beta1.ClusterRoleLister {
	return v.factory.Rbac().V1beta1().ClusterRoles().Lister()
}

// ClusterRoleBindings returns a ClusterRoleBindingLister.
func (v *version) ClusterRoleBindings() v1beta1.ClusterRoleBindingLister {
	return &clusterRoleBindingLister{client: v.client, tweakListOptions: v.tweakListOptions}
}

// ClusterRoleBindings returns a ClusterRoleBindingLister.
func (v *infromerVersion) ClusterRoleBindings() v1beta1.ClusterRoleBindingLister {
	return v.factory.Rbac().V1beta1().ClusterRoleBindings().Lister()
}

// Roles returns a RoleLister.
func (v *version) Roles() v1beta1.RoleLister {
	return &roleLister{client: v.client, tweakListOptions: v.tweakListOptions}
}

// Roles returns a RoleLister.
func (v *infromerVersion) Roles() v1beta1.RoleLister {
	return v.factory.Rbac().V1beta1().Roles().Lister()
}

// RoleBindings returns a RoleBindingLister.
func (v *version) RoleBindings() v1beta1.RoleBindingLister {
	return &roleBindingLister{client: v.client, tweakListOptions: v.tweakListOptions}
}

// RoleBindings returns a RoleBindingLister.
func (v *infromerVersion) RoleBindings() v1beta1.RoleBindingLister {
	return v.factory.Rbac().V1beta1().RoleBindings().Lister()
}
