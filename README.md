# websay

Websay!!!!

# Build image with Cloud Native Buildpacks

pack build -B gcr.io/paketo-buildpacks/builder:tiny martinborgman/websay:2 --publish

# Create deployment

kubectl create deployment websay --image martinborgman/websay:2

# Create service

kubectl create service clusterip websay --tcp=80:1323

# Create ingress controller

kubectl apply -f https://projectcontour.io/quickstart/contour.yaml

# Create HTTPProxy

kubectl apply -f work/ingress.yaml

