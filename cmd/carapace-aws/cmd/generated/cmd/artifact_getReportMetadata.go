package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var artifact_getReportMetadataCmd = &cobra.Command{
	Use:   "get-report-metadata",
	Short: "Get the metadata for a single report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifact_getReportMetadataCmd).Standalone()

	artifact_getReportMetadataCmd.Flags().String("report-id", "", "Unique resource ID for the report resource.")
	artifact_getReportMetadataCmd.Flags().String("report-version", "", "Version for the report resource.")
	artifact_getReportMetadataCmd.MarkFlagRequired("report-id")
	artifactCmd.AddCommand(artifact_getReportMetadataCmd)
}
