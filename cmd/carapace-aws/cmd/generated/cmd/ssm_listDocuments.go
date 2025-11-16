package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_listDocumentsCmd = &cobra.Command{
	Use:   "list-documents",
	Short: "Returns all Systems Manager (SSM) documents in the current Amazon Web Services account and Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_listDocumentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_listDocumentsCmd).Standalone()

		ssm_listDocumentsCmd.Flags().String("document-filter-list", "", "This data type is deprecated.")
		ssm_listDocumentsCmd.Flags().String("filters", "", "One or more `DocumentKeyValuesFilter` objects.")
		ssm_listDocumentsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_listDocumentsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	})
	ssmCmd.AddCommand(ssm_listDocumentsCmd)
}
