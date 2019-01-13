mkdir -p /firmament/config/
\cp firmament_scheduler.cfg /firmament/config/firmament_scheduler.cfg

kubectl delete -f firmament-deployment.yaml  
kubectl delete -f heapster-poseidon.yaml  
kubectl delete -f poseidon-deployment.yaml

sleep 30

kubectl create -f firmament-deployment.yaml  
sleep 5
kubectl create -f heapster-poseidon.yaml  
sleep 5
kubectl create -f poseidon-deployment.yaml
