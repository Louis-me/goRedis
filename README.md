## 说明

本次代码为go中对reids进行增删改查

## 环境

- win10, go 1.20.4
- 本地启动redis服务

## 运行代码

```
go run main.go
```

## 客户端测试

```python
data ={"name": "test11", "password": "1123456", "id": 1, "key": "t_key1"}
resp = requests.post("http://127.0.0.1:8000/userAdd", json=data)
print(resp.text)
resp = requests.get("http://127.0.0.1:8000/userGet/t_key2")
print(resp.text)

resp = requests.get("http://127.0.0.1:8000/userAll")
print(resp.text)


data ={"name": "test1131", "password": "123456811", "id": 111, "key": "t_key21"}
resp = requests.post("http://127.0.0.1:8000/userUpdate", json=data)
print(resp.text)

data ={"key": "t_key1"}
resp = requests.post("http://127.0.0.1:8000/userDel", json=data)
print(resp.text)


```

