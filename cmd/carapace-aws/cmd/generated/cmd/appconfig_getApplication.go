package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_getApplicationCmd = &cobra.Command{
	Use:   "get-application",
	Short: "Retrieves information about an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_getApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_getApplicationCmd).Standalone()

		appconfig_getApplicationCmd.Flags().String("application-id", "", "The ID of the application you want to get.")
		appconfig_getApplicationCmd.MarkFlagRequired("application-id")
	})
	appconfigCmd.AddCommand(appconfig_getApplicationCmd)
}
