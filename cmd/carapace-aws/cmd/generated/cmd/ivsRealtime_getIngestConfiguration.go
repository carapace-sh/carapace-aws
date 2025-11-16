package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_getIngestConfigurationCmd = &cobra.Command{
	Use:   "get-ingest-configuration",
	Short: "Gets information about the specified IngestConfiguration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_getIngestConfigurationCmd).Standalone()

	ivsRealtime_getIngestConfigurationCmd.Flags().String("arn", "", "ARN of the ingest for which the information is to be retrieved.")
	ivsRealtime_getIngestConfigurationCmd.MarkFlagRequired("arn")
	ivsRealtimeCmd.AddCommand(ivsRealtime_getIngestConfigurationCmd)
}
