package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeScalingPoliciesCmd = &cobra.Command{
	Use:   "describe-scaling-policies",
	Short: "**This API works with the following fleet types:** EC2\n\nRetrieves all scaling policies applied to a fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeScalingPoliciesCmd).Standalone()

	gamelift_describeScalingPoliciesCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet for which to retrieve scaling policies.")
	gamelift_describeScalingPoliciesCmd.Flags().String("limit", "", "The maximum number of results to return.")
	gamelift_describeScalingPoliciesCmd.Flags().String("location", "", "The fleet location.")
	gamelift_describeScalingPoliciesCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
	gamelift_describeScalingPoliciesCmd.Flags().String("status-filter", "", "Scaling policy status to filter results on.")
	gamelift_describeScalingPoliciesCmd.MarkFlagRequired("fleet-id")
	gameliftCmd.AddCommand(gamelift_describeScalingPoliciesCmd)
}
