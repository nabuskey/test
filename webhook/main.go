package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	admissionv1 "k8s.io/api/admission/v1"
)

func mutate(w http.ResponseWriter, r *http.Request) {

	admissionReview := &admissionv1.AdmissionReview{}

	log.Println("invoked")
	err := json.NewDecoder(r.Body).Decode(admissionReview)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("error decoding: %s", err)
		return
	}
	log.Printf("group: %s, kind: %s", admissionReview.Request.Resource.Group, admissionReview.Request.Kind)

	admissionResponse := &admissionv1.AdmissionResponse{
		UID:     admissionReview.Request.UID,
		Allowed: true,
	}

	resp := &admissionv1.AdmissionReview{
		Response: admissionResponse,
	}
	resp.SetGroupVersionKind(admissionReview.GroupVersionKind())

	b, err := json.Marshal(resp)
	if err != nil {
		msg := fmt.Sprintf("error marshalling response json: %v", err)
		log.Printf(msg)
		w.WriteHeader(500)
		w.Write([]byte(msg))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
	return
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.HandleFunc("/mutate", mutate)
	err := http.ListenAndServeTLS(fmt.Sprintf(":%s", port), "/tls/tls.crt", "/tls/tls.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
