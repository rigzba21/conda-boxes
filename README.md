# conda-boxes
A Kubernetes Controller and CLI to manage [Conda environments](https://docs.conda.io/projects/conda/en/latest/user-guide/tasks/manage-environments.html#creating-an-environment-file-manually)
as volumes or "Boxes" that can be mounted to Pods inside Kubernetes. 

⚠️ This project is a _Work-In-Progress Experiment + Proof-Of-Concept!_ ⚠️

When [building custom Singleuser Jupyter Notebook Containers](https://zero-to-jupyterhub.readthedocs.io/en/latest/jupyterhub/customizing/user-environment.html#customize-an-existing-docker-image)
Conda environments with a large number of dependencies not only increases container build time, but can significantly increase spawn time 
during scheduling, particularly if the container has not been pulled to the node running the Pod. 

The idea with `conda-boxes` is to manage Conda environments outside of the container build and mount them as volumes during Singleuser Pod scheduling.


### Usage:

>🤔 maybe this would be better as a 
[`krew plugin`](https://krew.sigs.k8s.io/docs/developer-guide/develop/plugin-development/)...

```bash
kubectl conda-box my-conda-environment.yaml
```


