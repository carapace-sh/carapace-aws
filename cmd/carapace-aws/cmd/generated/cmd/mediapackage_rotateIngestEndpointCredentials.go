package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackage_rotateIngestEndpointCredentialsCmd = &cobra.Command{
	Use:   "rotate-ingest-endpoint-credentials",
	Short: "Rotate the IngestEndpoint's username and password, as specified by the IngestEndpoint's id.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackage_rotateIngestEndpointCredentialsCmd).Standalone()

	mediapackage_rotateIngestEndpointCredentialsCmd.Flags().String("id", "", "The ID of the channel the IngestEndpoint is on.")
	mediapackage_rotateIngestEndpointCredentialsCmd.Flags().String("ingest-endpoint-id", "", "The id of the IngestEndpoint whose credentials should be rotated")
	mediapackage_rotateIngestEndpointCredentialsCmd.MarkFlagRequired("id")
	mediapackage_rotateIngestEndpointCredentialsCmd.MarkFlagRequired("ingest-endpoint-id")
	mediapackageCmd.AddCommand(mediapackage_rotateIngestEndpointCredentialsCmd)
}
