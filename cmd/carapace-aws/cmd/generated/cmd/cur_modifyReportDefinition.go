package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cur_modifyReportDefinitionCmd = &cobra.Command{
	Use:   "modify-report-definition",
	Short: "Allows you to programmatically update your report preferences.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cur_modifyReportDefinitionCmd).Standalone()

	cur_modifyReportDefinitionCmd.Flags().String("report-definition", "", "")
	cur_modifyReportDefinitionCmd.Flags().String("report-name", "", "")
	cur_modifyReportDefinitionCmd.MarkFlagRequired("report-definition")
	cur_modifyReportDefinitionCmd.MarkFlagRequired("report-name")
	curCmd.AddCommand(cur_modifyReportDefinitionCmd)
}
