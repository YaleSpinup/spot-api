package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/YaleSpinup/spot-api/apierror"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
)

// ElastigroupsListHandler handles listing elastigroups in SpotInst
func (s *server) ElastigroupsListHandler(w http.ResponseWriter, r *http.Request) {
	w = LogWriter{w}
	vars := mux.Vars(r)
	account := vars["account"]
	esService, ok := s.elastigroupServices[account]
	if !ok {
		log.Errorf("account not found: %s", account)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	output, err := esService.ListAWSElastigroups(r.Context())
	if err != nil {
		handleError(w, err)
		return
	}

	groups := []string{}
	for _, g := range output {
		groups = append(groups, spotinst.StringValue(g.ID))
	}

	j, err := json.Marshal(groups)
	if err != nil {
		log.Errorf("cannot marshal response (%v) into JSON: %s", groups, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// ElastigroupShowHandler handles getting details about an elastigroup from SpotInst
func (s *server) ElastigroupShowHandler(w http.ResponseWriter, r *http.Request) {
	w = LogWriter{w}
	vars := mux.Vars(r)
	account := vars["account"]
	esService, ok := s.elastigroupServices[account]
	if !ok {
		log.Errorf("account not found: %s", account)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	elastigroup := vars["elastigroup"]

	output, err := esService.GetAWSElastigroupByID(r.Context(), elastigroup)
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

// ElastigroupUpdateHandler handles updating an elastigroup in SpotInst
func (s *server) ElastigroupUpdateHandler(w http.ResponseWriter, r *http.Request) {
	w = LogWriter{w}
	vars := mux.Vars(r)
	account := vars["account"]
	esService, ok := s.elastigroupServices[account]
	if !ok {
		log.Errorf("account not found: %s", account)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	elastigroup := vars["elastigroup"]

	req := aws.Group{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		msg := fmt.Sprintf("cannot decode body into update elastigroup input: %s", err)
		handleError(w, apierror.New(apierror.ErrBadRequest, msg, err))
		return
	}

	req.ID = spotinst.String(elastigroup)

	output, err := esService.UpdateAWSElastigroup(r.Context(), &req)
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

// ElastigroupCreateHandler handles creating an elastigroup in SpotInst
func (s *server) ElastigroupCreateHandler(w http.ResponseWriter, r *http.Request) {
	w = LogWriter{w}
	vars := mux.Vars(r)
	account := vars["account"]
	esService, ok := s.elastigroupServices[account]
	if !ok {
		log.Errorf("account not found: %s", account)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	req := aws.Group{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		msg := fmt.Sprintf("cannot decode body into create elastigroup input: %s", err)
		handleError(w, apierror.New(apierror.ErrBadRequest, msg, err))
		return
	}

	output, err := esService.CreateAWSElastigroup(r.Context(), &req)
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

// ElastigroupDeleteHandler handles deleting an elastigroup from SpotInst
func (s *server) ElastigroupDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w = LogWriter{w}
	vars := mux.Vars(r)
	account := vars["account"]
	esService, ok := s.elastigroupServices[account]
	if !ok {
		log.Errorf("account not found: %s", account)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	elastigroup := vars["elastigroup"]

	err := esService.DeleteAWSElastigroupByID(r.Context(), elastigroup)
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
