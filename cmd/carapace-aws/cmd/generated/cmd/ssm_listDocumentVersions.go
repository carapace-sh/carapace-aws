package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_listDocumentVersionsCmd = &cobra.Command{
	Use:   "list-document-versions",
	Short: "List all versions for a document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_listDocumentVersionsCmd).Standalone()

	ssm_listDocumentVersionsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_listDocumentVersionsCmd.Flags().String("name", "", "The name of the document.")
	ssm_listDocumentVersionsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	ssm_listDocumentVersionsCmd.MarkFlagRequired("name")
	ssmCmd.AddCommand(ssm_listDocumentVersionsCmd)
}
