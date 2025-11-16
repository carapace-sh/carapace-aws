package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_getMetadataTransferJobCmd = &cobra.Command{
	Use:   "get-metadata-transfer-job",
	Short: "Gets a nmetadata transfer job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_getMetadataTransferJobCmd).Standalone()

	iottwinmaker_getMetadataTransferJobCmd.Flags().String("metadata-transfer-job-id", "", "The metadata transfer job Id.")
	iottwinmaker_getMetadataTransferJobCmd.MarkFlagRequired("metadata-transfer-job-id")
	iottwinmakerCmd.AddCommand(iottwinmaker_getMetadataTransferJobCmd)
}
