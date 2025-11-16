package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_createEnvironmentCmd = &cobra.Command{
	Use:   "create-environment",
	Short: "Creates an environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_createEnvironmentCmd).Standalone()

	appconfig_createEnvironmentCmd.Flags().String("application-id", "", "The application ID.")
	appconfig_createEnvironmentCmd.Flags().String("description", "", "A description of the environment.")
	appconfig_createEnvironmentCmd.Flags().String("monitors", "", "Amazon CloudWatch alarms to monitor during the deployment process.")
	appconfig_createEnvironmentCmd.Flags().String("name", "", "A name for the environment.")
	appconfig_createEnvironmentCmd.Flags().String("tags", "", "Metadata to assign to the environment.")
	appconfig_createEnvironmentCmd.MarkFlagRequired("application-id")
	appconfig_createEnvironmentCmd.MarkFlagRequired("name")
	appconfigCmd.AddCommand(appconfig_createEnvironmentCmd)
}
