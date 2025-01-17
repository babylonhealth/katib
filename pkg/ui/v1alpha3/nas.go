package v1alpha3

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	trialsv1alpha3 "github.com/kubeflow/katib/pkg/apis/controller/trials/v1alpha3"
	api_pb_v1alpha3 "github.com/kubeflow/katib/pkg/apis/manager/v1alpha3"
)

func (k *KatibUIHandler) FetchAllNASJobs(w http.ResponseWriter, r *http.Request) {
	//enableCors(&w)
	// Use "" to get experiments in all namespaces.
	jobs, err := k.getExperimentList(k.katibClient.GetClientNamespace(), JobTypeNAS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(jobs)
	if err != nil {
		log.Printf("Marshal NAS jobs failed: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

func (k *KatibUIHandler) FetchNASJobInfo(w http.ResponseWriter, r *http.Request) {
	//enableCors(&w)
	experimentName := r.URL.Query()["experimentName"][0]
	namespace := r.URL.Query()["namespace"][0]

	responseRaw := make([]NNView, 0)
	var architecture string
	var decoder string

	conn, c := k.connectManager()

	defer conn.Close()

	trials, err := k.katibClient.GetTrialList(experimentName, namespace)
	if err != nil {
		log.Printf("GetTrialList from NAS job failed: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Got Trial List")

	for i, t := range trials.Items {
		succeeded := false
		for _, condition := range t.Status.Conditions {
			if condition.Type == trialsv1alpha3.TrialSucceeded {
				succeeded = true
			}
		}
		if succeeded {
			obsLogResp, err := c.GetObservationLog(
				context.Background(),
				&api_pb_v1alpha3.GetObservationLogRequest{
					TrialName: t.Name,
					StartTime: "",
					EndTime:   "",
				},
			)
			if err != nil {
				log.Printf("GetObservationLog from NAS job failed: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			metricsName := make([]string, 0)
			metricsValue := make([]string, 0)
			for _, m := range obsLogResp.ObservationLog.MetricLogs {
				metricsName = append(metricsName, m.Metric.Name)
				metricsValue = append(metricsValue, m.Metric.Value)

			}
			for _, trialParam := range t.Spec.ParameterAssignments {
				if trialParam.Name == "architecture" {
					architecture = trialParam.Value
				}
				if trialParam.Name == "nn_config" {
					decoder = trialParam.Value
				}
			}
			responseRaw = append(responseRaw, NNView{
				Name:         "Generation " + strconv.Itoa(i),
				TrialName:    t.Name,
				Architecture: generateNNImage(architecture, decoder),
				MetricsName:  metricsName,
				MetricsValue: metricsValue,
			})
		}
	}
	log.Printf("Logs parsed, result: %v", responseRaw)

	response, err := json.Marshal(responseRaw)
	if err != nil {
		log.Printf("Marshal result in NAS job failed: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}
