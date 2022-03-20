use kube_client::{Client, CustomResourceExt};
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};
use kube::CustomResource;
use kube::api::{Api};

#[derive(CustomResource, Clone, Debug, Deserialize, Serialize, JsonSchema, Default)]
#[kube(group = "rigzba21.k8sconda.io", version = "v1", kind = "CondaEnvironment", namespaced)]
struct CondaEnvSpec {
    name: String, //conda env name   
    channel: String, //conda channel
    dependencies: Vec<String>, //dependencies
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    //create  cluster client using the default context or .kube/config
    // or the $KUBECONFIG environment variable
    let _client = Client::try_default().await?;
    
    let _apiserver_version = _client.apiserver_version().await.unwrap();
    println!("{:#?}", _apiserver_version);

    //let _sample_env: Api<CondaEnvironment> = Api::namespaced(_client, "default");
    let _sample = CondaEnvironment::new("sample-env", CondaEnvSpec::default());
    println!("{:#?}", _sample);
    println!("{:?}", serde_yaml::to_string(&CondaEnvironment::crd()));
    Ok(())
}
