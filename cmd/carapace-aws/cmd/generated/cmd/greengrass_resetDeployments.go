package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_resetDeploymentsCmd = &cobra.Command{
	Use:   "reset-deployments",
	Short: "Resets a group's deployments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_resetDeploymentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_resetDeploymentsCmd).Standalone()

		greengrass_resetDeploymentsCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
		greengrass_resetDeploymentsCmd.Flags().String("force", "", "If true, performs a best-effort only core reset.")
		greengrass_resetDeploymentsCmd.Flags().String("group-id", "", "The ID of the Greengrass group.")
		greengrass_resetDeploymentsCmd.MarkFlagRequired("group-id")
	})
	greengrassCmd.AddCommand(greengrass_resetDeploymentsCmd)
}
