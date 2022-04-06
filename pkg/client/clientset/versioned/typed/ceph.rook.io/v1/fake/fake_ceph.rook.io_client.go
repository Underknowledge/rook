/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/rook/rook/pkg/client/clientset/versioned/typed/ceph.rook.io/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCephV1 struct {
	*testing.Fake
}

func (c *FakeCephV1) CephBlockPools(namespace string) v1.CephBlockPoolInterface {
	return &FakeCephBlockPools{c, namespace}
}

func (c *FakeCephV1) CephBlockPoolRadosNamespaces(namespace string) v1.CephBlockPoolRadosNamespaceInterface {
	return &FakeCephBlockPoolRadosNamespaces{c, namespace}
}

func (c *FakeCephV1) CephBucketNotifications(namespace string) v1.CephBucketNotificationInterface {
	return &FakeCephBucketNotifications{c, namespace}
}

func (c *FakeCephV1) CephBucketTopics(namespace string) v1.CephBucketTopicInterface {
	return &FakeCephBucketTopics{c, namespace}
}

func (c *FakeCephV1) CephClients(namespace string) v1.CephClientInterface {
	return &FakeCephClients{c, namespace}
}

func (c *FakeCephV1) CephClusters(namespace string) v1.CephClusterInterface {
	return &FakeCephClusters{c, namespace}
}

func (c *FakeCephV1) CephFilesystems(namespace string) v1.CephFilesystemInterface {
	return &FakeCephFilesystems{c, namespace}
}

func (c *FakeCephV1) CephFilesystemMirrors(namespace string) v1.CephFilesystemMirrorInterface {
	return &FakeCephFilesystemMirrors{c, namespace}
}

func (c *FakeCephV1) CephFilesystemSubVolumeGroups(namespace string) v1.CephFilesystemSubVolumeGroupInterface {
	return &FakeCephFilesystemSubVolumeGroups{c, namespace}
}

func (c *FakeCephV1) CephNFSes(namespace string) v1.CephNFSInterface {
	return &FakeCephNFSes{c, namespace}
}

func (c *FakeCephV1) CephObjectRealms(namespace string) v1.CephObjectRealmInterface {
	return &FakeCephObjectRealms{c, namespace}
}

func (c *FakeCephV1) CephObjectStores(namespace string) v1.CephObjectStoreInterface {
	return &FakeCephObjectStores{c, namespace}
}

func (c *FakeCephV1) CephObjectStoreUsers(namespace string) v1.CephObjectStoreUserInterface {
	return &FakeCephObjectStoreUsers{c, namespace}
}

func (c *FakeCephV1) CephObjectZones(namespace string) v1.CephObjectZoneInterface {
	return &FakeCephObjectZones{c, namespace}
}

func (c *FakeCephV1) CephObjectZoneGroups(namespace string) v1.CephObjectZoneGroupInterface {
	return &FakeCephObjectZoneGroups{c, namespace}
}

func (c *FakeCephV1) CephRBDMirrors(namespace string) v1.CephRBDMirrorInterface {
	return &FakeCephRBDMirrors{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCephV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
