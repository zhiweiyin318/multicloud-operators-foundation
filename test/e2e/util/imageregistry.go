package util

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

const imageRegistryTemplate = `{
  "apiVersion": "cluster.open-cluster-management.io/v1alpha1",
  "kind": "ManagedClusterImageRegistry",
  "metadata": {
 	"labels": {
	  "test-automation": "true"
    },
    "namespace": "cluster1"
	"name": "registry1"
  },
  "spec": {
    "registry": "quay.io/yzw",
    "pullSecret" {
      "name": "pullSecret"
    },
    "placementRef": {
      "group": "cluster.open-cluster-management.io",
      "resource": "placements",
      "name": "placement"
    }
  }
}`

var imageRegistryGVR = schema.GroupVersionResource{
	Group:    "cluster.open-cluster-management.io",
	Version:  "v1alpha1",
	Resource: "managedclusterimageregistries",
}

func CreateImageRegistry(dynamicClient dynamic.Interface, namespace, name, placement string) error {
	obj, err := LoadResourceFromJSON(imageRegistryTemplate)
	if err != nil {
		return err
	}
	err = unstructured.SetNestedField(obj.Object, namespace, "metadata", "namespace")
	if err != nil {
		return err
	}
	err = unstructured.SetNestedField(obj.Object, name, "metadata", "name")
	if err != nil {
		return err
	}
	err = unstructured.SetNestedField(obj.Object, placement, "spec", "placementRef", "name")
	if err != nil {
		return err
	}

	_, err = CreateResource(dynamicClient, imageRegistryGVR, obj)
	return err
}

func DeleteImageRegistry(dynamicClient dynamic.Interface, namespace, name string) error {
	return DeleteResource(dynamicClient, imageRegistryGVR, namespace, name)
}
