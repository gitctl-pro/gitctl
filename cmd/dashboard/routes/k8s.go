package routes

import (
	"github.com/gitctl-pro/gitctl/controller/k8s"
)

func (r *RouteManager) addK8sRoutes(path string) {
	k8s := k8s.NewController()
	// route: /api/v1/cluster
	rg := r.gin.Group(path + "/cluster")
	rg.Use()
	{
		rg.GET("", k8s.Cluster.ListCluster)
		rg.GET("/:name", k8s.Cluster.GetCluster)
	}

	// route: /api/v1/node
	rg = r.gin.Group(path + "/node")
	rg.Use()
	{
		rg.GET("", k8s.Node.ListNode)
		rg.GET("/:name", k8s.Node.GetNode)
	}

	// route: /api/v1/namespace
	rg = r.gin.Group(path + "/namespace")
	rg.Use()
	{
		rg.GET("", k8s.Namespace.ListNamespace)
		rg.POST("/create", k8s.Namespace.CreateNamespace)
		rg.POST("/:name", k8s.Namespace.UpdateNamespace)
		rg.GET("/:name", k8s.Namespace.ListNamespace)
		rg.DELETE("/:name", k8s.Namespace.DeleteNamespace)
	}

	// route: /api/v1/application
	rg = r.gin.Group(path + "/application")
	rg.Use()
	{
		rg.GET("", k8s.Application.ListApplication)
		rg.POST("/create", k8s.Application.CreateApplication)
		rg.POST("/:name", k8s.Application.UpdateApplication)
		rg.DELETE("/:name", k8s.Application.DeleteApplication)
	}

	// route: /api/v1/deployment
	rg = r.gin.Group(path + "/deployment")
	rg.Use()
	{
		rg.GET("/:namespace", k8s.Deployment.List)
		rg.GET("/:namespace/:name", k8s.Deployment.Get)

		rg.GET("/:namespace/:name/events", k8s.Deployment.Events)
		rg.GET("/:namespace/:name/repliaces", k8s.Deployment.ReplicaSets)
		rg.GET("/:namespace/:name/repliaces/new", k8s.Deployment.NewReplicaSets)
		rg.GET("/:namespace/:name/repliaces/old", k8s.Deployment.OldReplicaSets)

		//rollout
		rg.GET("/:namespace/:name/restart", k8s.Deployment.RolloutRestart)
		rg.GET("/:namespace/:name/pause", k8s.Deployment.RolloutPause)
		rg.GET("/:namespace/:name/resume", k8s.Deployment.RolloutResume)
		rg.GET("/:namespace/:name/rollback", k8s.Deployment.RolloutRollback)
	}

	// route: /api/v1/replicaset
	rg = r.gin.Group(path + "/replicaset")
	rg.Use()
	{
		rg.GET("/list", k8s.ReplicaSet.List)
		rg.GET("/:namespace/:name", k8s.ReplicaSet.Get)
		rg.GET("/:namespace/:name/pod", k8s.ReplicaSet.Pods)
		rg.GET("/:namespace/:name/service", k8s.ReplicaSet.Service)
		rg.GET("/:namespace/:name/event", k8s.ReplicaSet.Events)
		rg.DELETE("/:namespace/:name", k8s.ReplicaSet.Delete)
	}

	// route: /api/v1/pod
	rg = r.gin.Group(path + "/pod")
	rg.Use()
	{
		rg.GET("/list", k8s.Pod.List)
		rg.GET("/namespace/:name", k8s.Pod.Get)
		rg.GET("/:namespace/:name/event", k8s.Pod.Events)
		rg.GET("/:namespace/:name/container", k8s.Pod.Containers)
		rg.GET("/:namespace/:name/pvc", k8s.Pod.PersistentVolumeClaims)
		rg.GET("/:namespace/:name/webshell/info", k8s.Pod.ExecShellInfo)
		rg.GET("/:namespace/:name/webshell", k8s.Pod.ExecShell)
	}

	// route: /api/v1/log {
	rg = r.gin.Group(path + "/log")
	rg.Use()
	{
		rg.GET("/log/:namespace/:name", k8s.Pod.LogDetail)
		//rg.GET("/log/:namespace/:name/file",)
	}

	// route: /api/v1/ingress
	rg = r.gin.Group(path + "/ingress")
	rg.Use()
	{
		rg.GET("/list", k8s.Ingress.ListIngress)
		rg.GET("/:namespace/:name", k8s.Ingress.GetIngress)
		rg.GET("/:namespace/:name/events", k8s.Ingress.Events)
	}

	// route: /api/v1/job
	rg = r.gin.Group(path + "/job")
	rg.Use()
	{
		rg.GET("", k8s.Job.ListJob)
		rg.GET("/:namespace/:name", k8s.Job.GetJob)
		rg.GET("/:namespace/:name/events", k8s.Job.Events)
	}

	// route: /api/v1/cronjob
	rg = r.gin.Group(path + "/cronjob")
	rg.Use()
	{
		rg.GET("", k8s.CronJob.ListCronJob)
		rg.GET("/:namespace/:name", k8s.CronJob.GetCronJob)
		rg.GET("/:namespace/:name/events", k8s.CronJob.Events)

	}
	// route: /api/v1/secret
	rg = r.gin.Group(path + "/secret")
	rg.Use()
	{
		rg.GET("", k8s.Secret.ListSecret)
		rg.GET("/:namespace/:name", k8s.Secret.GetSecret)
	}

	// route: /api/v1/pvc
	rg = r.gin.Group(path + "/pvc")
	rg.Use()
	{
		rg.GET("", k8s.PVC.ListPVC)
		rg.GET("/:namespace/:name", k8s.PVC.GetPVC)
	}

	// route: /api/v1/event
	rg = r.gin.Group(path + "/event")
	rg.Use()
	{
		rg.GET("", k8s.Event.ListEvents)
	}

	// route: /api/v1/hpa
	rg = r.gin.Group(path + "/hpa")
	rg.Use()
	{
		rg.GET("", k8s.HPA.ListHPA)
		rg.GET("/:namespace/:name", k8s.HPA.GetHPA)
		rg.POST("/:namespace/:name", k8s.HPA.UpdateHPA)
		rg.POST("/:namespace/:name", k8s.HPA.CreateHPA)
		rg.DELETE("/:namespace/:name", k8s.HPA.DeleteHPA)
	}

	// route: /api/v1/service
	rg = r.gin.Group(path + "/service")
	rg.Use()
	{
		rg.GET("", k8s.Service.ListService)
		rg.GET("/:namespace/:name", k8s.Service.GetService)
		rg.GET("/:namespace/:name/pod", k8s.Service.Pods)
		rg.GET("/:namespace/:name/event", k8s.Service.Events)

	}

	// route: /api/v1/configmap
	rg = r.gin.Group(path + "/configmap")
	rg.Use()
	{
		rg.GET("/list", k8s.Configmap.List)
		rg.GET("/:namespace/:name", k8s.Configmap.Get)
		rg.GET("/:namespace/:name", k8s.Configmap.Create)
		rg.GET("/:namespace/:name", k8s.Configmap.Update)
		rg.GET("/:namespace/:name/event", k8s.Configmap.Events)
	}

	// route: /api/v1/scale
	rg = r.gin.Group(path + "/scale")
	rg.Use()
	{
		rg.POST("/:kind/:namespace/:name", k8s.Scale.ScaleResource)
	}
}
