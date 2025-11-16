package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_updateApplicationCmd = &cobra.Command{
	Use:   "update-application",
	Short: "Updates an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_updateApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_updateApplicationCmd).Standalone()

		appconfig_updateApplicationCmd.Flags().String("application-id", "", "The application ID.")
		appconfig_updateApplicationCmd.Flags().String("description", "", "A description of the application.")
		appconfig_updateApplicationCmd.Flags().String("name", "", "The name of the application.")
		appconfig_updateApplicationCmd.MarkFlagRequired("application-id")
	})
	appconfigCmd.AddCommand(appconfig_updateApplicationCmd)
}
