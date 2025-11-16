package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_getDeploymentCmd = &cobra.Command{
	Use:   "get-deployment",
	Short: "Retrieves information about a configuration deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_getDeploymentCmd).Standalone()

	appconfig_getDeploymentCmd.Flags().String("application-id", "", "The ID of the application that includes the deployment you want to get.")
	appconfig_getDeploymentCmd.Flags().String("deployment-number", "", "The sequence number of the deployment.")
	appconfig_getDeploymentCmd.Flags().String("environment-id", "", "The ID of the environment that includes the deployment you want to get.")
	appconfig_getDeploymentCmd.MarkFlagRequired("application-id")
	appconfig_getDeploymentCmd.MarkFlagRequired("deployment-number")
	appconfig_getDeploymentCmd.MarkFlagRequired("environment-id")
	appconfigCmd.AddCommand(appconfig_getDeploymentCmd)
}
