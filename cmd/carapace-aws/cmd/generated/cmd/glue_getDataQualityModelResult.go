package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getDataQualityModelResultCmd = &cobra.Command{
	Use:   "get-data-quality-model-result",
	Short: "Retrieve a statistic's predictions for a given Profile ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getDataQualityModelResultCmd).Standalone()

	glue_getDataQualityModelResultCmd.Flags().String("profile-id", "", "The Profile ID.")
	glue_getDataQualityModelResultCmd.Flags().String("statistic-id", "", "The Statistic ID.")
	glue_getDataQualityModelResultCmd.MarkFlagRequired("profile-id")
	glue_getDataQualityModelResultCmd.MarkFlagRequired("statistic-id")
	glueCmd.AddCommand(glue_getDataQualityModelResultCmd)
}
