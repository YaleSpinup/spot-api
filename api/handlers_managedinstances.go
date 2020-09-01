package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/YaleSpinup/spot-api/apierror"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spotinst/spotinst-sdk-go/service/managedinstance/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
)

// ManagedInstanceListHandler handles listing managed instances in SpotInst
func (s *server) ManagedInstanceListHandler(w http.ResponseWriter, r *http.Request) {
	w = LogWriter{w}
	vars := mux.Vars(r)
	account := vars["account"]
	miService, ok := s.managedinstanceServices[account]
	if !ok {
		log.Errorf("account not found: %s", account)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	output, err := miService.ListAWSManagedInstances(r.Context())
	if err != nil {
		handleError(w, err)
		return
	}

	instances := []string{}
	for _, i := range output {
		instances = append(instances, spotinst.StringValue(i.ID))
	}

	j, err := json.Marshal(instances)
	if err != nil {
		log.Errorf("cannot marshal response (%v) into JSON: %s", instances, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// ManagedInstanceShowHandler handles getting details about a managed instance from SpotInst
func (s *server) ManagedInstanceShowHandler(w http.ResponseWriter, r *http.Request) {
	w = LogWriter{w}
	vars := mux.Vars(r)
	account := vars["account"]
	miService, ok := s.managedinstanceServices[account]
	if !ok {
		log.Errorf("account not found: %s", account)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	managedinstance := vars["instance"]

	output, err := miService.GetAWSManagedInstanceByID(r.Context(), managedinstance)
	if err != nil {
		handleError(w, err)
		return
	}

	j, err := json.Marshal(output)
	if err != nil {
		log.Errorf("cannot marshal response (%v) into JSON: %s", output, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// ManagedInstanceCreateHandler handles creating a managed instance in SpotInst
func (s *server) ManagedInstanceCreateHandler(w http.ResponseWriter, r *http.Request) {
	w = LogWriter{w}
	vars := mux.Vars(r)
	account := vars["account"]
	miService, ok := s.managedinstanceServices[account]
	if !ok {
		log.Errorf("account not found: %s", account)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	req := aws.ManagedInstance{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		msg := fmt.Sprintf("cannot decode body into create elastigroup input: %s", err)
		handleError(w, apierror.New(apierror.ErrBadRequest, msg, err))
		return
	}

	output, err := miService.CreateAWSManagedInstance(r.Context(), &req)
	if err != nil {
		handleError(w, err)
		return
	}

	j, err := json.Marshal(output)
	if err != nil {
		log.Errorf("cannot marshal response (%v) into JSON: %s", output, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// ManagedInstanceUpdateHandler handles updating a managed instance in SpotInst
func (s *server) ManagedInstanceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	w = LogWriter{w}
	vars := mux.Vars(r)
	account := vars["account"]
	miService, ok := s.managedinstanceServices[account]
	if !ok {
		log.Errorf("account not found: %s", account)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	managedinstance := vars["instance"]

	req := aws.ManagedInstance{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		msg := fmt.Sprintf("cannot decode body into update managed instance input: %s", err)
		handleError(w, apierror.New(apierror.ErrBadRequest, msg, err))
		return
	}

	// assert {managedinstance} from querystring route on req object
	req.ID = spotinst.String(managedinstance)

	output, err := miService.UpdateAWSManagedInstance(r.Context(), &req)
	if err != nil {
		msg := fmt.Sprintf("%s", err)
		handleError(w, apierror.New(apierror.ErrBadRequest, msg, err))
		return
	}

	j, err := json.Marshal(output)
	if err != nil {
		handleError(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// ManagedInstanceDeleteHandler handles deleting a managed instance from SpotInst
func (s *server) ManagedInstanceDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w = LogWriter{w}
	vars := mux.Vars(r)
	account := vars["account"]
	miService, ok := s.managedinstanceServices[account]
	if !ok {
		log.Errorf("account not found: %s", account)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	managedinstance := vars["instance"]

	err := miService.DeleteAWSManagedInstanceByID(r.Context(), managedinstance)
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
