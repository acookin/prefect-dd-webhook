package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

type FlowStateChangeEvent struct {
	TenantSlug  string `json:"tenant_slug"`
	FlowName    string `json:"flow_name"`
	FlowRunLink string `json:"flow_run_link"`
	State       string `json:"state"`
	FlowRunId   string `json:"flow_run_id"`
}

func sendDataDogEvent(event FlowStateChangeEvent) error {
	body := datadog.EventCreateRequest{
		Title: fmt.Sprintf("prefect_flow_%s_%s", event.FlowName, event.State),
		Text:  fmt.Sprintf("Flow run %s has %s. Link: %s", event.FlowName, event.State, event.FlowRunLink),
		Tags: []string{
			fmt.Sprintf("env:%s", event.TenantSlug),
			fmt.Sprintf("flow_run_id:%s", event.FlowRunId),
			fmt.Sprintf("flow_name:%s", event.FlowName),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	_, _, err := apiClient.EventsApi.CreateEvent(ctx, body)

	if err != nil {
		return err
	}
	return nil
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	var event FlowStateChangeEvent
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sendDataDogEvent(event)
}

func main() {
	log.Println("server started")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
