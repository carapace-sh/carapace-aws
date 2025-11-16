package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_createFindingsReportCmd = &cobra.Command{
	Use:   "create-findings-report",
	Short: "Creates a finding report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_createFindingsReportCmd).Standalone()

	inspector2_createFindingsReportCmd.Flags().String("filter-criteria", "", "The filter criteria to apply to the results of the finding report.")
	inspector2_createFindingsReportCmd.Flags().String("report-format", "", "The format to generate the report in.")
	inspector2_createFindingsReportCmd.Flags().String("s3-destination", "", "The Amazon S3 export destination for the report.")
	inspector2_createFindingsReportCmd.MarkFlagRequired("report-format")
	inspector2_createFindingsReportCmd.MarkFlagRequired("s3-destination")
	inspector2Cmd.AddCommand(inspector2_createFindingsReportCmd)
}
