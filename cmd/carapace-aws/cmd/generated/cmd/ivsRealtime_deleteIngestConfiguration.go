package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_deleteIngestConfigurationCmd = &cobra.Command{
	Use:   "delete-ingest-configuration",
	Short: "Deletes a specified IngestConfiguration, so it can no longer be used to broadcast.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_deleteIngestConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivsRealtime_deleteIngestConfigurationCmd).Standalone()

		ivsRealtime_deleteIngestConfigurationCmd.Flags().String("arn", "", "ARN of the IngestConfiguration.")
		ivsRealtime_deleteIngestConfigurationCmd.Flags().Bool("force", false, "Optional field to force deletion of the IngestConfiguration.")
		ivsRealtime_deleteIngestConfigurationCmd.Flags().Bool("no-force", false, "Optional field to force deletion of the IngestConfiguration.")
		ivsRealtime_deleteIngestConfigurationCmd.MarkFlagRequired("arn")
		ivsRealtime_deleteIngestConfigurationCmd.Flag("no-force").Hidden = true
	})
	ivsRealtimeCmd.AddCommand(ivsRealtime_deleteIngestConfigurationCmd)
}
