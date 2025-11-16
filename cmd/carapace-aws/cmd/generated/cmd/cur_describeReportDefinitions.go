package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cur_describeReportDefinitionsCmd = &cobra.Command{
	Use:   "describe-report-definitions",
	Short: "Lists the Amazon Web Services Cost and Usage Report available to this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cur_describeReportDefinitionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cur_describeReportDefinitionsCmd).Standalone()

		cur_describeReportDefinitionsCmd.Flags().String("max-results", "", "")
		cur_describeReportDefinitionsCmd.Flags().String("next-token", "", "")
	})
	curCmd.AddCommand(cur_describeReportDefinitionsCmd)
}
