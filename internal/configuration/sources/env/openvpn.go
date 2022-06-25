package env

import (
	"fmt"
	"strings"

	"github.com/qdm12/gluetun/internal/configuration/settings"
	"github.com/qdm12/govalid/binary"
)

func (r *Reader) readOpenVPN() (
	openVPN settings.OpenVPN, err error) {
	defer func() {
		err = unsetEnvKeys([]string{"OPENVPN_CLIENTKEY", "OPENVPN_CLIENTCRT"}, err)
	}()

	openVPN.Version = getCleanedEnv("OPENVPN_VERSION")
	openVPN.User = r.readOpenVPNUser()
	openVPN.Password = r.readOpenVPNPassword()
	confFile := getCleanedEnv("OPENVPN_CUSTOM_CONFIG")
	if confFile != "" {
		openVPN.ConfFile = &confFile
	}

	ciphersKey, _ := r.getEnvWithRetro("OPENVPN_CIPHERS", "OPENVPN_CIPHER")
	openVPN.Ciphers = envToCSV(ciphersKey)

	auth := getCleanedEnv("OPENVPN_AUTH")
	if auth != "" {
		openVPN.Auth = &auth
	}

	openVPN.ClientCrt, err = readBase64OrNil("OPENVPN_CLIENTCRT")
	if err != nil {
		return openVPN, fmt.Errorf("environment variable OPENVPN_CLIENTCRT: %w", err)
	}

	openVPN.ClientKey, err = readBase64OrNil("OPENVPN_CLIENTKEY")
	if err != nil {
		return openVPN, fmt.Errorf("environment variable OPENVPN_CLIENTKEY: %w", err)
	}

	openVPN.KeyPassphrase = r.readOpenVPNKeyPassphrase()

	openVPN.PIAEncPreset = r.readPIAEncryptionPreset()

	openVPN.IPv6, err = envToBoolPtr("OPENVPN_IPV6")
	if err != nil {
		return openVPN, fmt.Errorf("environment variable OPENVPN_IPV6: %w", err)
	}

	openVPN.MSSFix, err = envToUint16Ptr("OPENVPN_MSSFIX")
	if err != nil {
		return openVPN, fmt.Errorf("environment variable OPENVPN_MSSFIX: %w", err)
	}

	_, openVPN.Interface = r.getEnvWithRetro("VPN_INTERFACE", "OPENVPN_INTERFACE")

	openVPN.ProcessUser, err = r.readOpenVPNProcessUser()
	if err != nil {
		return openVPN, err
	}

	openVPN.Verbosity, err = envToIntPtr("OPENVPN_VERBOSITY")
	if err != nil {
		return openVPN, fmt.Errorf("environment variable OPENVPN_VERBOSITY: %w", err)
	}

	flagsStr := getCleanedEnv("OPENVPN_FLAGS")
	if flagsStr != "" {
		openVPN.Flags = strings.Fields(flagsStr)
	}

	return openVPN, nil
}

func (r *Reader) readOpenVPNUser() (user string) {
	_, user = r.getEnvWithRetro("OPENVPN_USER", "USER")
	// Remove spaces in user ID to simplify user's life, thanks @JeordyR
	return strings.ReplaceAll(user, " ", "")
}

func (r *Reader) readOpenVPNPassword() (password string) {
	_, password = r.getEnvWithRetro("OPENVPN_PASSWORD", "PASSWORD")
	return password
}

func (r *Reader) readOpenVPNKeyPassphrase() (passphrase *string) {
	passphrase = new(string)
	*passphrase = getCleanedEnv("OPENVPN_KEY_PASSPHRASE")
	if *passphrase == "" {
		return nil
	}
	return passphrase
}

func readBase64OrNil(envKey string) (valueOrNil *string, err error) {
	value := getCleanedEnv(envKey)
	if value == "" {
		return nil, nil //nolint:nilnil
	}

	decoded, err := decodeBase64(value)
	if err != nil {
		return nil, err
	}

	return &decoded, nil
}

func (r *Reader) readPIAEncryptionPreset() (presetPtr *string) {
	_, preset := r.getEnvWithRetro(
		"PRIVATE_INTERNET_ACCESS_OPENVPN_ENCRYPTION_PRESET",
		"PIA_ENCRYPTION", "ENCRYPTION")
	if preset != "" {
		return &preset
	}
	return nil
}

func (r *Reader) readOpenVPNProcessUser() (processUser string, err error) {
	key, value := r.getEnvWithRetro("OPENVPN_PROCESS_USER", "OPENVPN_ROOT")
	if key == "OPENVPN_PROCESS_USER" {
		return value, nil
	}

	// Retro-compatibility
	if value == "" {
		return "", nil
	}
	root, err := binary.Validate(value)
	if err != nil {
		return "", fmt.Errorf("environment variable %s: %w", key, err)
	}
	if root {
		return "root", nil
	}
	const defaultNonRootUser = "nonrootuser"
	return defaultNonRootUser, nil
}
