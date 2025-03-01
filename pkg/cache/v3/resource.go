// Copyright 2018 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package cache

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/golang/protobuf/proto"

	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	auth "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	runtime "github.com/envoyproxy/go-control-plane/envoy/service/runtime/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
)

// GetResponseType returns the enumeration for a valid xDS type URL
func GetResponseType(typeURL string) types.ResponseType {
	switch typeURL {
	case resource.EndpointType:
		return types.Endpoint
	case resource.ClusterType:
		return types.Cluster
	case resource.RouteType:
		return types.Route
	case resource.ListenerType:
		return types.Listener
	case resource.SecretType:
		return types.Secret
	case resource.RuntimeType:
		return types.Runtime
	case resource.ExtensionConfigType:
		return types.ExtensionConfig
	}
	return types.UnknownType
}

// GetResourceName returns the resource name for a valid xDS response type.
func GetResourceName(res types.Resource) string {
	switch v := res.(type) {
	case *endpoint.ClusterLoadAssignment:
		return v.GetClusterName()
	case *cluster.Cluster:
		return v.GetName()
	case *route.RouteConfiguration:
		return v.GetName()
	case *listener.Listener:
		return v.GetName()
	case *auth.Secret:
		return v.GetName()
	case *runtime.Runtime:
		return v.GetName()
	case *core.TypedExtensionConfig:
		// This is a V3 proto, but this is the easiest workaround for the fact that there is no V2 proto.
		return v.GetName()
	default:
		return ""
	}
}

// MarshalResource converts the Resource to MarshaledResource
func MarshalResource(resource types.Resource) (types.MarshaledResource, error) {
	b := proto.NewBuffer(nil)
	b.SetDeterministic(true)
	err := b.Marshal(resource)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// GetResourceReferences returns the names for dependent resources (EDS cluster
// names for CDS, RDS routes names for LDS).
func GetResourceReferences(resources map[string]types.ResourceWithTtl) map[string]bool {
	out := make(map[string]bool)
	for _, res := range resources {
		if res.Resource == nil {
			continue
		}
		switch v := res.Resource.(type) {
		case *endpoint.ClusterLoadAssignment:
			// no dependencies
		case *cluster.Cluster:
			// for EDS type, use cluster name or ServiceName override
			switch typ := v.ClusterDiscoveryType.(type) {
			case *cluster.Cluster_Type:
				if typ.Type == cluster.Cluster_EDS {
					if v.EdsClusterConfig != nil && v.EdsClusterConfig.ServiceName != "" {
						out[v.EdsClusterConfig.ServiceName] = true
					} else {
						out[v.Name] = true
					}
				}
			}
		case *route.RouteConfiguration:
			// References to clusters in both routes (and listeners) are not included
			// in the result, because the clusters are retrieved in bulk currently,
			// and not by name.
		case *listener.Listener:
			// extract route configuration names from HTTP connection manager
			for _, chain := range v.FilterChains {
				for _, filter := range chain.Filters {
					if filter.Name != wellknown.HTTPConnectionManager {
						continue
					}

					config := resource.GetHTTPConnectionManager(filter)

					if config == nil {
						continue
					}

					if rds, ok := config.RouteSpecifier.(*hcm.HttpConnectionManager_Rds); ok && rds != nil && rds.Rds != nil {
						out[rds.Rds.RouteConfigName] = true
					}
				}
			}
		case *runtime.Runtime:
			// no dependencies
		}
	}
	return out
}

// HashResource will take a resource and create a SHA256 hash sum out of the marshaled bytes
func HashResource(resource []byte) string {
	hasher := sha256.New()
	hasher.Write(resource)

	return hex.EncodeToString(hasher.Sum(nil))
}
