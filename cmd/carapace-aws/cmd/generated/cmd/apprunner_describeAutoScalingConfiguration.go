package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_describeAutoScalingConfigurationCmd = &cobra.Command{
	Use:   "describe-auto-scaling-configuration",
	Short: "Return a full description of an App Runner automatic scaling configuration resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_describeAutoScalingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_describeAutoScalingConfigurationCmd).Standalone()

		apprunner_describeAutoScalingConfigurationCmd.Flags().String("auto-scaling-configuration-arn", "", "The Amazon Resource Name (ARN) of the App Runner auto scaling configuration that you want a description for.")
		apprunner_describeAutoScalingConfigurationCmd.MarkFlagRequired("auto-scaling-configuration-arn")
	})
	apprunnerCmd.AddCommand(apprunner_describeAutoScalingConfigurationCmd)
}
