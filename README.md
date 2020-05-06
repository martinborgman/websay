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

# Add CI/CD using Tekton

kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml

# Because we are using Cloud Native Buildpacks we need to add support in Tekton

kubectl apply -f https://raw.githubusercontent.com/tektoncd/catalog/master/buildpacks/buildpacks-v3.yaml

# Create resource for git repo

kubectl apply -f work/websay-git.yaml 

# Run the image build job

kubectl apply -f work/run-build-image.yaml 

This will fail because the job dous not have access to your Docker Hub account.

# Create secret for Docker Hub account

On Mac and Windows you must first edit your $HOME/.docker/config.json to remove te credStore value and log in te the Docker Hub again. This trick will result in a normal (linux) ~/.docker/config.json
Beware, this trick stores your Docker Hub credentials in a unsecure manner.

kubectl create secret generic regcred \
        --from-file=.dockerconfigjson=/Users/martinborgman/.docker/config.json \
        --type=kubernetes.io/dockerconfigjson

# Create a service account

kubectl create sa build-service

Edit the service account to include regcred.

 