package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	ui "github.com/kubeflow/katib/pkg/ui/v1alpha3"
)

type namespaceList []string

func (n *namespaceList) String() string {
	return fmt.Sprintf("%v", *n)
}

func (n *namespaceList) Set(value string) error {
	*n = append(*n, value)
	return nil
}

var (
	port, host, buildDir *string
)

var namespaces namespaceList

func init() {
	port = flag.String("port", "80", "the port to listen to for incoming HTTP connections")
	host = flag.String("host", "0.0.0.0", "the host to listen to for incoming HTTP connections")
	flag.Var(&namespaces, "namespace", "the namespace list where the UI operates")
	buildDir = flag.String("build-dir", "/app/build", "the dir of frontend")
}
func main() {
	flag.Parse()
	log.Printf("Creating Katib UI Handler in namespace %s", namespaces)
	kuh := ui.NewKatibUIHandler(namespaces)

	log.Printf("Serving the frontend dir %s", *buildDir)
	frontend := http.FileServer(http.Dir(*buildDir))
	http.Handle("/katib/", http.StripPrefix("/katib/", frontend))

	http.HandleFunc("/katib/fetch_hp_jobs/", kuh.FetchAllHPJobs)
	http.HandleFunc("/katib/fetch_nas_jobs/", kuh.FetchAllNASJobs)
	http.HandleFunc("/katib/submit_yaml/", kuh.SubmitYamlJob)
	http.HandleFunc("/katib/submit_hp_job/", kuh.SubmitParamsJob)
	http.HandleFunc("/katib/submit_nas_job/", kuh.SubmitParamsJob)

	http.HandleFunc("/katib/delete_experiment/", kuh.DeleteExperiment)

	http.HandleFunc("/katib/fetch_hp_job/", kuh.FetchHPJob)
	http.HandleFunc("/katib/fetch_hp_job_info/", kuh.FetchHPJobInfo)
	http.HandleFunc("/katib/fetch_hp_job_trial_info/", kuh.FetchHPJobTrialInfo)
	http.HandleFunc("/katib/fetch_nas_job_info/", kuh.FetchNASJobInfo)

	http.HandleFunc("/katib/fetch_trial_templates/", kuh.FetchTrialTemplates)
	http.HandleFunc("/katib/update_template/", kuh.AddEditDeleteTemplate)

	http.HandleFunc("/katib/fetch_namespaces", kuh.FetchNamespaces)

	log.Printf("Serving at %s:%s", *host, *port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", *host, *port), nil); err != nil {
		panic(err)
	}
}
