package events

import (
	"context"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/openshift/assisted-service/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/openshift/assisted-service/internal/common"
	logutil "github.com/openshift/assisted-service/pkg/log"
	"github.com/openshift/assisted-service/restapi"
	"github.com/openshift/assisted-service/restapi/operations/events"
	"github.com/sirupsen/logrus"
)

var _ restapi.EventsAPI = &Api{}

type Api struct {
	db      *gorm.DB
	handler Handler
	log     logrus.FieldLogger
}

func NewApi(db *gorm.DB, handler Handler, log logrus.FieldLogger) *Api {
	return &Api{
		db:      db,
		handler: handler,
		log:     log,
	}
}

func (a *Api) ListEvents(ctx context.Context, params events.ListEventsParams) middleware.Responder {
	log := logutil.FromContext(ctx, a.log)
	var hostID string

	if params.HostID != nil {
		hostID = (*params.HostID).String()
		dbReply := a.db.Model(&models.Host{}).Where("id = ?", hostID)
		if dbReply.Error != nil {
			return common.NewApiError(http.StatusNotFound, dbReply.Error)
		}
	} else {
		dbReply := a.db.Model(&models.Cluster{}).Where("id = ?", params.ClusterID.String())
		if dbReply.Error != nil {
			return common.NewApiError(http.StatusNotFound, dbReply.Error)
		}
		hostID = "<none>"
	}
	evs, err := a.handler.GetEvents(params.ClusterID, params.HostID)
	if err != nil {
		log.Errorf("failed to get events for cluster %s host %s", params.ClusterID.String(), hostID)
		return common.NewApiError(http.StatusInternalServerError, err)
	}
	ret := make(models.EventList, len(evs))
	for i, ev := range evs {
		ret[i] = &models.Event{
			ClusterID: ev.ClusterID,
			HostID:    ev.HostID,
			Severity:  ev.Severity,
			EventTime: ev.EventTime,
			Message:   ev.Message,
		}
	}
	return events.NewListEventsOK().WithPayload(ret)

}
