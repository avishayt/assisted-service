package events

import (
	"context"

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
	handler Handler
	log     logrus.FieldLogger
}

func NewApi(handler Handler, log logrus.FieldLogger) *Api {
	return &Api{
		handler: handler,
		log:     log,
	}
}

func (a *Api) ListEvents(ctx context.Context, params events.ListEventsParams) middleware.Responder {
	log := logutil.FromContext(ctx, a.log)
	var hostID string
	if params.HostID != nil {
		hostID = (*params.HostID).String()
	} else {
		hostID = "<none>"
	}
	evs, err := a.handler.GetEvents(params.ClusterID, params.HostID)
	if err != nil {
		log.Errorf("failed to get events for cluster %s host %s", params.ClusterID.String(), hostID)
		return events.NewListEventsInternalServerError().
			WithPayload(common.GenerateInternalFromError(err))
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
