package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_getLaunchConfigurationCmd = &cobra.Command{
	Use:   "get-launch-configuration",
	Short: "Gets a LaunchConfiguration, filtered by Source Server IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_getLaunchConfigurationCmd).Standalone()

	drs_getLaunchConfigurationCmd.Flags().String("source-server-id", "", "The ID of the Source Server that we want to retrieve a Launch Configuration for.")
	drs_getLaunchConfigurationCmd.MarkFlagRequired("source-server-id")
	drsCmd.AddCommand(drs_getLaunchConfigurationCmd)
}
