package main

import (
	boshcmd "github.com/cloudfoundry/bosh-cli/cmd"
	cmdconf "github.com/cloudfoundry/bosh-cli/cmd/config"
	boshdir "github.com/cloudfoundry/bosh-cli/director"
	boshui "github.com/cloudfoundry/bosh-cli/ui"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

var privateBoshDir boshdir.Director = nil

func boshDirector() (boshdir.Director, error) {
	if privateBoshDir != nil {
		return privateBoshDir, nil
	}
	logger := boshlog.NewLogger(boshlog.LevelError)
	ui := boshui.NewConfUI(logger)
	basicDeps := boshcmd.NewBasicDeps(ui, logger)
	cmdFactory := boshcmd.NewFactory(basicDeps)
	fs := basicDeps.FS
	cmd, err := cmdFactory.New([]string{"releases"})
	if err != nil {
		return nil, err
	}
	config, err := cmdconf.NewFSConfigFromPath(cmd.BoshOpts.ConfigPathOpt, fs)
	if err != nil {
		return nil, err
	}
	session := boshcmd.NewSessionFromOpts(cmd.BoshOpts, config, ui, true, true, fs, logger)
	privateBoshDir, err = session.Director()
	if err != nil {
		privateBoshDir = nil
	}
	return privateBoshDir, err
}
