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


