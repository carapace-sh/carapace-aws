package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_createAutoScalingConfigurationCmd = &cobra.Command{
	Use:   "create-auto-scaling-configuration",
	Short: "Create an App Runner automatic scaling configuration resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_createAutoScalingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_createAutoScalingConfigurationCmd).Standalone()

		apprunner_createAutoScalingConfigurationCmd.Flags().String("auto-scaling-configuration-name", "", "A name for the auto scaling configuration.")
		apprunner_createAutoScalingConfigurationCmd.Flags().String("max-concurrency", "", "The maximum number of concurrent requests that you want an instance to process.")
		apprunner_createAutoScalingConfigurationCmd.Flags().String("max-size", "", "The maximum number of instances that your service scales up to.")
		apprunner_createAutoScalingConfigurationCmd.Flags().String("min-size", "", "The minimum number of instances that App Runner provisions for your service.")
		apprunner_createAutoScalingConfigurationCmd.Flags().String("tags", "", "A list of metadata items that you can associate with your auto scaling configuration resource.")
		apprunner_createAutoScalingConfigurationCmd.MarkFlagRequired("auto-scaling-configuration-name")
	})
	apprunnerCmd.AddCommand(apprunner_createAutoScalingConfigurationCmd)
}
