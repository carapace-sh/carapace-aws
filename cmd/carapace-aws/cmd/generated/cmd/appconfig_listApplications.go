package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_listApplicationsCmd = &cobra.Command{
	Use:   "list-applications",
	Short: "Lists all applications in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_listApplicationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_listApplicationsCmd).Standalone()

		appconfig_listApplicationsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		appconfig_listApplicationsCmd.Flags().String("next-token", "", "A token to start the list.")
	})
	appconfigCmd.AddCommand(appconfig_listApplicationsCmd)
}
