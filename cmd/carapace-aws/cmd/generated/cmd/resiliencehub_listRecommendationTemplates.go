package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listRecommendationTemplatesCmd = &cobra.Command{
	Use:   "list-recommendation-templates",
	Short: "Lists the recommendation templates for the Resilience Hub applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listRecommendationTemplatesCmd).Standalone()

	resiliencehub_listRecommendationTemplatesCmd.Flags().String("assessment-arn", "", "Amazon Resource Name (ARN) of the assessment.")
	resiliencehub_listRecommendationTemplatesCmd.Flags().String("max-results", "", "Maximum number of results to include in the response.")
	resiliencehub_listRecommendationTemplatesCmd.Flags().String("name", "", "The name for one of the listed recommendation templates.")
	resiliencehub_listRecommendationTemplatesCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
	resiliencehub_listRecommendationTemplatesCmd.Flags().String("recommendation-template-arn", "", "The Amazon Resource Name (ARN) for a recommendation template.")
	resiliencehub_listRecommendationTemplatesCmd.Flags().String("reverse-order", "", "The default is to sort by ascending **startTime**.")
	resiliencehub_listRecommendationTemplatesCmd.Flags().String("status", "", "Status of the action.")
	resiliencehubCmd.AddCommand(resiliencehub_listRecommendationTemplatesCmd)
}
