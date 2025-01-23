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


# INSTALLATION

## Operator Installation
To install the operator, run the following command:

```bash
helm repo add k8sgpt https://charts.k8sgpt.ai/
helm repo update
helm install release k8sgpt/k8sgpt-operator -n k8sgpt-operator-system --create-namespace
```

This will install the Operator into the cluster, which will await a K8sGPT resource before anything happens.

## Create secret for backend LLM model if using Ollama API (which is PAID version). SKIP for local running ollaama model

## Update the baseUrl in `backend-ollama-local.yaml` to point to the ollama server and apply the yaml file
```bash
kubectl apply -f backend-ollama-local.yaml
```

