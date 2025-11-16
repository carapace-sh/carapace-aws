package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_startDeploymentCmd = &cobra.Command{
	Use:   "start-deployment",
	Short: "Starts a deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_startDeploymentCmd).Standalone()

	appconfig_startDeploymentCmd.Flags().String("application-id", "", "The application ID.")
	appconfig_startDeploymentCmd.Flags().String("configuration-profile-id", "", "The configuration profile ID.")
	appconfig_startDeploymentCmd.Flags().String("configuration-version", "", "The configuration version to deploy.")
	appconfig_startDeploymentCmd.Flags().String("deployment-strategy-id", "", "The deployment strategy ID.")
	appconfig_startDeploymentCmd.Flags().String("description", "", "A description of the deployment.")
	appconfig_startDeploymentCmd.Flags().String("dynamic-extension-parameters", "", "A map of dynamic extension parameter names to values to pass to associated extensions with `PRE_START_DEPLOYMENT` actions.")
	appconfig_startDeploymentCmd.Flags().String("environment-id", "", "The environment ID.")
	appconfig_startDeploymentCmd.Flags().String("kms-key-identifier", "", "The KMS key identifier (key ID, key alias, or key ARN).")
	appconfig_startDeploymentCmd.Flags().String("tags", "", "Metadata to assign to the deployment.")
	appconfig_startDeploymentCmd.MarkFlagRequired("application-id")
	appconfig_startDeploymentCmd.MarkFlagRequired("configuration-profile-id")
	appconfig_startDeploymentCmd.MarkFlagRequired("configuration-version")
	appconfig_startDeploymentCmd.MarkFlagRequired("deployment-strategy-id")
	appconfig_startDeploymentCmd.MarkFlagRequired("environment-id")
	appconfigCmd.AddCommand(appconfig_startDeploymentCmd)
}
