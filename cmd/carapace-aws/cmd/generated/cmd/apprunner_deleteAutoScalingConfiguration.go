package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_deleteAutoScalingConfigurationCmd = &cobra.Command{
	Use:   "delete-auto-scaling-configuration",
	Short: "Delete an App Runner automatic scaling configuration resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_deleteAutoScalingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_deleteAutoScalingConfigurationCmd).Standalone()

		apprunner_deleteAutoScalingConfigurationCmd.Flags().String("auto-scaling-configuration-arn", "", "The Amazon Resource Name (ARN) of the App Runner auto scaling configuration that you want to delete.")
		apprunner_deleteAutoScalingConfigurationCmd.Flags().Bool("delete-all-revisions", false, "Set to `true` to delete all of the revisions associated with the `AutoScalingConfigurationArn` parameter value.")
		apprunner_deleteAutoScalingConfigurationCmd.Flags().Bool("no-delete-all-revisions", false, "Set to `true` to delete all of the revisions associated with the `AutoScalingConfigurationArn` parameter value.")
		apprunner_deleteAutoScalingConfigurationCmd.MarkFlagRequired("auto-scaling-configuration-arn")
		apprunner_deleteAutoScalingConfigurationCmd.Flag("no-delete-all-revisions").Hidden = true
	})
	apprunnerCmd.AddCommand(apprunner_deleteAutoScalingConfigurationCmd)
}
