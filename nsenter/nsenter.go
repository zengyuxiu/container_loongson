package nsenter

/*
#define _GNU_SOURCE
#include <errno.h>
#include <sched.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>
#include <unistd.h>

__attribute__((constructor)) void enter_namespace(void) {
	char *container_pid;
	container_pid = getenv("container_pid");
	if (container_pid) {
		fprintf(stdout, "got container_pid=%s\n", container_pid);
	} else {
		fprintf(stdout, "missing container_pid env skip nsenter");
		return;
	}
	char *container_cmd;
	container_cmd = getenv("container_cmd");
	if (container_cmd) {
		fprintf(stdout, "got container_cmd=%s\n", container_cmd);
	} else {
		fprintf(stdout, "missing container_cmd env skip nsenter");
		return;
	}
	int i;
	char nspath[1024];
	char *namespaces[] = { "ipc", "uts", "net", "pid", "mnt" };

	for (i=0; i<5; i++) {
		sprintf(nspath, "/proc/%s/ns/%s", container_pid, namespaces[i]);
		int fd = open(nspath, O_RDONLY,0644);

		if (setns(fd, 0) == -1) {
			fprintf(stderr, "setns on %s namespace failed: %s\n", namespaces[i], strerror(errno));
		} else {
			fprintf(stdout, "setns on %s namespace succeeded\n", namespaces[i]);
		}
		close(fd);
	}
	int res = system(container_cmd);
	exit(0);
	return;
}
*/
import "C"
