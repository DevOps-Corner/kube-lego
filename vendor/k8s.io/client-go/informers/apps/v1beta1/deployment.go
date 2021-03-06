/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package v1beta1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/client-go/listers/apps/v1beta1"
	apps_v1beta1 "k8s.io/client-go/pkg/apis/apps/v1beta1"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// DeploymentInformer provides access to a shared informer and lister for
// Deployments.
type DeploymentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.DeploymentLister
}

type deploymentInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newDeploymentInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.AppsV1beta1().Deployments(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.AppsV1beta1().Deployments(v1.NamespaceAll).Watch(options)
			},
		},
		&apps_v1beta1.Deployment{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *deploymentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apps_v1beta1.Deployment{}, newDeploymentInformer)
}

func (f *deploymentInformer) Lister() v1beta1.DeploymentLister {
	return v1beta1.NewDeploymentLister(f.Informer().GetIndexer())
}
