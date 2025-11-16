package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraphCmd = &cobra.Command{
	Use:   "iotthingsgraph",
	Short: "AWS IoT Things Graph\n\nAWS IoT Things Graph provides an integrated set of tools that enable developers to connect devices and services that use different standards, such as units of measure and communication protocols.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraphCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotthingsgraphCmd).Standalone()

	})
	rootCmd.AddCommand(iotthingsgraphCmd)
}
