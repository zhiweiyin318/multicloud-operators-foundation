// Licensed Materials - Property of IBM
// 5737-E67
// (C) Copyright IBM Corporation 2016, 2019 All Rights Reserved
// US Government Users Restricted Rights - Use, duplication or disclosure restricted by GSA ADP Schedule Contract with IBM Corp.

package controllers

import (
	"testing"

	actionv1beta1 "github.com/open-cluster-management/multicloud-operators-foundation/pkg/apis/action/v1beta1"
	restutils "github.com/open-cluster-management/multicloud-operators-foundation/pkg/utils/rest"
	corev1 "k8s.io/api/core/v1"
	extensionv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

func NewKubeWorkSpec() *actionv1beta1.KubeWorkSpec {
	klusterletIngress := &extensionv1beta1.Ingress{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Ingress",
			APIVersion: "extension/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:              "mcm-ingress-testcluster-klusterlet",
			Namespace:         "kube-system",
			CreationTimestamp: metav1.Now(),
		},
		Spec: extensionv1beta1.IngressSpec{
			Rules: []extensionv1beta1.IngressRule{
				{
					Host: "test.com",
				},
			},
		},
		Status: extensionv1beta1.IngressStatus{
			LoadBalancer: corev1.LoadBalancerStatus{
				Ingress: []corev1.LoadBalancerIngress{
					{
						IP: "127.0.0.1",
					},
				},
			},
		},
	}
	return &actionv1beta1.KubeWorkSpec{
		Resource:  "Ingress",
		Namespace: "default",
		Name:      "test-deployment",
		ObjectTemplate: runtime.RawExtension{
			Object: klusterletIngress,
		},
	}
}
func NewClusterAction(name, namespace string, actiontype actionv1beta1.ActionType, kubework *actionv1beta1.KubeWorkSpec) *actionv1beta1.ClusterAction {
	return &actionv1beta1.ClusterAction{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: actionv1beta1.ClusterActionSpec{
			ActionType: actiontype,
			KubeWork:   kubework,
		},
	}
}
func TestActionReconciler_handleClusterAction(t *testing.T) {
	fakekubecontrol := restutils.NewFakeKubeControl()
	log := ctrl.Log.WithName("controllers").WithName("ClusterAction")
	ar := NewActionReconciler(nil, log, &runtime.Scheme{}, nil, fakekubecontrol, false)
	kubework := NewKubeWorkSpec()
	nca1 := NewClusterAction("ca1", "can1", actionv1beta1.CreateActionType, kubework)
	err := ar.handleClusterAction(nca1)
	if err != nil {
		t.Errorf("Create kube work error. %v", err)
	}
	nca2 := NewClusterAction("ca1", "can1", actionv1beta1.UpdateActionType, kubework)
	err = ar.handleClusterAction(nca2)
	if err == nil {
		t.Error("Update kube work should have error.")
	}
	nca3 := NewClusterAction("ca1", "can1", actionv1beta1.DeleteActionType, kubework)
	err = ar.handleClusterAction(nca3)
	if err != nil {
		t.Errorf("Create kube work error. %v", err)
	}

	arf := NewActionReconciler(nil, log, &runtime.Scheme{}, nil, fakekubecontrol, true)
	err = arf.handleClusterAction(nca1)
	if err != nil {
		t.Errorf("Create kube work error. %v", err)
	}
	err = arf.handleClusterAction(nca2)
	if err == nil {
		t.Error("Update kube work should have error.")
	}
	err = arf.handleClusterAction(nca3)
	if err != nil {
		t.Errorf("Create kube work error. %v", err)
	}
}