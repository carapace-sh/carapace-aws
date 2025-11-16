package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_listApplicationAssociationsCmd = &cobra.Command{
	Use:   "list-application-associations",
	Short: "Returns a paginated list of application associations for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_listApplicationAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appintegrations_listApplicationAssociationsCmd).Standalone()

		appintegrations_listApplicationAssociationsCmd.Flags().String("application-id", "", "A unique identifier for the Application.")
		appintegrations_listApplicationAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		appintegrations_listApplicationAssociationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		appintegrations_listApplicationAssociationsCmd.MarkFlagRequired("application-id")
	})
	appintegrationsCmd.AddCommand(appintegrations_listApplicationAssociationsCmd)
}
