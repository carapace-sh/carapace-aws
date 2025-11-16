package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getDataQualityModelCmd = &cobra.Command{
	Use:   "get-data-quality-model",
	Short: "Retrieve the training status of the model along with more information (CompletedOn, StartedOn, FailureReason).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getDataQualityModelCmd).Standalone()

	glue_getDataQualityModelCmd.Flags().String("profile-id", "", "The Profile ID.")
	glue_getDataQualityModelCmd.Flags().String("statistic-id", "", "The Statistic ID.")
	glue_getDataQualityModelCmd.MarkFlagRequired("profile-id")
	glueCmd.AddCommand(glue_getDataQualityModelCmd)
}
