package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_listServicesForAutoScalingConfigurationCmd = &cobra.Command{
	Use:   "list-services-for-auto-scaling-configuration",
	Short: "Returns a list of the associated App Runner services using an auto scaling configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_listServicesForAutoScalingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_listServicesForAutoScalingConfigurationCmd).Standalone()

		apprunner_listServicesForAutoScalingConfigurationCmd.Flags().String("auto-scaling-configuration-arn", "", "The Amazon Resource Name (ARN) of the App Runner auto scaling configuration that you want to list the services for.")
		apprunner_listServicesForAutoScalingConfigurationCmd.Flags().String("max-results", "", "The maximum number of results to include in each response (result page).")
		apprunner_listServicesForAutoScalingConfigurationCmd.Flags().String("next-token", "", "A token from a previous result page.")
		apprunner_listServicesForAutoScalingConfigurationCmd.MarkFlagRequired("auto-scaling-configuration-arn")
	})
	apprunnerCmd.AddCommand(apprunner_listServicesForAutoScalingConfigurationCmd)
}
