# Google Cloud Vision API
POC - Do not use in production. The intend of this repository is to explore Google Cloud Vision capabilities.
It can serve as a starting point for a custom integration.

## Google Cloud Setup
1. Sign up for Google Cloud
2. Go to API Manager -> Credentials
3. Create Credentials -> Service account key. Download as JSON
4. Go to API Manager -> Dashboard
5. Search for Google Cloud Vision API and enable it

## Building & Running
Get the project ```github.com/fallenhitokiri/cloudvisionapi```.
Run it ```GOOGLE_APPLICATION_CREDENTIALS=/path/to/credentials.json go run -ldflags -s cmd/cloudvisionapi.go```

You can skip the ```-ldflags -s``` part if you are not using the latest XCode release -
this is a known bug and hopefully fixed soon.
 
Post an image to the running service
```curl -X "POST" "http://localhost:8000/" -H "Accept: application/json" -H "Content-Type: multipart/form-data" -H "boundary=Nonce" -F "image=@/path/to/image.jpg"```