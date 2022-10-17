# webhook-test

1. Install Cert Manager: https://cert-manager.io/docs/installation/ This is needed for webhooks to work. Webhooks require TLS
2. Create certs for the webhook application: 
    ```bash
   kubectl create ns hello-webhook
   kubectl apply -f manifests/cert-manager.yaml
   ```
3. Create deployment, services, and webhook:
   ```bash
   kubectl apply -f manifests/deployment.yaml
   kubectl apply -f manifests/service.yaml
   kubectl apply -f manifests/webhook.yaml   
   ```
4. The webhook pod prints API group and kind then approves it. So open another terminal and follow logs:
   ```bash
   kubectl logs -n hello-webhook <POD> hello-webhook -f
   ```
5. Apply object-storage. Note that this CR doesn't exist on the cluster yet.
   ```bash
    $ kubectl apply -f manifests/object-storage.yaml                                          
    Error from server (NotFound): error when creating "manifests/object-storage.yaml": the server could not find the requested resource (post objectstorages.awsblueprints.io)
   ```
6. Observe that nothing is printed in the pod stdout
7. Apply CRD:
   ```bash
   kubectl apply -f manifests/example-crd.yaml
   ```
8. ```bash
    $ kubectl apply -f manifests/object-storage.yaml                                          
    objectstorage.awsblueprints.io/standard-object-storage created
   ```
9. Observe the pod prints out expected outputs.
