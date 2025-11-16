package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_cancelMetadataTransferJobCmd = &cobra.Command{
	Use:   "cancel-metadata-transfer-job",
	Short: "Cancels the metadata transfer job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_cancelMetadataTransferJobCmd).Standalone()

	iottwinmaker_cancelMetadataTransferJobCmd.Flags().String("metadata-transfer-job-id", "", "The metadata transfer job Id.")
	iottwinmaker_cancelMetadataTransferJobCmd.MarkFlagRequired("metadata-transfer-job-id")
	iottwinmakerCmd.AddCommand(iottwinmaker_cancelMetadataTransferJobCmd)
}
