package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_createMetadataTransferJobCmd = &cobra.Command{
	Use:   "create-metadata-transfer-job",
	Short: "Creates a new metadata transfer job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_createMetadataTransferJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_createMetadataTransferJobCmd).Standalone()

		iottwinmaker_createMetadataTransferJobCmd.Flags().String("description", "", "The metadata transfer job description.")
		iottwinmaker_createMetadataTransferJobCmd.Flags().String("destination", "", "The metadata transfer job destination.")
		iottwinmaker_createMetadataTransferJobCmd.Flags().String("metadata-transfer-job-id", "", "The metadata transfer job Id.")
		iottwinmaker_createMetadataTransferJobCmd.Flags().String("sources", "", "The metadata transfer job sources.")
		iottwinmaker_createMetadataTransferJobCmd.MarkFlagRequired("destination")
		iottwinmaker_createMetadataTransferJobCmd.MarkFlagRequired("sources")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_createMetadataTransferJobCmd)
}
