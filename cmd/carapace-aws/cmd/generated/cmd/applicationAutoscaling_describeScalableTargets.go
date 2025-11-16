package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationAutoscaling_describeScalableTargetsCmd = &cobra.Command{
	Use:   "describe-scalable-targets",
	Short: "Gets information about the scalable targets in the specified namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationAutoscaling_describeScalableTargetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationAutoscaling_describeScalableTargetsCmd).Standalone()

		applicationAutoscaling_describeScalableTargetsCmd.Flags().String("max-results", "", "The maximum number of scalable targets.")
		applicationAutoscaling_describeScalableTargetsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		applicationAutoscaling_describeScalableTargetsCmd.Flags().String("resource-ids", "", "The identifier of the resource associated with the scalable target.")
		applicationAutoscaling_describeScalableTargetsCmd.Flags().String("scalable-dimension", "", "The scalable dimension associated with the scalable target.")
		applicationAutoscaling_describeScalableTargetsCmd.Flags().String("service-namespace", "", "The namespace of the Amazon Web Services service that provides the resource.")
		applicationAutoscaling_describeScalableTargetsCmd.MarkFlagRequired("service-namespace")
	})
	applicationAutoscalingCmd.AddCommand(applicationAutoscaling_describeScalableTargetsCmd)
}
