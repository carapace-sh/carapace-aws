package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_getLaunchConfigurationCmd = &cobra.Command{
	Use:   "get-launch-configuration",
	Short: "Lists all LaunchConfigurations available, filtered by Source Server IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_getLaunchConfigurationCmd).Standalone()

	mgn_getLaunchConfigurationCmd.Flags().String("account-id", "", "Request to get Launch Configuration information by Account ID.")
	mgn_getLaunchConfigurationCmd.Flags().String("source-server-id", "", "Request to get Launch Configuration information by Source Server ID.")
	mgn_getLaunchConfigurationCmd.MarkFlagRequired("source-server-id")
	mgnCmd.AddCommand(mgn_getLaunchConfigurationCmd)
}
