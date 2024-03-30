/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	Version   string = "dev"
	Commit    string = "---"
	BuildDate string = "---"
)

var (
	address     string
	port        int
	proxy       string
	vmid        string
	username    string
	token       string
	secret      string
	viewer_path string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "pveSpiceRun",
	Short:   "Simple client to start vm/lxc on proxmox and attach Spice console",
	Long:    fmt.Sprintf("Simple client to start vm/lxc on proxmox and attach Spice console\nVersion: %s compiled on: %s from commit: %s", Version, BuildDate, Commit),
	Run:     rootCommandExecute,
	Version: Version,
}

func Execute() {

	rootCmd.SetVersionTemplate(Version)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().StringVarP(&address, "address", "a", "", "host address")
	rootCmd.MarkPersistentFlagRequired("address")

	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 0, "host port")

	rootCmd.PersistentFlags().StringVarP(&proxy, "proxy", "x", "", "spice proxy, if not provided then equals \"address\"")

	rootCmd.PersistentFlags().StringVarP(&vmid, "vmid", "i", "", "VM/LXC ID")
	rootCmd.MarkPersistentFlagRequired("vmid")

	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "username")
	rootCmd.MarkPersistentFlagRequired("username")

	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "token")
	rootCmd.MarkPersistentFlagRequired("token")

	rootCmd.PersistentFlags().StringVarP(&secret, "secret", "s", "", "secret")
	rootCmd.MarkPersistentFlagRequired("secret")

	default_virtviewer_path := "/usr/bin/remote-viewer"
	os := runtime.GOOS
	if os == "windows" {
		default_virtviewer_path = "C:\\ProgramData\\Microsoft\\Windows\\Start Menu\\Programs\\VirtViewer\\Remote viewer.lnk"
	}
	rootCmd.PersistentFlags().StringVarP(&viewer_path, "viewer_path", "", default_virtviewer_path, "path to remote-viewer")
}

func rootCommandExecute(cmd *cobra.Command, args []string) {

	// authHeader := map[string]string{"Authorization": "PVEAPIToken=" + username + "!" + token + "=" + secret}
	// apiUrl := "https://" + address + "/api2/json"
	// if port != 0 {
	// 	apiUrl = "https://" + address + ":" + strconv.Itoa(port) + "/api2/json"
	// }
	// if proxy == "" {
	// 	proxy = address
	// }

}
