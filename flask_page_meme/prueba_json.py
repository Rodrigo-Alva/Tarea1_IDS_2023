import requests
import pprint
url="https://meme-api.com/gimme"

response = requests.get(url)

response.json() # nos brinda la información en formato JSON
print(response.json())
productos=response.json()
a=productos["postLink"]
print(productos.values())
print(a)
pprint.pprint(productos)  