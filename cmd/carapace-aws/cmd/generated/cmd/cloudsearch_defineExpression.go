package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_defineExpressionCmd = &cobra.Command{
	Use:   "define-expression",
	Short: "Configures an `Expression` for the search domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_defineExpressionCmd).Standalone()

	cloudsearch_defineExpressionCmd.Flags().String("domain-name", "", "")
	cloudsearch_defineExpressionCmd.Flags().String("expression", "", "")
	cloudsearch_defineExpressionCmd.MarkFlagRequired("domain-name")
	cloudsearch_defineExpressionCmd.MarkFlagRequired("expression")
	cloudsearchCmd.AddCommand(cloudsearch_defineExpressionCmd)
}
