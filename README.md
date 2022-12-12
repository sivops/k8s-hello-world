# k8s-hello-world

`go` folder contains application related files as well as Dockerfile \
`k8s` folder contains the kubernetes deployment files

To create a docker image of your choice please follow Image build process onwards, ifnot please skip to Deployment step

## Image build
Change directory to `go` folder and run the following commands to build, tag and push to docker hub
```
docker build -t <image_name>:<tag> .
docker tag <source_image_name>:<tag> <hub-user>/<repo-name>[:<tag>]
docker push <hub-user>/<repo-name>[:<tag>]
```
Note: Make sure you are logged in docker hub on local using ```docker login``` to run above commands
## Deployment
Here I am using minikube so run `minikube start` to start the k8s cluster
Now change the directory to k8s folder and apply the `deploy.yaml` and `svc.yaml`
Note: make sure to change the image name if you created the docker image of your choice
After applying the files, you will see pods and service running as shown in below sample output
```
➜  ~ kubectl get po
NAME                             READY   STATUS    RESTARTS   AGE
sample-go-app-7c6848db9d-ms6f5   1/1     Running   0          59m
➜  ~ kubectl get svc
NAME                TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
kubernetes          ClusterIP   10.96.0.1       <none>        443/TCP          60m
sample-go-app-svc   NodePort    10.107.59.170   <none>        3000:30370/TCP   47m
➜  ~
```
## Test the deployment
To test the k8s deployment, I always use port-forward method as below explains
Run the below k8s command to port-forward the deployment to your local
```
kubectl port-forward service/sample-go-app-svc 9000:3000
```
Now if you hit browser with `http://localhost:9000/` and `http://localhost:9000/health_check` you should see as below output

<img width="660" alt="Screen Shot 2022-12-12 at 5 06 00 PM" src="https://user-images.githubusercontent.com/65404162/207165317-29952193-07d1-4abe-89b3-33582db6190e.png">

<img width="471" alt="Screen Shot 2022-12-12 at 5 07 27 PM" src="https://user-images.githubusercontent.com/65404162/207165521-174da701-1596-4587-ac4f-a1269d9d9a63.png">

