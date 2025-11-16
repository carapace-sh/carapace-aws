package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_getEc2InstanceRecommendationsCmd = &cobra.Command{
	Use:   "get-ec2-instance-recommendations",
	Short: "Returns Amazon EC2 instance recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_getEc2InstanceRecommendationsCmd).Standalone()

	computeOptimizer_getEc2InstanceRecommendationsCmd.Flags().String("account-ids", "", "The ID of the Amazon Web Services account for which to return instance recommendations.")
	computeOptimizer_getEc2InstanceRecommendationsCmd.Flags().String("filters", "", "An array of objects to specify a filter that returns a more specific list of instance recommendations.")
	computeOptimizer_getEc2InstanceRecommendationsCmd.Flags().String("instance-arns", "", "The Amazon Resource Name (ARN) of the instances for which to return recommendations.")
	computeOptimizer_getEc2InstanceRecommendationsCmd.Flags().String("max-results", "", "The maximum number of instance recommendations to return with a single request.")
	computeOptimizer_getEc2InstanceRecommendationsCmd.Flags().String("next-token", "", "The token to advance to the next page of instance recommendations.")
	computeOptimizer_getEc2InstanceRecommendationsCmd.Flags().String("recommendation-preferences", "", "An object to specify the preferences for the Amazon EC2 instance recommendations to return in the response.")
	computeOptimizerCmd.AddCommand(computeOptimizer_getEc2InstanceRecommendationsCmd)
}
