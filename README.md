String-Services

To run locally:
1. Start the container
2. curl -H "Content-Type: application/json; charset=UTF-8" -d '{"Input":"विकिपीडिया"}' http://localhost:8080/reverse
3. curl -H "Content-Type: application/json; charset=UTF-8" -d '{"Input":"विकिपीडिया"}' http://localhost:8080/echo

-----------------------------

To run on GCP:
1. push the container on gcloud:
    - docker tag rjain/wobe_v2 us.gcr.io/deep-geography-164119/wobe_v2
    - gcloud docker -- push us.gcr.io/deep-geography-164119/wobe_v2

2. Copy the code repository to gcloud:

3. Deploy services: kubectl create -f service.yaml
4. Deploy the pods: kubectl create -f controller.yaml
5. Check pods: kubectl get pods
6. Check the replication controller: kubectl get rc
7. Check the services: kubectl get services

8. to resize the cluster: gcloud container clusters --zone ZONE resize CLUSTER --size SIZE

9. Test same as locally but with instance IP:port.