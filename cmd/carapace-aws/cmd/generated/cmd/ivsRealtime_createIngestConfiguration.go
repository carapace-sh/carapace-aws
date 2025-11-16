package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_createIngestConfigurationCmd = &cobra.Command{
	Use:   "create-ingest-configuration",
	Short: "Creates a new IngestConfiguration resource, used to specify the ingest protocol for a stage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_createIngestConfigurationCmd).Standalone()

	ivsRealtime_createIngestConfigurationCmd.Flags().String("attributes", "", "Application-provided attributes to store in the IngestConfiguration and attach to a stage.")
	ivsRealtime_createIngestConfigurationCmd.Flags().String("ingest-protocol", "", "Type of ingest protocol that the user employs to broadcast.")
	ivsRealtime_createIngestConfigurationCmd.Flags().String("insecure-ingest", "", "Whether the stage allows insecure RTMP ingest.")
	ivsRealtime_createIngestConfigurationCmd.Flags().String("name", "", "Optional name that can be specified for the IngestConfiguration being created.")
	ivsRealtime_createIngestConfigurationCmd.Flags().String("stage-arn", "", "ARN of the stage with which the IngestConfiguration is associated.")
	ivsRealtime_createIngestConfigurationCmd.Flags().String("tags", "", "Tags attached to the resource.")
	ivsRealtime_createIngestConfigurationCmd.Flags().String("user-id", "", "Customer-assigned name to help identify the participant using the IngestConfiguration; this can be used to link a participant to a user in the customer’s own systems.")
	ivsRealtime_createIngestConfigurationCmd.MarkFlagRequired("ingest-protocol")
	ivsRealtimeCmd.AddCommand(ivsRealtime_createIngestConfigurationCmd)
}
