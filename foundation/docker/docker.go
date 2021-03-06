package docker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type Container struct {
	Id   string
	Host string
}

// StartContainer starts docker image with extracting id and port of the container
// then constructs and returns Container
func StartContainer(image string, port string, args ...string) (*Container, error) {

	arg := []string{"run"}
	arg = append(arg, "-P", "-d")
	arg = append(arg, args...)
	arg = append(arg, image)

	var out bytes.Buffer
	cmd := exec.Command("docker", arg...)
	fmt.Print("docker")
	fmt.Println(arg)
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	id := string(out.Next(12))
	fmt.Println(id)
	host, err := ExtractHost(id, port)
	if err != nil {
		return nil, err
	}
	return &Container{
		Id:   id,
		Host: host,
	}, nil
}

func StopContainer(id string) error {
	cmd := exec.Command("docker", "stop", id)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("could not stop the container: %w", err)
	}
	fmt.Printf("Stopped: %s\n", id)

	//cmd = exec.Command("docker", "rm", id)
	return nil

}

func ExtractHost(id string, port string) (string, error) {
	tmpl := fmt.Sprintf(" | jq '.[].NetworkSettings.Ports.\"%s/tcp\" | .[]'", port)
	dockerCmd := fmt.Sprintf("docker inspect %s", id)
	dockerCmd += tmpl

	var out bytes.Buffer
	cmd := exec.Command("/bin/sh", "-c", dockerCmd)
	fmt.Println(dockerCmd)
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("cannot inspect docker %s", err)
	}

	var hst struct {
		HostIp   string
		HostPort string
	}
	if err := json.Unmarshal(out.Bytes(), &hst); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s:%s", hst.HostIp, hst.HostPort), nil
}
