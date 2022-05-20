# README
- image rsync service
- project default URL: http://127.0.0.1:8080/


### build/deploy
- build package
```
# 执行 go build, 并制作 images
bash cmd/linux_build.sh
```

- kubernetes deploy
```
kubectl create namespace go
kubectl -n go create configmap go-default-service-configmap --from-file=configs/config.yaml
kubectl apply -f build/go_default_service-deploy.yaml
```

### 设计
- 系统图
![系統圖](./doc/img/imagersyncv1.0-system.png)

#### API 协议设计
