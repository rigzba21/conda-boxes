use kube_client::{Client, CustomResourceExt};
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};
use k8s_openapi::serde_json::json;
use kube::CustomResource;

#[derive(CustomResource, Clone, Debug, Deserialize, Serialize, JsonSchema)]
#[kube(group = "rigzba21.k8sconda.io", version = "v1", kind = "CondaEnvironment", namespaced)]
struct CondaEnvSpec {
    name: String,    
    info: String,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    //create  cluster client using the default context or .kube/config
    // or the $KUBECONFIG environment variable
    let _client = Client::try_default().await?;
    let _apiserver_version = _client.apiserver_version().await.unwrap();
    println!("{:#?}", _apiserver_version);

    Ok(())
}
