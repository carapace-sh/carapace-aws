package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_describeRouterConfigurationCmd = &cobra.Command{
	Use:   "describe-router-configuration",
	Short: "Details about the router.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_describeRouterConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_describeRouterConfigurationCmd).Standalone()

		directconnect_describeRouterConfigurationCmd.Flags().String("router-type-identifier", "", "Identifies the router by a combination of vendor, platform, and software version.")
		directconnect_describeRouterConfigurationCmd.Flags().String("virtual-interface-id", "", "The ID of the virtual interface.")
		directconnect_describeRouterConfigurationCmd.MarkFlagRequired("virtual-interface-id")
	})
	directconnectCmd.AddCommand(directconnect_describeRouterConfigurationCmd)
}
