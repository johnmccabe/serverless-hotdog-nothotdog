# Serverless Hotdog Not-Hotdog

[![serverless hotdog not hotdog video](https://img.youtube.com/vi/Ga3-jzQlp6U/0.jpg)](http://www.youtube.com/watch?v=Ga3-jzQlp6U)

# Functions
This demo consists of two OpenFaaS functions which you can deploy yourself.

## hotdog-classifier
This is a forked version of the work by [Magnus Erik Hvass Pedersen](https://github.com/Hvass-Labs/TensorFlow-Tutorials) - its a tweaked version of the [faas-and-furious/inception-function](https://github.com/faas-and-furious/inception-function) which has been updated to classify images submitted as `base64` image data in a POST body (rather than a submitted URL).

## hotdog-app
This is a very basic function which returns a static webpage which acts as a frontend to the `hotdog-classifier` function above.

> Note that only Android or Desktops browsers are supported (Chrome/Firefox) as it uses the `getUserMedia` web API to capture the image to be classified from a local camera.
