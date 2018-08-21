# go-devspaces
Hello world Go sample app with instructions to run with Dev Spaces

## Get Started
1. Set up AKS and Azure Dev Spaces
    1. [Prerequisites](https://docs.microsoft.com/en-us/azure/dev-spaces/quickstart-nodejs#prerequisites)
    1. [Set up Dev Spaces](https://docs.microsoft.com/en-us/azure/dev-spaces/quickstart-nodejs#set-up-azure-dev-spaces)   
1. Clone this source repo, and open a Terminal window to the repo's root folder.
1. Generate a Helm chart: `azds prep`
1. Build and run the Go api app in AKS: `azds up`
1. Open web app in a browser (URL appears in `azds up` output)
