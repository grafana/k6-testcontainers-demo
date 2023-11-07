package tck6demo

import (
	"bytes"
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/testcontainers/testcontainers-go/modules/k3s"
	"github.com/testcontainers/testcontainers-go/modules/k6"
)



func Test_Demo(t *testing.T) {
	ctx := context.Background()

	// start a k3s container
	k3sContainer, err := k3s.RunContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Clean up the container
	defer func() {
		if err := k3sContainer.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	}()

	err = k3sContainer.CopyFileToContainer(
		ctx,
		"./manifests/quickpizza.yaml",
		"quickpizza.yaml",
		0x644,
	)
	if err != nil {
		t.Fatalf("copying files to cluster %v", err)
		
	}

	_, _, err = k3sContainer.Exec(
		ctx, 
		[]string{
			"kubectl",
			"apply",
			"-f",
			"quickpizza.yaml",
		},
	)
	if err != nil {
		t.Fatalf("deploying app: %v", err)
		
	}

	rc, stdout, err := k3sContainer.Exec(
		ctx, 
		[]string{
			"kubectl",
			"wait",
			"pods",
			"--all",
			"--for=condition=Ready",
			"--timeout=90s",
		},
	)
	if err != nil {
		t.Fatalf("waiting for pods ready %v", err)
	}
	if rc != 0 {
		output := bytes.Buffer{}
		output.ReadFrom(stdout)
		t.Fatalf("pods not ready \n%s\n", output.String())
	}

	k3sIP, err := k3sContainer.ContainerIP(ctx)
	if err != nil {
		t.Fatalf("failed to get k3s IP:\n%v", err)
	}
	frontEndUrl := fmt.Sprintf("http://%s:3333", k3sIP)

	// path to script must be absolute
	scriptPath, err := filepath.Abs("scripts/test.js")
	if err != nil {
		t.Fatalf("failed to get path to test script: %v", err)
	}

	k6Container, err := k6.RunContainer(
		ctx,
		k6.WithCache(),
		k6.WithTestScript(scriptPath),
		k6.SetEnvVar("FRONTEND_URL", frontEndUrl),
	)

	if err != nil {
		t.Fatal(err)
	}

	// Clean up the container after the test is complete
	t.Cleanup(func() {
		if err := k6Container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	// assert the result of the test
	state, err := k6Container.State(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if state.ExitCode != 0 {
		logs := bytes.Buffer{}
		logReader, err := k6Container.Logs(ctx)
		if err != nil {
			t.Logf("getting logs %v", err)
		} else {
			logs.ReadFrom(logReader)
		}
		
		t.Fatalf("test failed with code %d\n%s\n",  state.ExitCode, logs.String())
	}
}