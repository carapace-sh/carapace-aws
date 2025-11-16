package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationAutoscaling_deregisterScalableTargetCmd = &cobra.Command{
	Use:   "deregister-scalable-target",
	Short: "Deregisters an Application Auto Scaling scalable target when you have finished using it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationAutoscaling_deregisterScalableTargetCmd).Standalone()

	applicationAutoscaling_deregisterScalableTargetCmd.Flags().String("resource-id", "", "The identifier of the resource associated with the scalable target.")
	applicationAutoscaling_deregisterScalableTargetCmd.Flags().String("scalable-dimension", "", "The scalable dimension associated with the scalable target.")
	applicationAutoscaling_deregisterScalableTargetCmd.Flags().String("service-namespace", "", "The namespace of the Amazon Web Services service that provides the resource.")
	applicationAutoscaling_deregisterScalableTargetCmd.MarkFlagRequired("resource-id")
	applicationAutoscaling_deregisterScalableTargetCmd.MarkFlagRequired("scalable-dimension")
	applicationAutoscaling_deregisterScalableTargetCmd.MarkFlagRequired("service-namespace")
	applicationAutoscalingCmd.AddCommand(applicationAutoscaling_deregisterScalableTargetCmd)
}
