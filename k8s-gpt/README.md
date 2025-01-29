# PRE-REQUISITES
Ollama server and model should be running and accessible from the cluster.

Install ollama on MacOS
```bash
brew install ollama
```

Start ollama server 
```bash
OLLAMA_HOST=0.0.0.0  ollama serve
```

Run the model
```bash
ollama run llama3
```
# Create a ChatBox UI using open-web-ui
```bash
docker run -d -p 9783:8080 -v open-webui:/app/backend/data --name open-webui ghcr.io/open-webui/open-webui:main
```

# INSTALLATION

## Operator Installation
To install the operator, run the following command:

```bash
helm repo add k8sgpt https://charts.k8sgpt.ai/
helm repo update
# helm install release k8sgpt/k8sgpt-operator -n k8sgpt-operator-system --create-namespace

# k8sGPT with Prometheus and Grafana

helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install prometheus prometheus-community/kube-prometheus-stack --namespace monitoring --create-namespace

helm install release k8sgpt/k8sgpt-operator -n k8sgpt-operator-system --create-namespace --set interplex.enabled=true --set grafanaDashboard.enabled=true --set serviceMonitor.enabled=true

```

This will install the Operator into the cluster, which will await a K8sGPT resource before anything happens.

## Create secret for backend LLM model if using Ollama API (which is PAID version). SKIP for local running ollaama model

## Update the baseUrl in `backend-ollama-local.yaml` to point to the ollama server and apply the yaml file
```bash
kubectl apply -f backend-ollama-local.yml
```

## Looking at the results of the analysis made by k8sgpt 
```bash
kubectl get result -n k8sgpt-operator-system -o json | jq .
```

# Secrity scanning using Trivy

```bash
 k8sgpt integration list
```

## Activate Trivy 
```bash
k8sgpt integration activate trivy
```

## New filters are added now 
ConfigAuditReport (integration) and VulnerabilityReport (integration)
```bash
k8sgpt filters list
```



## Option 2 - Run Deepseek model (WORKING)

Install ollama on MacOS
```bash
brew install ollama
```

Start ollama server 
```bash
OLLAMA_HOST=0.0.0.0  ollama serve
```

Run the model = Deepseek
```bash
ollama run deepseek-r1:8b
```

To run custom Model
```bash
ollama create praj-deepseek-r1 -f deepseek-model/Modelfile
ollama run praj-deepseek-r1
```

## Update the baseUrl in `backend-deepseek-local.yaml` to point to the ollama server and apply the yaml file
```bash
kubectl apply -f backend-deepseek-local.yaml
```

This might take anywhere between 10-30 min depending on your laptop configurtion and deepseek model size

## Looking at the results of the analysis made by k8sgpt 
```bash
kubectl get result -n k8sgpt-operator-system -o json | jq .
```