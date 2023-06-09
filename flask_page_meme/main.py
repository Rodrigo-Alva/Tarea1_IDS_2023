from flask import Flask,render_template
import requests
import json


def getMeme ():
    url="https://meme-api.com/gimme"
    response=json.loads(requests.get(url).text)#aca con .text luego vere el motivo
    meme_img=response["preview"][2]
    subreddit=response["subreddit"]
    return meme_img, subreddit

app=Flask(__name__)#aca le decimos que todo lo que necesita para correr este servidor web esta en este archivo y carpeta
#@ == decorador de python
@app.route("/")#esto para navegar entre las urls

def index():
    meme_img, subreddit=getMeme()
    return render_template("meme_index.html", meme_img=meme_img, subreddit=subreddit)


if __name__=='__main__':
    app.run(debug=True)
