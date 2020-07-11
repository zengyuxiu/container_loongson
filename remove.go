package main

import (
	"container_loongson/cgroups"
	"container_loongson/container"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func removeContainer(containerName string) {
	containerInfo, err := getContainerInfoByName(containerName)
	if err != nil {
		log.Errorf("Get container %s info error %v", containerName, err)
		return
	}
	if containerInfo.Status != container.STOP {
		log.Errorf("Couldn't remove running container")
		return
	}

	//释放cg
	cgroupManager := cgroups.NewCgroupManager("container-cgroup")
	defer cgroupManager.Destroy()

	dirURL := fmt.Sprintf(container.DefaultInfoLocation, containerName)
	if err := os.RemoveAll(dirURL); err != nil {
		log.Errorf("Remove file %s error %v", dirURL, err)
		return
	}
}
