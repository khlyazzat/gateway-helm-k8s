## Structure
services/ — source code of services (auth, profile)

kube/— Helm charts and Kubernetes manifests

migrations/ — SQL migrations for each database

## Database
Postgres (you can use the Bitnami chart or minikube-postgres)

## Service Build
helm upgrade --install auth-service kube/auth-service --namespace architecture-lab -f kube/auth-service/values.yaml

helm upgrade --install profile-service kube/profile-service --namespace architecture-lab -f kube/profile-service/values.yaml

### Install PostgreSQL via Helm with custom values
helm install auth-db oci://registry-1.docker.io/bitnamicharts/postgresql -f kube/databases/values/auth-db.yaml

helm install profile-db oci://registry-1.docker.io/bitnamicharts/postgresql -f kube/databases/values/profile-db.yaml

### Redis
helm repo add bitnami https://charts.bitnami.com/bitnami

helm repo update

helm install redis bitnami/redis -f kube/databases/values/redis.yaml --namespace default

### Apply, get, delete commands
kubectl apply -f kube/auth-service/configs/ -n architecture-lab

kubectl apply -f kube/auth-service/secrets/ -n architecture-lab


kubectl delete deployment profile-service -n architecture-lab

kubectl delete job profile-service-migrate -n architecture-lab


kubectl delete deployment auth-service -n architecture-lab

kubectl delete job auth-service-migrate -n architecture-lab


kubectl get pods -n architecture-lab
