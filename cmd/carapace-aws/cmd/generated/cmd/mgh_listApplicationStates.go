package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_listApplicationStatesCmd = &cobra.Command{
	Use:   "list-application-states",
	Short: "Lists all the migration statuses for your applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_listApplicationStatesCmd).Standalone()

	mgh_listApplicationStatesCmd.Flags().String("application-ids", "", "The configurationIds from the Application Discovery Service that uniquely identifies your applications.")
	mgh_listApplicationStatesCmd.Flags().String("max-results", "", "Maximum number of results to be returned per page.")
	mgh_listApplicationStatesCmd.Flags().String("next-token", "", "If a `NextToken` was returned by a previous call, there are more results available.")
	mghCmd.AddCommand(mgh_listApplicationStatesCmd)
}
