package common

import (
	"context"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	clusterPkg "github.com/openshift/assisted-service/internal/cluster"
	"github.com/openshift/assisted-service/internal/host"
	"github.com/openshift/assisted-service/pkg/leader"
	"github.com/openshift/assisted-service/pkg/s3wrapper"
	"github.com/openshift/assisted-service/restapi/operations/installer"
)

type Config struct {
	DeletedUnregisteredAfter        time.Duration     `envconfig:"DELETED_UNREGISTERED_AFTER" default:"168h"`
	DeregisterInactiveAfter         time.Duration     `envconfig:"DELETED_INACTIVE_AFTER" default:"720h"`
}

type GarbageCollectors interface {
	InstallClusterInternal(ctx context.Context, params installer.InstallClusterParams) (*Cluster, error)
	DeregisterClusterInternal(ctx context.Context, params installer.DeregisterClusterParams) error
}
func NewGarbageCollectors(
	db *gorm.DB,
	log logrus.FieldLogger,
	hostApi host.API,
	clusterApi clusterPkg.API,
	objectHandler s3wrapper.API,
	leaderElector leader.Leader,
) *garbageCollector {
	return &garbageCollector{
		db:                  db,
		log:                 log,
		hostApi:             hostApi,
		clusterApi:          clusterApi,
		objectHandler:       objectHandler,
		leaderElector:       leaderElector,
	}
}

type garbageCollector struct {
	Config
	db                  *gorm.DB
	log                 logrus.FieldLogger
	hostApi             host.API
	clusterApi          clusterPkg.API
	objectHandler       s3wrapper.API
	leaderElector       leader.Leader
}

func (g garbageCollector) DeregisterInactiveClusters() {
	olderThan := strfmt.DateTime(time.Now().Add(-g.Config.DeregisterInactiveAfter))
	if err := g.clusterApi.InactiveClusterDeregister(context.Background(), olderThan, g.objectHandler); err != nil {
		g.log.WithError(err).Errorf("Failed deregister inactive clusters")
		return
	}
}

func (g garbageCollector) PermanentlyDeleteUnregisteredClustersAndHosts() {
	if !g.leaderElector.IsLeader() {
		g.log.Debugf("Not a leader, exiting periodic clusters and hosts deletion")
		return
	}

	olderThan := strfmt.DateTime(time.Now().Add(-g.Config.DeletedUnregisteredAfter))
	g.log.Debugf(
		"Permanently deleting all clusters that were de-registered before %s",
		olderThan)
	if err := g.clusterApi.PermanentClustersDeletion(context.Background(), olderThan, g.objectHandler); err != nil {
		g.log.WithError(err).Errorf("Failed deleting de-registered clusters")
		return
	}

	g.log.Debugf(
		"Permanently deleting all hosts that were soft-deleted before %s",
		olderThan)
	if err := g.hostApi.PermanentHostsDeletion(olderThan); err != nil {
		g.log.WithError(err).Errorf("Failed deleting soft-deleted hosts")
		return
	}
}