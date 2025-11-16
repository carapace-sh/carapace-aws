package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_deleteExpressionCmd = &cobra.Command{
	Use:   "delete-expression",
	Short: "Removes an `Expression` from the search domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_deleteExpressionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearch_deleteExpressionCmd).Standalone()

		cloudsearch_deleteExpressionCmd.Flags().String("domain-name", "", "")
		cloudsearch_deleteExpressionCmd.Flags().String("expression-name", "", "The name of the `Expression` to delete.")
		cloudsearch_deleteExpressionCmd.MarkFlagRequired("domain-name")
		cloudsearch_deleteExpressionCmd.MarkFlagRequired("expression-name")
	})
	cloudsearchCmd.AddCommand(cloudsearch_deleteExpressionCmd)
}
