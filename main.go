package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func getClientSet() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func listNamespaces(c *gin.Context) {
	clientset, err := getClientSet()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	nsList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	names := []string{}
	for _, ns := range nsList.Items {
		names = append(names, ns.Name)
	}
	c.JSON(http.StatusOK, gin.H{"namespaces": names})
}

func main() {
	r := gin.Default()
	r.Static("/", "./frontend/dist")
	r.GET("/api/namespaces", listNamespaces)
	r.Run(":8080")
}
