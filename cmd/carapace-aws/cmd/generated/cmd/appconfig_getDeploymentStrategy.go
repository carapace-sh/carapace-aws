package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_getDeploymentStrategyCmd = &cobra.Command{
	Use:   "get-deployment-strategy",
	Short: "Retrieves information about a deployment strategy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_getDeploymentStrategyCmd).Standalone()

	appconfig_getDeploymentStrategyCmd.Flags().String("deployment-strategy-id", "", "The ID of the deployment strategy to get.")
	appconfig_getDeploymentStrategyCmd.MarkFlagRequired("deployment-strategy-id")
	appconfigCmd.AddCommand(appconfig_getDeploymentStrategyCmd)
}
