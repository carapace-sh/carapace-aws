package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_updateIngestConfigurationCmd = &cobra.Command{
	Use:   "update-ingest-configuration",
	Short: "Updates a specified IngestConfiguration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_updateIngestConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivsRealtime_updateIngestConfigurationCmd).Standalone()

		ivsRealtime_updateIngestConfigurationCmd.Flags().String("arn", "", "ARN of the IngestConfiguration, for which the related stage ARN needs to be updated.")
		ivsRealtime_updateIngestConfigurationCmd.Flags().String("stage-arn", "", "Stage ARN that needs to be updated.")
		ivsRealtime_updateIngestConfigurationCmd.MarkFlagRequired("arn")
	})
	ivsRealtimeCmd.AddCommand(ivsRealtime_updateIngestConfigurationCmd)
}
