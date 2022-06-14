## Prometheus-Client

### Steps to deploy
1. `docker build -t username/prometheus-client .`
2. `docker push username/prometheus-client:latest`
3. `kubectl create -f deployment.yml`
4. `kubectl create -f service.yml`
5. `kubectl port-forward service/prometheus-client-svc 5555:5555`
5. Open [localhost:5555/metrics](http://localhost:5555/metrics) in the browser
