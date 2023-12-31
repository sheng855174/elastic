// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/module_include_list/module_include_list.go - DO NOT EDIT.

package include

import (
	// Import packages that need to register themselves.
	_ "github.com/sheng855174/elastic/metricbeat/module/aerospike"
	_ "github.com/sheng855174/elastic/metricbeat/module/aerospike/namespace"
	_ "github.com/sheng855174/elastic/metricbeat/module/apache"
	_ "github.com/sheng855174/elastic/metricbeat/module/apache/status"
	_ "github.com/sheng855174/elastic/metricbeat/module/beat"
	_ "github.com/sheng855174/elastic/metricbeat/module/beat/state"
	_ "github.com/sheng855174/elastic/metricbeat/module/beat/stats"
	_ "github.com/sheng855174/elastic/metricbeat/module/ceph"
	_ "github.com/sheng855174/elastic/metricbeat/module/ceph/cluster_disk"
	_ "github.com/sheng855174/elastic/metricbeat/module/ceph/cluster_health"
	_ "github.com/sheng855174/elastic/metricbeat/module/ceph/cluster_status"
	_ "github.com/sheng855174/elastic/metricbeat/module/ceph/mgr_cluster_disk"
	_ "github.com/sheng855174/elastic/metricbeat/module/ceph/mgr_cluster_health"
	_ "github.com/sheng855174/elastic/metricbeat/module/ceph/mgr_osd_perf"
	_ "github.com/sheng855174/elastic/metricbeat/module/ceph/mgr_osd_pool_stats"
	_ "github.com/sheng855174/elastic/metricbeat/module/ceph/mgr_osd_tree"
	_ "github.com/sheng855174/elastic/metricbeat/module/ceph/mgr_pool_disk"
	_ "github.com/sheng855174/elastic/metricbeat/module/ceph/monitor_health"
	_ "github.com/sheng855174/elastic/metricbeat/module/ceph/osd_df"
	_ "github.com/sheng855174/elastic/metricbeat/module/ceph/osd_tree"
	_ "github.com/sheng855174/elastic/metricbeat/module/ceph/pool_disk"
	_ "github.com/sheng855174/elastic/metricbeat/module/consul"
	_ "github.com/sheng855174/elastic/metricbeat/module/consul/agent"
	_ "github.com/sheng855174/elastic/metricbeat/module/couchbase"
	_ "github.com/sheng855174/elastic/metricbeat/module/couchbase/bucket"
	_ "github.com/sheng855174/elastic/metricbeat/module/couchbase/cluster"
	_ "github.com/sheng855174/elastic/metricbeat/module/couchbase/node"
	_ "github.com/sheng855174/elastic/metricbeat/module/couchdb"
	_ "github.com/sheng855174/elastic/metricbeat/module/couchdb/server"
	_ "github.com/sheng855174/elastic/metricbeat/module/dropwizard"
	_ "github.com/sheng855174/elastic/metricbeat/module/dropwizard/collector"
	_ "github.com/sheng855174/elastic/metricbeat/module/elasticsearch"
	_ "github.com/sheng855174/elastic/metricbeat/module/elasticsearch/ccr"
	_ "github.com/sheng855174/elastic/metricbeat/module/elasticsearch/cluster_stats"
	_ "github.com/sheng855174/elastic/metricbeat/module/elasticsearch/enrich"
	_ "github.com/sheng855174/elastic/metricbeat/module/elasticsearch/index"
	_ "github.com/sheng855174/elastic/metricbeat/module/elasticsearch/index_recovery"
	_ "github.com/sheng855174/elastic/metricbeat/module/elasticsearch/index_summary"
	_ "github.com/sheng855174/elastic/metricbeat/module/elasticsearch/ml_job"
	_ "github.com/sheng855174/elastic/metricbeat/module/elasticsearch/node"
	_ "github.com/sheng855174/elastic/metricbeat/module/elasticsearch/node_stats"
	_ "github.com/sheng855174/elastic/metricbeat/module/elasticsearch/pending_tasks"
	_ "github.com/sheng855174/elastic/metricbeat/module/elasticsearch/shard"
	_ "github.com/sheng855174/elastic/metricbeat/module/envoyproxy"
	_ "github.com/sheng855174/elastic/metricbeat/module/envoyproxy/server"
	_ "github.com/sheng855174/elastic/metricbeat/module/etcd"
	_ "github.com/sheng855174/elastic/metricbeat/module/etcd/leader"
	_ "github.com/sheng855174/elastic/metricbeat/module/etcd/metrics"
	_ "github.com/sheng855174/elastic/metricbeat/module/etcd/self"
	_ "github.com/sheng855174/elastic/metricbeat/module/etcd/store"
	_ "github.com/sheng855174/elastic/metricbeat/module/golang"
	_ "github.com/sheng855174/elastic/metricbeat/module/golang/expvar"
	_ "github.com/sheng855174/elastic/metricbeat/module/golang/heap"
	_ "github.com/sheng855174/elastic/metricbeat/module/graphite"
	_ "github.com/sheng855174/elastic/metricbeat/module/graphite/server"
	_ "github.com/sheng855174/elastic/metricbeat/module/haproxy"
	_ "github.com/sheng855174/elastic/metricbeat/module/haproxy/info"
	_ "github.com/sheng855174/elastic/metricbeat/module/haproxy/stat"
	_ "github.com/sheng855174/elastic/metricbeat/module/http"
	_ "github.com/sheng855174/elastic/metricbeat/module/http/json"
	_ "github.com/sheng855174/elastic/metricbeat/module/http/server"
	_ "github.com/sheng855174/elastic/metricbeat/module/jolokia"
	_ "github.com/sheng855174/elastic/metricbeat/module/jolokia/jmx"
	_ "github.com/sheng855174/elastic/metricbeat/module/kafka"
	_ "github.com/sheng855174/elastic/metricbeat/module/kafka/consumergroup"
	_ "github.com/sheng855174/elastic/metricbeat/module/kafka/partition"
	_ "github.com/sheng855174/elastic/metricbeat/module/kibana"
	_ "github.com/sheng855174/elastic/metricbeat/module/kibana/stats"
	_ "github.com/sheng855174/elastic/metricbeat/module/kibana/status"
	_ "github.com/sheng855174/elastic/metricbeat/module/kvm"
	_ "github.com/sheng855174/elastic/metricbeat/module/kvm/dommemstat"
	_ "github.com/sheng855174/elastic/metricbeat/module/kvm/status"
	_ "github.com/sheng855174/elastic/metricbeat/module/linux"
	_ "github.com/sheng855174/elastic/metricbeat/module/linux/conntrack"
	_ "github.com/sheng855174/elastic/metricbeat/module/linux/iostat"
	_ "github.com/sheng855174/elastic/metricbeat/module/linux/ksm"
	_ "github.com/sheng855174/elastic/metricbeat/module/linux/memory"
	_ "github.com/sheng855174/elastic/metricbeat/module/linux/pageinfo"
	_ "github.com/sheng855174/elastic/metricbeat/module/linux/pressure"
	_ "github.com/sheng855174/elastic/metricbeat/module/logstash"
	_ "github.com/sheng855174/elastic/metricbeat/module/logstash/node"
	_ "github.com/sheng855174/elastic/metricbeat/module/logstash/node_stats"
	_ "github.com/sheng855174/elastic/metricbeat/module/memcached"
	_ "github.com/sheng855174/elastic/metricbeat/module/memcached/stats"
	_ "github.com/sheng855174/elastic/metricbeat/module/mongodb"
	_ "github.com/sheng855174/elastic/metricbeat/module/mongodb/collstats"
	_ "github.com/sheng855174/elastic/metricbeat/module/mongodb/dbstats"
	_ "github.com/sheng855174/elastic/metricbeat/module/mongodb/metrics"
	_ "github.com/sheng855174/elastic/metricbeat/module/mongodb/replstatus"
	_ "github.com/sheng855174/elastic/metricbeat/module/mongodb/status"
	_ "github.com/sheng855174/elastic/metricbeat/module/munin"
	_ "github.com/sheng855174/elastic/metricbeat/module/munin/node"
	_ "github.com/sheng855174/elastic/metricbeat/module/mysql"
	_ "github.com/sheng855174/elastic/metricbeat/module/mysql/galera_status"
	_ "github.com/sheng855174/elastic/metricbeat/module/mysql/query"
	_ "github.com/sheng855174/elastic/metricbeat/module/mysql/status"
	_ "github.com/sheng855174/elastic/metricbeat/module/nats"
	_ "github.com/sheng855174/elastic/metricbeat/module/nats/connection"
	_ "github.com/sheng855174/elastic/metricbeat/module/nats/connections"
	_ "github.com/sheng855174/elastic/metricbeat/module/nats/route"
	_ "github.com/sheng855174/elastic/metricbeat/module/nats/routes"
	_ "github.com/sheng855174/elastic/metricbeat/module/nats/stats"
	_ "github.com/sheng855174/elastic/metricbeat/module/nats/subscriptions"
	_ "github.com/sheng855174/elastic/metricbeat/module/nginx"
	_ "github.com/sheng855174/elastic/metricbeat/module/nginx/stubstatus"
	_ "github.com/sheng855174/elastic/metricbeat/module/openmetrics"
	_ "github.com/sheng855174/elastic/metricbeat/module/openmetrics/collector"
	_ "github.com/sheng855174/elastic/metricbeat/module/php_fpm"
	_ "github.com/sheng855174/elastic/metricbeat/module/php_fpm/pool"
	_ "github.com/sheng855174/elastic/metricbeat/module/php_fpm/process"
	_ "github.com/sheng855174/elastic/metricbeat/module/postgresql"
	_ "github.com/sheng855174/elastic/metricbeat/module/postgresql/activity"
	_ "github.com/sheng855174/elastic/metricbeat/module/postgresql/bgwriter"
	_ "github.com/sheng855174/elastic/metricbeat/module/postgresql/database"
	_ "github.com/sheng855174/elastic/metricbeat/module/postgresql/statement"
	_ "github.com/sheng855174/elastic/metricbeat/module/prometheus"
	_ "github.com/sheng855174/elastic/metricbeat/module/prometheus/collector"
	_ "github.com/sheng855174/elastic/metricbeat/module/prometheus/query"
	_ "github.com/sheng855174/elastic/metricbeat/module/prometheus/remote_write"
	_ "github.com/sheng855174/elastic/metricbeat/module/rabbitmq"
	_ "github.com/sheng855174/elastic/metricbeat/module/rabbitmq/connection"
	_ "github.com/sheng855174/elastic/metricbeat/module/rabbitmq/exchange"
	_ "github.com/sheng855174/elastic/metricbeat/module/rabbitmq/node"
	_ "github.com/sheng855174/elastic/metricbeat/module/rabbitmq/queue"
	_ "github.com/sheng855174/elastic/metricbeat/module/redis"
	_ "github.com/sheng855174/elastic/metricbeat/module/redis/info"
	_ "github.com/sheng855174/elastic/metricbeat/module/redis/key"
	_ "github.com/sheng855174/elastic/metricbeat/module/redis/keyspace"
	_ "github.com/sheng855174/elastic/metricbeat/module/system"
	_ "github.com/sheng855174/elastic/metricbeat/module/system/core"
	_ "github.com/sheng855174/elastic/metricbeat/module/system/cpu"
	_ "github.com/sheng855174/elastic/metricbeat/module/system/diskio"
	_ "github.com/sheng855174/elastic/metricbeat/module/system/entropy"
	_ "github.com/sheng855174/elastic/metricbeat/module/system/filesystem"
	_ "github.com/sheng855174/elastic/metricbeat/module/system/fsstat"
	_ "github.com/sheng855174/elastic/metricbeat/module/system/load"
	_ "github.com/sheng855174/elastic/metricbeat/module/system/memory"
	_ "github.com/sheng855174/elastic/metricbeat/module/system/network"
	_ "github.com/sheng855174/elastic/metricbeat/module/system/network_summary"
	_ "github.com/sheng855174/elastic/metricbeat/module/system/process"
	_ "github.com/sheng855174/elastic/metricbeat/module/system/process_summary"
	_ "github.com/sheng855174/elastic/metricbeat/module/system/raid"
	_ "github.com/sheng855174/elastic/metricbeat/module/system/service"
	_ "github.com/sheng855174/elastic/metricbeat/module/system/socket"
	_ "github.com/sheng855174/elastic/metricbeat/module/system/socket_summary"
	_ "github.com/sheng855174/elastic/metricbeat/module/system/uptime"
	_ "github.com/sheng855174/elastic/metricbeat/module/system/users"
	_ "github.com/sheng855174/elastic/metricbeat/module/traefik"
	_ "github.com/sheng855174/elastic/metricbeat/module/traefik/health"
	_ "github.com/sheng855174/elastic/metricbeat/module/uwsgi"
	_ "github.com/sheng855174/elastic/metricbeat/module/uwsgi/status"
	_ "github.com/sheng855174/elastic/metricbeat/module/vsphere"
	_ "github.com/sheng855174/elastic/metricbeat/module/vsphere/datastore"
	_ "github.com/sheng855174/elastic/metricbeat/module/vsphere/host"
	_ "github.com/sheng855174/elastic/metricbeat/module/vsphere/virtualmachine"
	_ "github.com/sheng855174/elastic/metricbeat/module/windows"
	_ "github.com/sheng855174/elastic/metricbeat/module/windows/perfmon"
	_ "github.com/sheng855174/elastic/metricbeat/module/windows/service"
	_ "github.com/sheng855174/elastic/metricbeat/module/zookeeper"
	_ "github.com/sheng855174/elastic/metricbeat/module/zookeeper/connection"
	_ "github.com/sheng855174/elastic/metricbeat/module/zookeeper/mntr"
	_ "github.com/sheng855174/elastic/metricbeat/module/zookeeper/server"
)
