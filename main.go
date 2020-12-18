package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

func createRegistryContainerUnlessAlreadyExists(name string, port int) string {

	result := fmt.Sprintf(`		
		running="$(docker inspect -f '{{.State.Running}}' "%s" 2>/dev/null || true)"
		if [ "${running}" != 'true' ]; then
		docker run \
			-d --restart=always -p "%s:5000" --name "%s" \
			registry:2
		fi			
	`, name, port, name)

	return result
}

func main() {
	reg_name := flag.String("name", "kind-registry", "enter the name of the kind registry")

	reg_port := flag.Int("port", 5000, "enter the port of the kind registry")

	cmd := exec.Command("sh", createRegistryContainerUnlessAlreadyExists(reg_name, reg_port))

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", stdoutStderr)
}
