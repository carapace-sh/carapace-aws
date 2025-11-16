package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_batchGetDataQualityResultCmd = &cobra.Command{
	Use:   "batch-get-data-quality-result",
	Short: "Retrieves a list of data quality results for the specified result IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_batchGetDataQualityResultCmd).Standalone()

	glue_batchGetDataQualityResultCmd.Flags().String("result-ids", "", "A list of unique result IDs for the data quality results.")
	glue_batchGetDataQualityResultCmd.MarkFlagRequired("result-ids")
	glueCmd.AddCommand(glue_batchGetDataQualityResultCmd)
}
