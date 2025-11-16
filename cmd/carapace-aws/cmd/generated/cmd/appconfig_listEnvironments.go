package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_listEnvironmentsCmd = &cobra.Command{
	Use:   "list-environments",
	Short: "Lists the environments for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_listEnvironmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_listEnvironmentsCmd).Standalone()

		appconfig_listEnvironmentsCmd.Flags().String("application-id", "", "The application ID.")
		appconfig_listEnvironmentsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		appconfig_listEnvironmentsCmd.Flags().String("next-token", "", "A token to start the list.")
		appconfig_listEnvironmentsCmd.MarkFlagRequired("application-id")
	})
	appconfigCmd.AddCommand(appconfig_listEnvironmentsCmd)
}
