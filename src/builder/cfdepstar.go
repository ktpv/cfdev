package main

// ./scripts/build-cf-deps-tar -m output/cf.yml -c output/cloud-config.yml -r images/cf/configs/dns-runtime-config.yml
func buildCfDepsTar(cfYml, cloudYml, dnsYml string) error {
	// download_warden_stemcell
	// download_compiled_releases
	compileReleases(cfYml)
	compileReleases(dnsYml)
	// tar_deps
	return nil
}

func compileReleases(file string) {
}
