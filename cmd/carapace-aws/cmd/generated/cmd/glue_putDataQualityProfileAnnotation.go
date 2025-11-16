package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_putDataQualityProfileAnnotationCmd = &cobra.Command{
	Use:   "put-data-quality-profile-annotation",
	Short: "Annotate all datapoints for a Profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_putDataQualityProfileAnnotationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_putDataQualityProfileAnnotationCmd).Standalone()

		glue_putDataQualityProfileAnnotationCmd.Flags().String("inclusion-annotation", "", "The inclusion annotation value to apply to the profile.")
		glue_putDataQualityProfileAnnotationCmd.Flags().String("profile-id", "", "The ID of the data quality monitoring profile to annotate.")
		glue_putDataQualityProfileAnnotationCmd.MarkFlagRequired("inclusion-annotation")
		glue_putDataQualityProfileAnnotationCmd.MarkFlagRequired("profile-id")
	})
	glueCmd.AddCommand(glue_putDataQualityProfileAnnotationCmd)
}
