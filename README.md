```
curl -v -X POST -d '{"Arr":[2,345,43,99,3,111,61,2,1,9,1], "name":"Casey"}' http://34.139.109.32/
```

```
for x in $(seq 100); do curl -X POST -d '{"Arr":[99,3,111,61,2,1,9,1], "name":"Casey"}' http://34.139.109.32/; sleep 10s;echo "\n";done
```

docker build -t docker.io/cmwylie19/test-argo-rollouts:v2 .; docker push docker.io/cmwylie19/test-argo-rollouts:v2


k argo rollouts get rollout app-rollout --watch      

k get ro app -oyaml  