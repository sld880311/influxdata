/*
Copyright 2014 The Kubernetes Authors.

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

// The types defined in this package are used by the meta package to represent
// the in-memory version of objects. We cannot reuse the __internal version of
// API objects because it causes import cycle.
package metatypes

import "k8s.io/client-go/1.5/pkg/types"

type OwnerReference struct {
	APIVersion string
	Kind       string
	UID        types.UID
	Name       string
	Controller *bool
}
