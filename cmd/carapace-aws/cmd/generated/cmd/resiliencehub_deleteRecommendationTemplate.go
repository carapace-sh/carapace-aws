package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_deleteRecommendationTemplateCmd = &cobra.Command{
	Use:   "delete-recommendation-template",
	Short: "Deletes a recommendation template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_deleteRecommendationTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_deleteRecommendationTemplateCmd).Standalone()

		resiliencehub_deleteRecommendationTemplateCmd.Flags().String("client-token", "", "Used for an idempotency token.")
		resiliencehub_deleteRecommendationTemplateCmd.Flags().String("recommendation-template-arn", "", "The Amazon Resource Name (ARN) for a recommendation template.")
		resiliencehub_deleteRecommendationTemplateCmd.MarkFlagRequired("recommendation-template-arn")
	})
	resiliencehubCmd.AddCommand(resiliencehub_deleteRecommendationTemplateCmd)
}
