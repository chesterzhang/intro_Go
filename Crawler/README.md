### Docker 与 ElasticSearch 配置 步骤

1. 安装 docker, 如果是 windows 系统, 还需要 安装 WSL, 全程默认安装即可
2. 拉取 ElasticSearch7.42 镜像 
``` 
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.4.2
```
3. 启动 ElasticSearch7.42 环境
``` 
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.4.2
``` 


### IDEA 访问 ElasticSearch 数据 步骤
1. IDEA 点击 Tools -> HTTP Client -> Test RESTFUL Service

2. 如果查询所有对象, HTTP Method 选择 GET 方法, host/port 输入 http://localhost:9200, Path 输入  index/type/_search (在这个项目中就是 dating_profile/zhenai/_search), Request Parameters: pretty = true, size = 100(自选)

3. 如果删除所有对象, host/port 输入 http://localhost:9200, Path 输入 _all, Request Parameters: pretty = true

4. 如果删除指定对象 , host/port 输入 http://localhost:9200, Path index/type/id, Request Parameters: pretty = true

5. 如果删除一个 index/type 内的所有对象,host/port 输入 http://localhost:9200, Path 输入 index_name/type_name (如 test1/zhenai)  , Request Parameters: pretty = true




