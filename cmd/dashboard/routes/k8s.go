package routes

import (
	"github.com/gitctl-pro/gitctl/controller/k8s"
)

func (r *RouteManager) addK8sRoutes(path string) {
	k8s := k8s.NewController(nil, nil)

	// route: /k8s/node
	rg := r.gin.Group(path + "/node")
	rg.Use()
	{
		rg.GET("", k8s.Node.ListNode)
		rg.GET("/:name", k8s.Node.Get)
	}

	// route: /k8s/namespace
	rg = r.gin.Group(path + "/namespace")
	rg.Use()
	{
		rg.GET("", k8s.Namespace.List)
		rg.POST("/:name", k8s.Namespace.Create)
		rg.PUT("/:name", k8s.Namespace.Put)
		rg.GET("/:name", k8s.Namespace.List)
		rg.DELETE("/:name", k8s.Namespace.Delete)
	}

	// route: /k8s/deployment
	rg = r.gin.Group(path + "/deployment")
	rg.Use()
	{
		rg.GET("/:namespace", k8s.Deployment.List)
		rg.GET("/:namespace/:name", k8s.Deployment.Get)
		rg.PUT("/:namespace/:name", k8s.Deployment.Put)

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

	// route: /k8s/replicaset
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

	// route: /k8s/pod
	rg = r.gin.Group(path + "/pod")
	rg.Use()
	{
		rg.GET("", k8s.Pod.ListPod)
		rg.GET("/namespace/:name", k8s.Pod.Get)
		rg.GET("/:namespace/:name/event", k8s.Pod.Events)
		rg.GET("/:namespace/:name/container", k8s.Pod.Containers)
		rg.GET("/:namespace/:name/pvc", k8s.Pod.PersistentVolumeClaims)
		rg.GET("/:namespace/:name/webshell/info", k8s.Pod.ExecShellInfo)
		rg.GET("/:namespace/:name/webshell", k8s.Pod.ExecShell)
	}

	// route: /k8s/log {
	rg = r.gin.Group(path + "/log")
	rg.Use()
	{
		rg.GET("/log/:namespace/:name", k8s.Pod.GetLogs)
		//rg.GET("/log/:namespace/:name/file",)
	}

	// route: /k8s/ingress
	rg = r.gin.Group(path + "/ingress")
	rg.Use()
	{
		rg.GET("", k8s.Ingress.ListIngress)
		rg.GET("/:namespace/:name", k8s.Ingress.Get)
		rg.POST("/:namespace/:name", k8s.Ingress.Create)
		rg.PUT("/:namespace/:name", k8s.Ingress.Put)
		rg.PUT("/:namespace/:name/update", k8s.Ingress.Put)
		rg.GET("/:namespace/:name/events", k8s.Ingress.Events)
		rg.DELETE("/:namespace/:name", k8s.Ingress.Delete)
	}

	// route: /k8s/job
	rg = r.gin.Group(path + "/job")
	rg.Use()
	{
		rg.GET("", k8s.Job.ListJob)
		rg.GET("/:namespace/:name", k8s.Job.Get)
		rg.POST("/:namespace/:name", k8s.Job.Create)
		rg.PUT("/:namespace/:name", k8s.Job.Put)
		rg.DELETE("/:namespace/:name", k8s.Job.Delete)
		rg.GET("/:namespace/:name/events", k8s.Job.Events)
	}

	// route: /k8s/cronjob
	rg = r.gin.Group(path + "/cronjob")
	rg.Use()
	{
		rg.GET("", k8s.CronJob.ListCronjob)
		rg.GET("/:namespace/:name", k8s.CronJob.Get)
		rg.PUT("/:namespace/:name/update", k8s.CronJob.Put)
		rg.POST("/:namespace/:name", k8s.CronJob.Create)
		rg.DELETE("/:namespace/:name", k8s.Job.Delete)
		rg.GET("/:namespace/:name/events", k8s.CronJob.Events)
	}
	// route: /k8s/secret
	rg = r.gin.Group(path + "/secret")
	rg.Use()
	{
		rg.GET("", k8s.Secret.ListSecret)
		rg.GET("/:namespace/:name", k8s.Secret.Get)
		rg.POST("/:namespace/:name", k8s.Secret.Create)
		rg.PUT("/:namespace/:name", k8s.Secret.Put)
		rg.DELETE("/:namespace/:name", k8s.Secret.Delete)
	}

	// route: /k8s/pvc
	rg = r.gin.Group(path + "/pvc")
	rg.Use()
	{
		rg.GET("", k8s.PVC.ListPVC)
		rg.GET("/:namespace/:name", k8s.PVC.Get)
		rg.POST("/:namespace/:name", k8s.PVC.Create)
		rg.DELETE("/:namespace/:name", k8s.PVC.Delete)
	}

	// route: /k8s/pv
	rg = r.gin.Group(path + "/pv")
	rg.Use()
	{
		rg.GET("", k8s.PV.ListPV)
		rg.GET("/:name", k8s.PV.Get)
		rg.DELETE("/:name", k8s.PV.Delete)
		rg.POST("/:name", k8s.PV.Create)
	}

	// route: /k8s/event
	rg = r.gin.Group(path + "/event")
	rg.Use()
	{
		rg.GET("", k8s.Event.ListEvents)
	}

	// route: /k8s/hpa
	rg = r.gin.Group(path + "/hpa")
	rg.Use()
	{
		rg.GET("", k8s.HPA.ListHPA)
		rg.GET("/:namespace/:name", k8s.HPA.Get)
		rg.POST("/:namespace/:name", k8s.HPA.Create)
		rg.PUT("/:namespace/:name", k8s.HPA.Put)
		rg.DELETE("/:namespace/:name", k8s.HPA.Delete)
	}

	// route: /k8s/service
	rg = r.gin.Group(path + "/service")
	rg.Use()
	{
		rg.GET("", k8s.Service.ListService)
		rg.GET("/:namespace/:name", k8s.Service.Get)
		rg.POST("/:namespace/:name", k8s.Service.Create)
		rg.PUT("/:namespace/:name", k8s.Service.Put)
		rg.DELETE("/:namespace/:name", k8s.Service.Delete)
		rg.GET("/:namespace/:name/pod", k8s.Service.Pods)
		rg.GET("/:namespace/:name/event", k8s.Service.Events)

	}

	// route: /k8s/configmap
	rg = r.gin.Group(path + "/configmap")
	rg.Use()
	{
		rg.GET("", k8s.ConfigMap.ListConfigmap)
		rg.GET("/:namespace/:name", k8s.ConfigMap.Get)
		rg.POST("/:namespace/:name", k8s.ConfigMap.Create)
		rg.PUT("/:namespace/:name", k8s.ConfigMap.Put)
		rg.DELETE("/:namespace/:name", k8s.ConfigMap.Delete)
	}

	// route: /k8s/scale
	rg = r.gin.Group(path + "/scale")
	rg.Use()
	{
		rg.POST("/:kind/:namespace/:name", k8s.Scale.ScaleResource)
	}

	// route: /k8s/serviceAccount
	rg = r.gin.Group(path + "/serviceaccount")
	rg.Use()
	{
		rg.GET("", k8s.ServiceAccount.ListServiceAccount)
		rg.GET("/:name", k8s.ServiceAccount.Get)
		rg.DELETE("/:name", k8s.ServiceAccount.Delete)
		rg.POST("/:namespace/:name", k8s.ServiceAccount.Create)
		rg.PUT("/:namespace/:name", k8s.ServiceAccount.Put)
	}

	// route: /k8s/clusterRole
	rg = r.gin.Group(path + "/clusterRole")
	rg.Use()
	{
		rg.GET("", k8s.ClusterRole.ListClusterRole)
		rg.GET("/:name", k8s.ClusterRole.Get)
		rg.DELETE("/:name", k8s.ClusterRole.Delete)
		rg.POST("/:name", k8s.ClusterRole.Create)
		rg.PUT("/:name", k8s.ClusterRole.Put)
	}

	// route: /k8s/role
	rg = r.gin.Group(path + "/clusterRole")
	rg.Use()
	{
		rg.GET("", k8s.Role.ListRole)
		rg.GET("/:namespace/:name", k8s.Role.Get)
		rg.DELETE("/:namespace/:name", k8s.Role.Delete)
		rg.POST("/:namespace/:name", k8s.Role.Create)
		rg.PUT("/:namespace/:name", k8s.Role.Put)
	}
}
