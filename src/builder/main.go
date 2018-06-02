package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	// "github.com/cppforlife/go-patch/patch"
)

var rootDir, outputDir string

func init() {
	var err error
	rootDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	outputDir = filepath.Join(rootDir, "output")
}

// ./scripts/generate-cloud-config -c ../cf-deployment/
func generateCloudConfig(cfDeployment string) error {
	// TODO assert that directory exists (and isn't empty string)
	opsDir := filepath.Join(rootDir, "images", "cf", "cf-operations")
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(
		"bosh", "int", "iaas-support/bosh-lite/cloud-config.yml",
		"-o", filepath.Join(opsDir, "set-cloud-config-subnet.yml"),
	)
	cmd.Dir = cfDeployment
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(stderr.String())
		return err
	}
	return ioutil.WriteFile(filepath.Join(outputDir, "cloud-config.yml"), stdout.Bytes(), 0644)
}

// ./scripts/generate-cf-manifest -c ../cf-deployment/
func generateCfManifest(cfDeployment string) error {
	// TODO assert that directory exists (and isn't empty string)
	opsDir := filepath.Join(rootDir, "images", "cf", "cf-operations")
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(
		"bosh", "int", "cf-deployment.yml",
		"-o", "operations/use-compiled-releases.yml",
		"-o", "operations/experimental/skip-consul-cell-registrations.yml",
		"-o", "operations/experimental/skip-consul-locks.yml",
		"-o", "operations/experimental/use-bosh-dns-for-containers.yml",
		"-o", "operations/experimental/disable-consul.yml",
		"-o", "operations/bosh-lite.yml",
		"-o", "operations/experimental/disable-consul-bosh-lite.yml",
		"-o", filepath.Join(opsDir, "allow-local-docker-registry.yml"),
		"-o", filepath.Join(opsDir, "garden-disable-app-armour.yml"),
		"-o", filepath.Join(opsDir, "collocate-tcp-router.yml"),
		"-o", filepath.Join(opsDir, "set-cfdev-subnet.yml"),
		"-o", filepath.Join(opsDir, "lower-memory.yml"),
		"-v", "cf_admin_password=admin",
		"-v", "uaa_admin_client_secret=admin-client-secret",
	)
	cmd.Dir = cfDeployment
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(stderr.String())
		return err
	}
	return ioutil.WriteFile(filepath.Join(outputDir, "cf.yml"), stdout.Bytes(), 0644)
}

// ./scripts/generate-bosh-manifest -b ../bosh-deployment/
func generateBoshManifest(boshDeployment string, uaa, credhub bool) error {
	// TODO assert that directory exists (and isn't empty string)
	opsDir := filepath.Join(rootDir, "images", "cf", "bosh-operations")
	args := []string{
		"int", "bosh.yml",
		"-o", "bosh-lite.yml",
		"-o", "bosh-lite-runc.yml",
		"-o", "bosh-lite-grootfs.yml",
		"-o", "warden/cpi.yml",
		"-o", "warden/cpi-grootfs.yml",
		"-o", "jumpbox-user.yml",
		"-o", filepath.Join(opsDir, "disable-app-armor.yml"),
		"-o", filepath.Join(opsDir, "remove-ports.yml"),
		"-o", filepath.Join(opsDir, "use-warden-cpi-v39.yml"),
		"-o", filepath.Join(opsDir, "use-stemcell-3586.7.yml"),
		"-v", `director_name="warden"`,
		"-v", "internal_cidr=10.245.0.0/24",
		"-v", "internal_gw=10.245.0.1",
		"-v", "internal_ip=10.245.0.2",
		"-v", "garden_host=10.0.0.10",
	}

	if uaa {
		args = append(args, "-o", "uaa.yml")
	}
	if credhub {
		args = append(args, "-o", "credhub.yml")
	}

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("bosh", args...)
	cmd.Dir = boshDeployment
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(stderr.String())
		return err
	}
	return ioutil.WriteFile(filepath.Join(outputDir, "bosh.yml"), stdout.Bytes(), 0644)
}

// ./scripts/generate-bosh-runtime-config -b ../bosh-deployment/
func generateBoshRuntime(boshDeployment string) error {
	// TODO assert that directory exists (and isn't empty string)
	opsDir := filepath.Join(rootDir, "images", "cf", "bosh-operations")
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(
		"bosh", "int", "runtime-configs/dns.yml",
		"-o", filepath.Join(opsDir, "add-host-pcfdev-dns-record.yml"),
	)
	cmd.Dir = boshDeployment
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(stderr.String())
		return err
	}
	return ioutil.WriteFile(filepath.Join(outputDir, "dns-runtime-config.yml"), stdout.Bytes(), 0644)
}

func all() error {
	os.RemoveAll(outputDir)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	fmt.Println("GENERATE CLOUD CONFIG")
	if err := generateCloudConfig(filepath.Join(rootDir, "..", "cf-deployment")); err != nil {
		return err
	}

	fmt.Println("GENERATE CF MANIFEST")
	if err := generateCfManifest(filepath.Join(rootDir, "..", "cf-deployment")); err != nil {
		return err
	}

	// fmt.Println("GENERATE CF DEPS TAR")
	// ./scripts/build-cf-deps-tar -m output/cf.yml -c output/cloud-config.yml -r images/cf/configs/dns-runtime-config.yml

	fmt.Println("GENERATE BOSH MANIFEST")
	if err := generateBoshManifest(filepath.Join(rootDir, "..", "bosh-deployment"), false, false); err != nil {
		return err
	}

	fmt.Println("GENERATE BOSH RUNTIME CONFIG")
	if err := generateBoshRuntime(filepath.Join(rootDir, "..", "bosh-deployment")); err != nil {
		return err
	}

	// fmt.Println("GENERATE BOSH TAR")
	// ./scripts/build-bosh-deps-tar -m output/bosh.yml -r output/dns-runtime-config.yml
	//
	// fmt.Println("GENERATE CF ISO")
	// ./scripts/build-cf-deps-iso -c output/cf.tgz  -b output/bosh.tgz
	//
	// fmt.Println("BUILD EFI IMAGE")
	// ./scripts/build-image
	//
	// fmt.Println("NOW, PLEASE GENERATE CF PLUGIN VIA: $PWD/src/code.cloudfoundry.org/cfdev/generate-plugin.sh")
	return nil
}

func main() {
	if err := all(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
