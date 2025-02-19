package dynamic

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/weibaohui/k8m/pkg/comm/utils/amis"
	"github.com/weibaohui/k8m/pkg/service"
	"github.com/weibaohui/kom/kom"
	v1 "k8s.io/api/core/v1"
)

var linkCacheTTL = 5 * time.Minute

func getPod(selectedCluster string, ctx context.Context, ns string, name string, kind string) (*v1.Pod, error) {
	var pod *v1.Pod
	var err error
	kk := kom.Cluster(selectedCluster).WithContext(ctx).
		Resource(&v1.Pod{}).
		Namespace(ns).
		Name(name).
		WithCache(linkCacheTTL)
	switch kind {
	case "Pod":
		err = kk.Get(&pod).Error
	case "Deployment":
		pod, err = kk.Ctl().Deployment().ManagedPod()
	case "StatefulSet":
		pod, err = kk.Ctl().StatefulSet().ManagedPod()
	case "DaemonSet":
		pod, err = kk.Ctl().DaemonSet().ManagedPod()
	case "ReplicaSet":
		pod, err = kk.Ctl().ReplicaSet().ManagedPod()
	}
	return pod, err
}
func getPods(selectedCluster string, ctx context.Context, ns string, name string, kind string) ([]*v1.Pod, error) {
	var pods []*v1.Pod
	var err error
	kk := kom.Cluster(selectedCluster).WithContext(ctx).
		Resource(&v1.Pod{}).
		Namespace(ns).
		Name(name).
		WithCache(linkCacheTTL)
	switch kind {
	case "Deployment":
		pods, err = kk.Ctl().Deployment().ManagedPods()
	case "StatefulSet":
		pods, err = kk.Ctl().StatefulSet().ManagedPods()
	case "DaemonSet":
		pods, err = kk.Ctl().DaemonSet().ManagedPods()
	case "ReplicaSet":
		pods, err = kk.Ctl().ReplicaSet().ManagedPods()
	}
	return pods, err
}
func LinksServices(c *gin.Context) {
	name := c.Param("name")
	ns := c.Param("ns")
	ctx := c.Request.Context()
	kind := c.Param("kind")
	selectedCluster := amis.GetSelectedCluster(c)

	pod, err := getPod(selectedCluster, ctx, ns, name, kind)

	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	services, err := service.PodService().LinksServices(ctx, selectedCluster, pod)
	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	amis.WriteJsonList(c, services)
}

func LinksEndpoints(c *gin.Context) {
	name := c.Param("name")
	ns := c.Param("ns")
	ctx := c.Request.Context()
	kind := c.Param("kind")
	selectedCluster := amis.GetSelectedCluster(c)

	pod, err := getPod(selectedCluster, ctx, ns, name, kind)

	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}

	endpoints, err := service.PodService().LinksEndpoints(ctx, selectedCluster, pod)
	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	amis.WriteJsonList(c, endpoints)

}

func LinksPVC(c *gin.Context) {
	name := c.Param("name")
	ns := c.Param("ns")
	ctx := c.Request.Context()
	kind := c.Param("kind")
	selectedCluster := amis.GetSelectedCluster(c)

	pod, err := getPod(selectedCluster, ctx, ns, name, kind)

	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	pvc, err := service.PodService().LinksPVC(ctx, selectedCluster, pod)
	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	amis.WriteJsonList(c, pvc)
}

func LinksPV(c *gin.Context) {
	name := c.Param("name")
	ns := c.Param("ns")
	ctx := c.Request.Context()
	kind := c.Param("kind")
	selectedCluster := amis.GetSelectedCluster(c)

	pod, err := getPod(selectedCluster, ctx, ns, name, kind)

	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	pv, err := service.PodService().LinksPV(ctx, selectedCluster, pod)
	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	amis.WriteJsonList(c, pv)
}

func LinksIngress(c *gin.Context) {
	name := c.Param("name")
	ns := c.Param("ns")
	ctx := c.Request.Context()
	kind := c.Param("kind")
	selectedCluster := amis.GetSelectedCluster(c)

	pod, err := getPod(selectedCluster, ctx, ns, name, kind)

	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	ingress, err := service.PodService().LinksIngress(ctx, selectedCluster, pod)
	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	amis.WriteJsonList(c, ingress)
}

func LinksEnv(c *gin.Context) {
	name := c.Param("name")
	ns := c.Param("ns")
	ctx := c.Request.Context()
	kind := c.Param("kind")
	selectedCluster := amis.GetSelectedCluster(c)

	pod, err := getPod(selectedCluster, ctx, ns, name, kind)

	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	env, err := service.PodService().LinksEnv(ctx, selectedCluster, pod)
	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	amis.WriteJsonList(c, env)
}

func LinksEnvFromPod(c *gin.Context) {
	name := c.Param("name")
	ns := c.Param("ns")
	ctx := c.Request.Context()
	kind := c.Param("kind")
	selectedCluster := amis.GetSelectedCluster(c)

	pod, err := getPod(selectedCluster, ctx, ns, name, kind)

	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	env, err := service.PodService().LinksEnvFromPod(ctx, selectedCluster, pod)
	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	amis.WriteJsonList(c, env)
}

func LinksConfigMap(c *gin.Context) {
	name := c.Param("name")
	ns := c.Param("ns")
	ctx := c.Request.Context()
	kind := c.Param("kind")
	selectedCluster := amis.GetSelectedCluster(c)

	pod, err := getPod(selectedCluster, ctx, ns, name, kind)

	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	configMap, err := service.PodService().LinksConfigMap(ctx, selectedCluster, pod)
	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	amis.WriteJsonList(c, configMap)
}

func LinksSecret(c *gin.Context) {
	name := c.Param("name")
	ns := c.Param("ns")
	ctx := c.Request.Context()
	kind := c.Param("kind")
	selectedCluster := amis.GetSelectedCluster(c)

	pod, err := getPod(selectedCluster, ctx, ns, name, kind)

	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	secret, err := service.PodService().LinksSecret(ctx, selectedCluster, pod)
	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	amis.WriteJsonList(c, secret)
}

func LinksNode(c *gin.Context) {
	name := c.Param("name")
	ns := c.Param("ns")
	ctx := c.Request.Context()
	kind := c.Param("kind")
	selectedCluster := amis.GetSelectedCluster(c)

	pod, err := getPod(selectedCluster, ctx, ns, name, kind)

	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	nodes, err := service.PodService().LinksNode(ctx, selectedCluster, pod)
	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	amis.WriteJsonList(c, nodes)
}
func LinksPod(c *gin.Context) {
	name := c.Param("name")
	ns := c.Param("ns")
	ctx := c.Request.Context()
	kind := c.Param("kind")
	selectedCluster := amis.GetSelectedCluster(c)
	pods, err := getPods(selectedCluster, ctx, ns, name, kind)
	if err != nil {
		amis.WriteJsonError(c, err)
		return
	}
	amis.WriteJsonList(c, pods)
}
