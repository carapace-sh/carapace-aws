package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_listApplicationsCmd = &cobra.Command{
	Use:   "list-applications",
	Short: "Lists applications in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_listApplicationsCmd).Standalone()

	appintegrations_listApplicationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	appintegrations_listApplicationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	appintegrationsCmd.AddCommand(appintegrations_listApplicationsCmd)
}
