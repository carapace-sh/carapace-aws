package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describeScalingActivitiesCmd = &cobra.Command{
	Use:   "describe-scaling-activities",
	Short: "Gets information about the scaling activities in the account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describeScalingActivitiesCmd).Standalone()

	autoscaling_describeScalingActivitiesCmd.Flags().String("activity-ids", "", "The activity IDs of the desired scaling activities.")
	autoscaling_describeScalingActivitiesCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_describeScalingActivitiesCmd.Flags().String("include-deleted-groups", "", "Indicates whether to include scaling activity from deleted Auto Scaling groups.")
	autoscaling_describeScalingActivitiesCmd.Flags().String("max-records", "", "The maximum number of items to return with this call.")
	autoscaling_describeScalingActivitiesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	autoscalingCmd.AddCommand(autoscaling_describeScalingActivitiesCmd)
}
