package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_cancelIngestionCmd = &cobra.Command{
	Use:   "cancel-ingestion",
	Short: "Cancels an ongoing ingestion of data into SPICE.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_cancelIngestionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_cancelIngestionCmd).Standalone()

		quicksight_cancelIngestionCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
		quicksight_cancelIngestionCmd.Flags().String("data-set-id", "", "The ID of the dataset used in the ingestion.")
		quicksight_cancelIngestionCmd.Flags().String("ingestion-id", "", "An ID for the ingestion.")
		quicksight_cancelIngestionCmd.MarkFlagRequired("aws-account-id")
		quicksight_cancelIngestionCmd.MarkFlagRequired("data-set-id")
		quicksight_cancelIngestionCmd.MarkFlagRequired("ingestion-id")
	})
	quicksightCmd.AddCommand(quicksight_cancelIngestionCmd)
}
