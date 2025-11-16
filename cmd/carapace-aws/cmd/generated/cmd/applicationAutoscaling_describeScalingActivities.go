package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationAutoscaling_describeScalingActivitiesCmd = &cobra.Command{
	Use:   "describe-scaling-activities",
	Short: "Provides descriptive information about the scaling activities in the specified namespace from the previous six weeks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationAutoscaling_describeScalingActivitiesCmd).Standalone()

	applicationAutoscaling_describeScalingActivitiesCmd.Flags().String("include-not-scaled-activities", "", "Specifies whether to include activities that aren't scaled (*not scaled activities*) in the response.")
	applicationAutoscaling_describeScalingActivitiesCmd.Flags().String("max-results", "", "The maximum number of scalable targets.")
	applicationAutoscaling_describeScalingActivitiesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	applicationAutoscaling_describeScalingActivitiesCmd.Flags().String("resource-id", "", "The identifier of the resource associated with the scaling activity.")
	applicationAutoscaling_describeScalingActivitiesCmd.Flags().String("scalable-dimension", "", "The scalable dimension.")
	applicationAutoscaling_describeScalingActivitiesCmd.Flags().String("service-namespace", "", "The namespace of the Amazon Web Services service that provides the resource.")
	applicationAutoscaling_describeScalingActivitiesCmd.MarkFlagRequired("service-namespace")
	applicationAutoscalingCmd.AddCommand(applicationAutoscaling_describeScalingActivitiesCmd)
}
