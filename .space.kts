job("Build and push Docker") {
     startOn {
        gitPush { 
        	branchFilter {
                +"refs/heads/master"
            }
        }
     }
     
    docker {
        build {
            context = "."
            file = "./docker/Dockerfile"
            labels["vendor"] = "Bermos Inc."
            args["HTTP_PROXY"] = "http://10.20.30.1:123"
        }

        push("bermos.registry.jetbrains.space/p/claerance/containers/claerance") {
            tags("version1.0")
        }
    }
}
