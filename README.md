# gaemu : Local server for Google App Engine
 
![gaemu](https://user-images.githubusercontent.com/4569916/103453053-2ed27e00-4d19-11eb-81cb-1df064c29477.png)

## What is

gaemu is local server for Google App Engine(Standard Environment).

gaemu provides a reverse proxy and a static contents server.

( It does not provide APIs for datastores, Memcache and ... other Google App Engine APIs. )

## Install

[Get here.](https://github.com/kurehajime/gaemu/releases)

or 

```
go get github.com/kurehajime/gaemu
```

## Usage 

1. Run Your app

`go run` ? `python main.py` ? `python manage.py runserver` ? `npm start` ?

Please refer to [Google's Help Page](https://cloud.google.com/appengine/docs/standard/python3/testing-and-deploying-your-app).

2. Lanch gaemu

```
gaemu -a "http://localhost:8080" -y "../YOUR_APP_PATH/app.yaml"
```

| option | description | default
----|---- |---- 
| -a | Your App URL | http://localhost:8080
| -y | Location of app.yaml  | 
| -p | gaemu's port  | 8081
