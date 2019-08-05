# dirty-cow-golang

Dirty Cow implement in Go for validating GKE-Sandbox

#Build Docker image

docker build -t gcr.io/gcp-project-name/rougetest/dirty-cow-golang -f Dockerfile .

#Run and verify on Docker

docker run gcr.io/gcp-project-name/rougetest/dirty-cow-golang:latest

#Create test-sandbox-cluster

gcloud container clusters get-credentials test-sandbox-cluster --zone us-central1-a --project gcp-project-name

#Select your sandbox

kubectl config use-context gke_gcp-project-name_us-central1-a_test-sandbox-cluster

#Push Docker image on GCR

gcloud docker -- push gcr.io/gcp-project-name/rougetest/dirty-cow-golang:latest

#Deploy on Kubernetes (GKE-Sandbox)

kubectl run dirty-cow-golang --image=gcr.io/gcp-project-name/rougetest/dirty-cow-golang:latest

#Update the Deployment.yml by adding following line thus POD can schedule:

runtimeClassName: gvisor (in: spec->template->spec)




