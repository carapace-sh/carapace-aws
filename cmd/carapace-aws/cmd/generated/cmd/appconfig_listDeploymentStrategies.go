package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_listDeploymentStrategiesCmd = &cobra.Command{
	Use:   "list-deployment-strategies",
	Short: "Lists deployment strategies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_listDeploymentStrategiesCmd).Standalone()

	appconfig_listDeploymentStrategiesCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	appconfig_listDeploymentStrategiesCmd.Flags().String("next-token", "", "A token to start the list.")
	appconfigCmd.AddCommand(appconfig_listDeploymentStrategiesCmd)
}
