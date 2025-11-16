package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_describeObservabilityConfigurationCmd = &cobra.Command{
	Use:   "describe-observability-configuration",
	Short: "Return a full description of an App Runner observability configuration resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_describeObservabilityConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_describeObservabilityConfigurationCmd).Standalone()

		apprunner_describeObservabilityConfigurationCmd.Flags().String("observability-configuration-arn", "", "The Amazon Resource Name (ARN) of the App Runner observability configuration that you want a description for.")
		apprunner_describeObservabilityConfigurationCmd.MarkFlagRequired("observability-configuration-arn")
	})
	apprunnerCmd.AddCommand(apprunner_describeObservabilityConfigurationCmd)
}
