package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_describeAnalysisSchemesCmd = &cobra.Command{
	Use:   "describe-analysis-schemes",
	Short: "Gets the analysis schemes configured for a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_describeAnalysisSchemesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearch_describeAnalysisSchemesCmd).Standalone()

		cloudsearch_describeAnalysisSchemesCmd.Flags().String("analysis-scheme-names", "", "The analysis schemes you want to describe.")
		cloudsearch_describeAnalysisSchemesCmd.Flags().Bool("deployed", false, "Whether to display the deployed configuration (`true`) or include any pending changes (`false`).")
		cloudsearch_describeAnalysisSchemesCmd.Flags().String("domain-name", "", "The name of the domain you want to describe.")
		cloudsearch_describeAnalysisSchemesCmd.Flags().Bool("no-deployed", false, "Whether to display the deployed configuration (`true`) or include any pending changes (`false`).")
		cloudsearch_describeAnalysisSchemesCmd.MarkFlagRequired("domain-name")
		cloudsearch_describeAnalysisSchemesCmd.Flag("no-deployed").Hidden = true
	})
	cloudsearchCmd.AddCommand(cloudsearch_describeAnalysisSchemesCmd)
}
