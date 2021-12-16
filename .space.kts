job("Build and push Docker") {
     startOn {
        gitPush { 
        	branchFilter {
                +"refs/heads/main"
            }
        }
     }
     
    docker {
        build {
            context = "."
            file = "./docker/Dockerfile"
            labels["vendor"] = "Bermos Inc."
            labels["version"] = "0.1"
            args["HTTP_PROXY"] = "http://10.20.30.1:123"
        }

        push("bermos.registry.jetbrains.space/p/claerance/containers/claerance") {
            tags("0.1")
        }
    }
}
