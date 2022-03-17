# kubectl-conda
A [`kubectl` Krew](https://krew.sigs.k8s.io/) plugin for managing  [Conda environments](https://docs.conda.io/projects/conda/en/latest/user-guide/tasks/manage-environments.html#creating-an-environment-file-manually)
as volumes in Kubernetes 

⚠️ This project is a _Work-In-Progress Experiment + Proof-Of-Concept!_ ⚠️

### Usage:

```bash
# build it
go build cmd/kubectl-conda.go

# tentative usage...
./kubectl-conda create my-conda-environment.yaml
```

