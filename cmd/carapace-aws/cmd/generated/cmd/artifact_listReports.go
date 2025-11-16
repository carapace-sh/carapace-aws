package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var artifact_listReportsCmd = &cobra.Command{
	Use:   "list-reports",
	Short: "List available reports.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifact_listReportsCmd).Standalone()

	artifact_listReportsCmd.Flags().String("max-results", "", "Maximum number of resources to return in the paginated response.")
	artifact_listReportsCmd.Flags().String("next-token", "", "Pagination token to request the next page of resources.")
	artifactCmd.AddCommand(artifact_listReportsCmd)
}
