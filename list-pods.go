package main

import (
        "k8s.io/client-go/kubernetes"
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        "k8s.io/client-go/tools/clientcmd"
        "log"
        "fmt"
        "os"
        "flag"
        "path/filepath"
)
var err error
func check(err error){
        if err != nil {
                log.Fatal(err)
        }   
     }
func main() {
        // Access to the Cluster using the kubeconfig file 
        kubeconfig := flag.String("kubeconfig", filepath.Join(os.Getenv("HOME"), ".kube","config"),"/home/aly/.kube/config")
        flag.Parse()
        
        config, err :=clientcmd.BuildConfigFromFlags("", *kubeconfig)
        check(err)
        
        clientset, err := kubernetes.NewForConfig(config)
        check(err)
        
        // List the available pods in the cluster using K8s API
        pods , err := clientset.CoreV1().Pods("") .List(metav1.ListOptions{})
        check(err)
        
        // Loop to all Pods and print them out 
        fmt.Printf("There are %d Pods available in the Cluster \n", len(pods.Items))
        for _, pod := range pods.Items {
                fmt.Println(pod.GetName())
        }
}
