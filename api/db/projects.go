package db

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Project struct {
	ProjectName                    string    `json:"project_name"`
	ProjectDescription             string    `json:"project_description"`
	ProjectStartDate               string    `json:"project_start_date"`
	ProjectEndDate                 string    `json:"project_end_date"`
	ProjectPriority                string    `json:"project_priority"`
	ProjectFrontEndGithub          string    `json:"project_front_end_github"`
	ProjectBackEndGithub           string    `json:"project_back_end_github"`
	ProjectFrontEndUrl             string    `json:"project_front_end_url"`
	ProjectBackEndUrl              string    `json:"project_back_end_url"`
	ProjectFrontEndHostingUrl      string    `json:"project_front_end_hosting_url"`
	ProjectBackEndHostingUrl       string    `json:"project_back_end_hosting_url"`
	ProjectFrontEndHostingProvider string    `json:"project_front_end_hosting_provider"`
	ProjectBackEndHostingProvider  string    `json:"project_back_end_hosting_provider"`
	CreatedAt                      time.Time `json:"created_at"`
	UpdatedAt                      time.Time `json:"updated_at"`
	Status                         string    `json:"status"`
	OwnerID                        string    `json:"owner_id"`
	ProjectID                      string    `json:"project_id"`
	AllowedUsers                   []string  `json:"allowed_users"`
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var project Project
	_ = json.NewDecoder(r.Body).Decode(&project)
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()
	project.Status = "active"
	project.OwnerID = r.Header.Get("User")
	project.AllowedUsers = []string{r.Header.Get("User")}

	client, err := GetFirebase().Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	_, _, err = client.Collection("projects").Add(ctx, project)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Project created successfully")

}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := GetFirebase().Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	iter := client.Collection("projects").Documents(ctx)
	var projects []Project
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var project Project
		doc.DataTo(&project)
		project.ProjectID = doc.Ref.ID
		if strings.Contains(strings.Join(project.AllowedUsers, ","), w.Header().Get("User")) && w.Header().Get("User") != "" {
			projects = append(projects, project)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projects)
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	vars := mux.Vars(r)
	projectID := vars["project_id"]
	client, err := GetFirebase().Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := client.Collection("projects").Doc(projectID).Get(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	var project Project
	doc.DataTo(&project)
	project.ProjectID = doc.Ref.ID

	if strings.Contains(strings.Join(project.AllowedUsers, ","), w.Header().Get("User")) && w.Header().Get("User") != "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(project)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
	}
}
