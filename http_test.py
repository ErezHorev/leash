import requests

url = "https://api.github.com/repos/docker/docker/issues/18614"
# url = "https://api.github.com/repos/Elastifile/ecs/issues/6"
r = requests.get(url)
result = r.json()
print(result["number"], result["state"])
