package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_describeExpressionsCmd = &cobra.Command{
	Use:   "describe-expressions",
	Short: "Gets the expressions configured for the search domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_describeExpressionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearch_describeExpressionsCmd).Standalone()

		cloudsearch_describeExpressionsCmd.Flags().Bool("deployed", false, "Whether to display the deployed configuration (`true`) or include any pending changes (`false`).")
		cloudsearch_describeExpressionsCmd.Flags().String("domain-name", "", "The name of the domain you want to describe.")
		cloudsearch_describeExpressionsCmd.Flags().String("expression-names", "", "Limits the `DescribeExpressions` response to the specified expressions.")
		cloudsearch_describeExpressionsCmd.Flags().Bool("no-deployed", false, "Whether to display the deployed configuration (`true`) or include any pending changes (`false`).")
		cloudsearch_describeExpressionsCmd.MarkFlagRequired("domain-name")
		cloudsearch_describeExpressionsCmd.Flag("no-deployed").Hidden = true
	})
	cloudsearchCmd.AddCommand(cloudsearch_describeExpressionsCmd)
}
