docker ps -a | awk '{print $1}' | xargs -i{} docker rm {}

docker images | grep "^<none" | awk '{print $3}' | xargs -i{} docker rmi {}
docker images | grep "scaler" | grep "<none" | awk '{print $3}' | xargs -i{} docker rmi {}

docker build --platform linux/amd64 -t registry.cn-shanghai.aliyuncs.com/namespacetest1/scaler:v1.0 .
docker push registry.cn-shanghai.aliyuncs.com/namespacetest1/scaler:v1.0

# tbh2akaswlyrqcp1ee@aliyun.com Test1234
