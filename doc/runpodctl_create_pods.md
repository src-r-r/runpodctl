## runpodctl create pods

create a group of pods

### Synopsis

create a group of pods on runpod.io

```
runpodctl create pods [flags]
```

### Options

```
      --args string             container arguments
      --communityCloud          create in community cloud
      --containerDiskSize int   container disk size in GB (default 20)
      --cost float32            $/hr price ceiling, if not defined, pod will be created with lowest price available
      --env strings             container arguments
      --gpuCount int            number of GPUs for the pod (default 1)
      --gpuType string          gpu type id, e.g. 'NVIDIA GeForce RTX 3090'
  -h, --help                    help for pods
      --imageName string        container image name
      --mem int                 minimum system memory needed (default 20)
      --name string             any pod name for easy reference
      --podCount int            number of pods to create with the same name (default 1)
      --ports strings           ports to expose; max only 1 http and 1 tcp allowed; e.g. '8888/http'
      --secureCloud             create in secure cloud
      --vcpu int                minimum vCPUs needed (default 1)
      --volumePath string       container volume path (default "/runpod")
      --volumeSize int          persistant volume disk size in GB (default 1)
```

### SEE ALSO

* [runpodctl create](runpodctl_create.md)	 - create a resource

###### Auto generated by spf13/cobra on 8-Feb-2023
