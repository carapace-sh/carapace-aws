package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var artifact_getReportCmd = &cobra.Command{
	Use:   "get-report",
	Short: "Get the content for a single report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifact_getReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(artifact_getReportCmd).Standalone()

		artifact_getReportCmd.Flags().String("report-id", "", "Unique resource ID for the report resource.")
		artifact_getReportCmd.Flags().String("report-version", "", "Version for the report resource.")
		artifact_getReportCmd.Flags().String("term-token", "", "Unique download token provided by GetTermForReport API.")
		artifact_getReportCmd.MarkFlagRequired("report-id")
		artifact_getReportCmd.MarkFlagRequired("term-token")
	})
	artifactCmd.AddCommand(artifact_getReportCmd)
}
