package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_getEnvironmentCmd = &cobra.Command{
	Use:   "get-environment",
	Short: "Retrieves information about an environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_getEnvironmentCmd).Standalone()

	appconfig_getEnvironmentCmd.Flags().String("application-id", "", "The ID of the application that includes the environment you want to get.")
	appconfig_getEnvironmentCmd.Flags().String("environment-id", "", "The ID of the environment that you want to get.")
	appconfig_getEnvironmentCmd.MarkFlagRequired("application-id")
	appconfig_getEnvironmentCmd.MarkFlagRequired("environment-id")
	appconfigCmd.AddCommand(appconfig_getEnvironmentCmd)
}
