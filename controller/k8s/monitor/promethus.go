package monitor

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"github.com/gitctl-pro/gitctl/pkg/monitor/promethus"
	"net/http"
	"net/url"
)

type monitor struct {
	clusterManager k8s.ClusterManager
	promethus      *promethus.Client
}

func NewMonitor(clusterManager k8s.ClusterManager) *monitor {
	client := promethus.NewPrometheusClient("promethus:8081")
	return &monitor{clusterManager: clusterManager, promethus: client}
}

func (ctl *monitor) PodSeries(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")
	queryMap := make(map[string]string, 0)
	ctx.BindQuery(&queryMap)
	queryMap["cluster"] = cluster
	queryMap["namespace"] = namespace

	exp := promethus.PodSeriesExp(name, queryMap)
	query := url.Values{}
	query.Set("match[]", exp)
	query.Set("start", ctx.Query("start"))
	query.Set("end", ctx.Query("end"))

	response, _ := ctl.promethus.Series(query)
	ctx.Data(http.StatusOK, "application/json", response)
}

func (ctl *monitor) PodQueryRange(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	namespace := ctx.Param("namespace")
	pod := ctx.Param("pod")
	queryMap := make(map[string]string, 0)
	ctx.BindQuery(&queryMap)
	queryMap["cluster"] = cluster
	queryMap["namespace"] = namespace
	queryMap["pod"] = pod
	metric := ctx.Query("metric")
	exp := promethus.PodQueryExp(metric, queryMap)

	query := url.Values{}
	query.Set("query", exp)
	query.Set("start", ctx.Query("start"))
	query.Set("end", ctx.Query("end"))
	query.Set("step", ctx.DefaultQuery("step", "5"))

	response, _ := ctl.promethus.QueryRange(query)
	ctx.Data(http.StatusOK, "application/json", response)
}
