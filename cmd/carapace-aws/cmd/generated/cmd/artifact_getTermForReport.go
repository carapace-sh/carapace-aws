package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var artifact_getTermForReportCmd = &cobra.Command{
	Use:   "get-term-for-report",
	Short: "Get the Term content associated with a single report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifact_getTermForReportCmd).Standalone()

	artifact_getTermForReportCmd.Flags().String("report-id", "", "Unique resource ID for the report resource.")
	artifact_getTermForReportCmd.Flags().String("report-version", "", "Version for the report resource.")
	artifact_getTermForReportCmd.MarkFlagRequired("report-id")
	artifactCmd.AddCommand(artifact_getTermForReportCmd)
}
