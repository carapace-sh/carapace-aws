package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeIngestionCmd = &cobra.Command{
	Use:   "describe-ingestion",
	Short: "Describes a SPICE ingestion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeIngestionCmd).Standalone()

	quicksight_describeIngestionCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
	quicksight_describeIngestionCmd.Flags().String("data-set-id", "", "The ID of the dataset used in the ingestion.")
	quicksight_describeIngestionCmd.Flags().String("ingestion-id", "", "An ID for the ingestion.")
	quicksight_describeIngestionCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeIngestionCmd.MarkFlagRequired("data-set-id")
	quicksight_describeIngestionCmd.MarkFlagRequired("ingestion-id")
	quicksightCmd.AddCommand(quicksight_describeIngestionCmd)
}
