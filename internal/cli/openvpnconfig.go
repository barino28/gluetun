package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/qdm12/gluetun/internal/configuration"
	"github.com/qdm12/gluetun/internal/constants"
	"github.com/qdm12/gluetun/internal/provider"
	"github.com/qdm12/gluetun/internal/storage"
	"github.com/qdm12/golibs/logging"
	"github.com/qdm12/golibs/params"
)

type OpenvpnConfigMaker interface {
	OpenvpnConfig(logger logging.Logger, env params.Interface) error
}

func (c *CLI) OpenvpnConfig(logger logging.Logger, env params.Interface) error {
	var allSettings configuration.Settings
	err := allSettings.Read(env, logger)
	if err != nil {
		return err
	}
	allServers, err := storage.New(logger, constants.ServersData).
		SyncServers(constants.GetAllServers())
	if err != nil {
		return err
	}
	providerConf := provider.New(allSettings.VPN.Provider.Name, allServers, time.Now)
	connection, err := providerConf.GetConnection(allSettings.VPN.Provider.ServerSelection)
	if err != nil {
		return err
	}
	lines := providerConf.BuildConf(connection, allSettings.VPN.OpenVPN)
	fmt.Println(strings.Join(lines, "\n"))
	return nil
}
