package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_deleteDeploymentStrategyCmd = &cobra.Command{
	Use:   "delete-deployment-strategy",
	Short: "Deletes a deployment strategy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_deleteDeploymentStrategyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_deleteDeploymentStrategyCmd).Standalone()

		appconfig_deleteDeploymentStrategyCmd.Flags().String("deployment-strategy-id", "", "The ID of the deployment strategy you want to delete.")
		appconfig_deleteDeploymentStrategyCmd.MarkFlagRequired("deployment-strategy-id")
	})
	appconfigCmd.AddCommand(appconfig_deleteDeploymentStrategyCmd)
}
