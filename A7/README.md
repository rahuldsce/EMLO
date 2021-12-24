## Description 
I was able to successfully run entire experiment from visualization to Explanation, but not able solve prediction setup error `503 internal server error`, I had set DNS config, but was not able expose the URL for external access which resulted in 503 error

To read minio access key and secret key we nee to execute `kubectl get secret mlpipeline-minio-artifact -n kubeflow -o yaml` 

## Setup Kubeflow Pipeline
<br/>

```
minikube start --vm-driver=hyperkit --kubernetes-version v1.20.12 --cpus=7 --memory=14g --disk-size=100g
```

<br/>

### Clone Manifest

```
git clone https://github.com/kubeflow/manifests
cd manifests
```
<br/>

<br/>

## Install kustomize on Mac OS
```
wget https://github.com/kubernetes-sigs/kustomize/releases/download/v3.2.0/kustomize_3.2.0_darwin_amd64
```
<br/>

## For linux change darwin to linux
```
chmod +x kustomize_3.2.0_darwin_amd64
mv kustomize_3.2.0_darwin_amd64 /usr/local/bin/kustomize
```
<br/>

### 5. Install kubectl based on minikube setup
Install kubectl version based on kubernetes version
```
minikube kubectl -- get po -A
```
<br/>

### Create minkube kubectl alias
```
alias kubectl="minikube kubectl --"
```
<br/>

### Install kubeflow piplines
```
while ! kustomize build example | kubectl apply -f -; do echo "Retrying to apply resources"; sleep 10; done
```

<br/>

### Start Kubeflow

```
kubectl port-forward svc/istio-ingressgateway -n istio-system 8080:80
```
<br/>

### Start Minio

```
kubectl port-forward -n kubeflow svc/minio-service 9090:9000
```
<br/>

### Read secret to get minio access and secret key

```
kubectl get secret mlpipeline-minio-artifact -n kubeflow -o yaml
```

<br/>

Apply secret key
```
Kubectl apply -f mino-secret.yaml -n kubeflow-user-example-com
```

<br/>

## DNS Config
Set the ingress ports:
```
export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')
export SECURE_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="https")].nodePort}')
export TCP_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="tcp")].nodePort}')

```
setting the ingress IP depends on the cluster provider:
minikube
```
$ export INGRESS_HOST=$(minikube ip)
```