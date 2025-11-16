package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cur_putReportDefinitionCmd = &cobra.Command{
	Use:   "put-report-definition",
	Short: "Creates a new report using the description that you provide.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cur_putReportDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cur_putReportDefinitionCmd).Standalone()

		cur_putReportDefinitionCmd.Flags().String("report-definition", "", "Represents the output of the PutReportDefinition operation.")
		cur_putReportDefinitionCmd.Flags().String("tags", "", "The tags to be assigned to the report definition resource.")
		cur_putReportDefinitionCmd.MarkFlagRequired("report-definition")
	})
	curCmd.AddCommand(cur_putReportDefinitionCmd)
}
