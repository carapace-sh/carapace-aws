package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_deleteIndexFieldCmd = &cobra.Command{
	Use:   "delete-index-field",
	Short: "Removes an `IndexField` from the search domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_deleteIndexFieldCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearch_deleteIndexFieldCmd).Standalone()

		cloudsearch_deleteIndexFieldCmd.Flags().String("domain-name", "", "")
		cloudsearch_deleteIndexFieldCmd.Flags().String("index-field-name", "", "The name of the index field your want to remove from the domain's indexing options.")
		cloudsearch_deleteIndexFieldCmd.MarkFlagRequired("domain-name")
		cloudsearch_deleteIndexFieldCmd.MarkFlagRequired("index-field-name")
	})
	cloudsearchCmd.AddCommand(cloudsearch_deleteIndexFieldCmd)
}
