package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_listApplicationsCmd = &cobra.Command{
	Use:   "list-applications",
	Short: "Lists all OpenSearch applications under your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_listApplicationsCmd).Standalone()

	opensearch_listApplicationsCmd.Flags().String("max-results", "", "")
	opensearch_listApplicationsCmd.Flags().String("next-token", "", "")
	opensearch_listApplicationsCmd.Flags().String("statuses", "", "Filters the list of OpenSearch applications by status.")
	opensearchCmd.AddCommand(opensearch_listApplicationsCmd)
}
