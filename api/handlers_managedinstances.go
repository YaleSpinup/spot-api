package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/YaleSpinup/apierror"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spotinst/spotinst-sdk-go/service/managedinstance/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
)

type ManagedInstanceListDetails struct {
	ID string `json:"id"`
}

// ManagedInstanceDetails defines the information about a managed instance
type ManagedInstanceDetails struct {
	ID         string              `json:"id"`
	CreatedAt  string              `json:"created_at,omitempty"`
	ModifiedAt string              `json:"modified_at,omitempty"`
	Name       string              `json:"name"`
	Size       string              `json:"type,omitempty"`
	State      string              `json:"state"`
	Tags       []map[string]string `json:"tags"`
}

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

	instances := []ManagedInstanceListDetails{}

	for _, i := range output {
		instances = append(instances, ManagedInstanceListDetails{spotinst.StringValue(i.ID)})
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

	instanceDetails := ManagedInstanceDetails{
		ID:         *output.ID,
		Name:       *output.Name,
		CreatedAt:  output.CreatedAt.Format("2006/01/02 15:04:05"),
		ModifiedAt: output.UpdatedAt.Format("2006/01/02 15:04:05"),
		Size:       *output.Compute.LaunchSpecification.InstanceTypes.PreferredType,
		Tags:       transformTags(Org, *output.Name, output.Compute.LaunchSpecification.Tags),
	}

	j, err := json.Marshal(instanceDetails)
	if err != nil {
		log.Errorf("cannot marshal response (%v) into JSON: %s", output, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// ManagedInstanceStatusHandler handles getting status details about a managed instance from SpotInst
func (s *server) ManagedInstanceStatusHandler(w http.ResponseWriter, r *http.Request) {
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

	// we'll try to get the instance first in case it doesn't exist
	// since the status call does not return a not found error
	if _, err := miService.GetAWSManagedInstanceByID(r.Context(), managedinstance); err != nil {
		handleError(w, err)
		return
	}

	output, err := miService.GetAWSManagedInstanceStatusByID(r.Context(), managedinstance)
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

// ManagedInstanceCostsHandler handles getting costs details about a managed instance from SpotInst
func (s *server) ManagedInstanceCostsHandler(w http.ResponseWriter, r *http.Request) {
	w = LogWriter{w}
	vars := mux.Vars(r)
	params := r.URL.Query()

	account := vars["account"]
	miService, ok := s.managedinstanceServices[account]
	if !ok {
		log.Errorf("account not found: %s", account)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	managedinstance := vars["instance"]

	startDate, endDate, err := parseDate(params.Get("start"), params.Get("end"))
	if err != nil {
		handleError(w, err)
		return
	}

	// we'll try to get the instance first in case it doesn't exist
	// since the costs call does not return a not found error
	if _, err := miService.GetAWSManagedInstanceByID(r.Context(), managedinstance); err != nil {
		handleError(w, err)
		return
	}

	output, err := miService.GetAWSManagedInstanceCostsByID(r.Context(), managedinstance, startDate, endDate)
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

	// use the DefaultVPC if VpcID is not passed
	if req.Compute != nil && req.Compute.VpcID == nil && miService.DefaultVPC != "" {
		req.Compute.VpcID = &miService.DefaultVPC
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

func transformTags(org, name string, tags []*aws.Tag) []map[string]string {
	transformedTags := []map[string]string{}
	for _, t := range tags {
		if *t.Key == "spinup:org" || *t.Key == "Name" {
			continue
		}
		transformedTags = append(transformedTags, map[string]string{*t.Key: *t.Value})
	}

	transformedTags = append(transformedTags,
		map[string]string{"Name": name},
		map[string]string{"spinup:org": org},
	)

	log.Debugf("returning transformed tags: %+v", transformedTags)
	return transformedTags
}

// ManagedVolumesListHandler handles listing managed instance volumes
func (s *server) ManagedVolumesListHandler(w http.ResponseWriter, r *http.Request) {
	w = LogWriter{w}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte{})
}

// parseDate returns date parsed from query parameters
func parseDate(startDate, endDate string) (string, string, error) {
	log.Debugf("startDate: %s, endDate: %s ", startDate, endDate)

	if startDate == "" && endDate == "" {
		return "", "", nil
	}

	if (startDate == "" && endDate != "") || (startDate != "" && endDate == "") {
		return "", "", apierror.New(apierror.ErrBadRequest, "both start and end dates must be specified", nil)
	}

	// sTmp and eTmp temporary vars to hold time.Time objects
	sTmp, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return "", "", apierror.New(apierror.ErrBadRequest, "error parsing start date from input", err)
	}

	eTmp, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return "", "", apierror.New(apierror.ErrBadRequest, "error parsing end date from input", err)
	}

	// check that end date is greater than start date
	timeValidity := eTmp.After(sTmp)
	if !timeValidity {
		return "", "", apierror.New(apierror.ErrBadRequest, "end date should be greater than start date", nil)
	}

	// convert time.Time to a string
	return sTmp.Format("2006-01-02"), eTmp.Format("2006-01-02"), nil
}
