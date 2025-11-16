package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createIngestionCmd = &cobra.Command{
	Use:   "create-ingestion",
	Short: "Creates and starts a new SPICE ingestion for a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createIngestionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_createIngestionCmd).Standalone()

		quicksight_createIngestionCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
		quicksight_createIngestionCmd.Flags().String("data-set-id", "", "The ID of the dataset used in the ingestion.")
		quicksight_createIngestionCmd.Flags().String("ingestion-id", "", "An ID for the ingestion.")
		quicksight_createIngestionCmd.Flags().String("ingestion-type", "", "The type of ingestion that you want to create.")
		quicksight_createIngestionCmd.MarkFlagRequired("aws-account-id")
		quicksight_createIngestionCmd.MarkFlagRequired("data-set-id")
		quicksight_createIngestionCmd.MarkFlagRequired("ingestion-id")
	})
	quicksightCmd.AddCommand(quicksight_createIngestionCmd)
}
