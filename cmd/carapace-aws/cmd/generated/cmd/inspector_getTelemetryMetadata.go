package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_getTelemetryMetadataCmd = &cobra.Command{
	Use:   "get-telemetry-metadata",
	Short: "Information about the data that is collected for the specified assessment run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_getTelemetryMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector_getTelemetryMetadataCmd).Standalone()

		inspector_getTelemetryMetadataCmd.Flags().String("assessment-run-arn", "", "The ARN that specifies the assessment run that has the telemetry data that you want to obtain.")
		inspector_getTelemetryMetadataCmd.MarkFlagRequired("assessment-run-arn")
	})
	inspectorCmd.AddCommand(inspector_getTelemetryMetadataCmd)
}
