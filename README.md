# Golang REST API with Mux

This project is following the tutorials of [Pragmatic Reviews' Golang Crash Course](https://www.youtube.com/playlist?list=PL3eAkoh7fypqUQUQPn-bXtfiYT_ZSVKmB).

## My notes

- More about Gorilla Mux toolkit [here](https://www.gorillatoolkit.org/).

- To use a firebase db do not forget to export the credentials like this:
```
export GOOGLE_APPLICATION_CREDENTIALS="/home/user/firebase/credentials.json"
```

- For the GET and POST operations with curl use:
```
# GET operation
curl localhost:8000/posts

# POST operation
curl -d '{"title":"Title 1", "text":"This book is about ..."}' -H "Content-Type: application/json" -X POST http://localhost:8000/posts
```
