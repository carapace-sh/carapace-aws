package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationAutoscaling_registerScalableTargetCmd = &cobra.Command{
	Use:   "register-scalable-target",
	Short: "Registers or updates a scalable target, which is the resource that you want to scale.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationAutoscaling_registerScalableTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationAutoscaling_registerScalableTargetCmd).Standalone()

		applicationAutoscaling_registerScalableTargetCmd.Flags().String("max-capacity", "", "The maximum value that you plan to scale out to.")
		applicationAutoscaling_registerScalableTargetCmd.Flags().String("min-capacity", "", "The minimum value that you plan to scale in to.")
		applicationAutoscaling_registerScalableTargetCmd.Flags().String("resource-id", "", "The identifier of the resource that is associated with the scalable target.")
		applicationAutoscaling_registerScalableTargetCmd.Flags().String("role-arn", "", "This parameter is required for services that do not support service-linked roles (such as Amazon EMR), and it must specify the ARN of an IAM role that allows Application Auto Scaling to modify the scalable target on your behalf.")
		applicationAutoscaling_registerScalableTargetCmd.Flags().String("scalable-dimension", "", "The scalable dimension associated with the scalable target.")
		applicationAutoscaling_registerScalableTargetCmd.Flags().String("service-namespace", "", "The namespace of the Amazon Web Services service that provides the resource.")
		applicationAutoscaling_registerScalableTargetCmd.Flags().String("suspended-state", "", "An embedded object that contains attributes and attribute values that are used to suspend and resume automatic scaling.")
		applicationAutoscaling_registerScalableTargetCmd.Flags().String("tags", "", "Assigns one or more tags to the scalable target.")
		applicationAutoscaling_registerScalableTargetCmd.MarkFlagRequired("resource-id")
		applicationAutoscaling_registerScalableTargetCmd.MarkFlagRequired("scalable-dimension")
		applicationAutoscaling_registerScalableTargetCmd.MarkFlagRequired("service-namespace")
	})
	applicationAutoscalingCmd.AddCommand(applicationAutoscaling_registerScalableTargetCmd)
}
