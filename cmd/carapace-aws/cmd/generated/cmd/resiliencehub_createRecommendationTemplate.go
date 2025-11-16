package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_createRecommendationTemplateCmd = &cobra.Command{
	Use:   "create-recommendation-template",
	Short: "Creates a new recommendation template for the Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_createRecommendationTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_createRecommendationTemplateCmd).Standalone()

		resiliencehub_createRecommendationTemplateCmd.Flags().String("assessment-arn", "", "Amazon Resource Name (ARN) of the assessment.")
		resiliencehub_createRecommendationTemplateCmd.Flags().String("bucket-name", "", "The name of the Amazon S3 bucket that will contain the recommendation template.")
		resiliencehub_createRecommendationTemplateCmd.Flags().String("client-token", "", "Used for an idempotency token.")
		resiliencehub_createRecommendationTemplateCmd.Flags().String("format", "", "The format for the recommendation template.")
		resiliencehub_createRecommendationTemplateCmd.Flags().String("name", "", "The name for the recommendation template.")
		resiliencehub_createRecommendationTemplateCmd.Flags().String("recommendation-ids", "", "Identifiers for the recommendations used to create a recommendation template.")
		resiliencehub_createRecommendationTemplateCmd.Flags().String("recommendation-types", "", "An array of strings that specify the recommendation template type or types.")
		resiliencehub_createRecommendationTemplateCmd.Flags().String("tags", "", "Tags assigned to the resource.")
		resiliencehub_createRecommendationTemplateCmd.MarkFlagRequired("assessment-arn")
		resiliencehub_createRecommendationTemplateCmd.MarkFlagRequired("name")
	})
	resiliencehubCmd.AddCommand(resiliencehub_createRecommendationTemplateCmd)
}
