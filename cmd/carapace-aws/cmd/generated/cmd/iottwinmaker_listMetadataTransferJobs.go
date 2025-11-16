package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_listMetadataTransferJobsCmd = &cobra.Command{
	Use:   "list-metadata-transfer-jobs",
	Short: "Lists the metadata transfer jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_listMetadataTransferJobsCmd).Standalone()

	iottwinmaker_listMetadataTransferJobsCmd.Flags().String("destination-type", "", "The metadata transfer job's destination type.")
	iottwinmaker_listMetadataTransferJobsCmd.Flags().String("filters", "", "An object that filters metadata transfer jobs.")
	iottwinmaker_listMetadataTransferJobsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iottwinmaker_listMetadataTransferJobsCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
	iottwinmaker_listMetadataTransferJobsCmd.Flags().String("source-type", "", "The metadata transfer job's source type.")
	iottwinmaker_listMetadataTransferJobsCmd.MarkFlagRequired("destination-type")
	iottwinmaker_listMetadataTransferJobsCmd.MarkFlagRequired("source-type")
	iottwinmakerCmd.AddCommand(iottwinmaker_listMetadataTransferJobsCmd)
}
