package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getDataQualityResultCmd = &cobra.Command{
	Use:   "get-data-quality-result",
	Short: "Retrieves the result of a data quality rule evaluation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getDataQualityResultCmd).Standalone()

	glue_getDataQualityResultCmd.Flags().String("result-id", "", "A unique result ID for the data quality result.")
	glue_getDataQualityResultCmd.MarkFlagRequired("result-id")
	glueCmd.AddCommand(glue_getDataQualityResultCmd)
}
