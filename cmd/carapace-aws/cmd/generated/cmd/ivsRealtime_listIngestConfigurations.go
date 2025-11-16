package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_listIngestConfigurationsCmd = &cobra.Command{
	Use:   "list-ingest-configurations",
	Short: "Lists all IngestConfigurations in your account, in the AWS region where the API request is processed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_listIngestConfigurationsCmd).Standalone()

	ivsRealtime_listIngestConfigurationsCmd.Flags().String("filter-by-stage-arn", "", "Filters the response list to match the specified stage ARN.")
	ivsRealtime_listIngestConfigurationsCmd.Flags().String("filter-by-state", "", "Filters the response list to match the specified state.")
	ivsRealtime_listIngestConfigurationsCmd.Flags().String("max-results", "", "Maximum number of results to return.")
	ivsRealtime_listIngestConfigurationsCmd.Flags().String("next-token", "", "The first IngestConfiguration to retrieve.")
	ivsRealtimeCmd.AddCommand(ivsRealtime_listIngestConfigurationsCmd)
}
