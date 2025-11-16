package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cur_deleteReportDefinitionCmd = &cobra.Command{
	Use:   "delete-report-definition",
	Short: "Deletes the specified report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cur_deleteReportDefinitionCmd).Standalone()

	cur_deleteReportDefinitionCmd.Flags().String("report-name", "", "The name of the report that you want to delete.")
	cur_deleteReportDefinitionCmd.MarkFlagRequired("report-name")
	curCmd.AddCommand(cur_deleteReportDefinitionCmd)
}
