package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_stopDeploymentCmd = &cobra.Command{
	Use:   "stop-deployment",
	Short: "Stops a deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_stopDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_stopDeploymentCmd).Standalone()

		appconfig_stopDeploymentCmd.Flags().Bool("allow-revert", false, "A Boolean that enables AppConfig to rollback a `COMPLETED` deployment to the previous configuration version.")
		appconfig_stopDeploymentCmd.Flags().String("application-id", "", "The application ID.")
		appconfig_stopDeploymentCmd.Flags().String("deployment-number", "", "The sequence number of the deployment.")
		appconfig_stopDeploymentCmd.Flags().String("environment-id", "", "The environment ID.")
		appconfig_stopDeploymentCmd.Flags().Bool("no-allow-revert", false, "A Boolean that enables AppConfig to rollback a `COMPLETED` deployment to the previous configuration version.")
		appconfig_stopDeploymentCmd.MarkFlagRequired("application-id")
		appconfig_stopDeploymentCmd.MarkFlagRequired("deployment-number")
		appconfig_stopDeploymentCmd.MarkFlagRequired("environment-id")
		appconfig_stopDeploymentCmd.Flag("no-allow-revert").Hidden = true
	})
	appconfigCmd.AddCommand(appconfig_stopDeploymentCmd)
}
