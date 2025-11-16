package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_batchPutDataQualityStatisticAnnotationCmd = &cobra.Command{
	Use:   "batch-put-data-quality-statistic-annotation",
	Short: "Annotate datapoints over time for a specific data quality statistic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_batchPutDataQualityStatisticAnnotationCmd).Standalone()

	glue_batchPutDataQualityStatisticAnnotationCmd.Flags().String("client-token", "", "Client Token.")
	glue_batchPutDataQualityStatisticAnnotationCmd.Flags().String("inclusion-annotations", "", "A list of `DatapointInclusionAnnotation`'s.")
	glue_batchPutDataQualityStatisticAnnotationCmd.MarkFlagRequired("inclusion-annotations")
	glueCmd.AddCommand(glue_batchPutDataQualityStatisticAnnotationCmd)
}
