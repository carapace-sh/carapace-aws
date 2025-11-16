package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_updateDefaultAutoScalingConfigurationCmd = &cobra.Command{
	Use:   "update-default-auto-scaling-configuration",
	Short: "Update an auto scaling configuration to be the default.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_updateDefaultAutoScalingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_updateDefaultAutoScalingConfigurationCmd).Standalone()

		apprunner_updateDefaultAutoScalingConfigurationCmd.Flags().String("auto-scaling-configuration-arn", "", "The Amazon Resource Name (ARN) of the App Runner auto scaling configuration that you want to set as the default.")
		apprunner_updateDefaultAutoScalingConfigurationCmd.MarkFlagRequired("auto-scaling-configuration-arn")
	})
	apprunnerCmd.AddCommand(apprunner_updateDefaultAutoScalingConfigurationCmd)
}
