// Licensed under the Apache License, Version 2.0
// Details: https://raw.githubusercontent.com/maniksurtani/quotaservice/master/LICENSE

package admin

import (
	pb "github.com/maniksurtani/quotaservice/protos/config"
	"github.com/maniksurtani/quotaservice/stats"
)

// Administrable defines something that can be administered via this package.
type Administrable interface {
	Configs() *pb.ServiceConfig
	HistoricalConfigs() ([]*pb.ServiceConfig, error)

	UpdateConfig(*pb.ServiceConfig, string) error

	DeleteBucket(string, string) error
	AddBucket(string, *pb.BucketConfig) error
	UpdateBucket(string, *pb.BucketConfig) error

	DeleteNamespace(string) error
	AddNamespace(*pb.NamespaceConfig) error
	UpdateNamespace(*pb.NamespaceConfig) error

	TopDynamicHits(string) []*stats.BucketScore
	TopDynamicMisses(string) []*stats.BucketScore
	DynamicBucketStats(string, string) *stats.BucketScores
}
