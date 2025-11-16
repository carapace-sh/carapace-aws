package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowDeviceManagementCmd = &cobra.Command{
	Use:   "snow-device-management",
	Short: "Amazon Web Services Snow Device Management documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowDeviceManagementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snowDeviceManagementCmd).Standalone()

	})
	rootCmd.AddCommand(snowDeviceManagementCmd)
}
