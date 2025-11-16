package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_createObservabilityConfigurationCmd = &cobra.Command{
	Use:   "create-observability-configuration",
	Short: "Create an App Runner observability configuration resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_createObservabilityConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_createObservabilityConfigurationCmd).Standalone()

		apprunner_createObservabilityConfigurationCmd.Flags().String("observability-configuration-name", "", "A name for the observability configuration.")
		apprunner_createObservabilityConfigurationCmd.Flags().String("tags", "", "A list of metadata items that you can associate with your observability configuration resource.")
		apprunner_createObservabilityConfigurationCmd.Flags().String("trace-configuration", "", "The configuration of the tracing feature within this observability configuration.")
		apprunner_createObservabilityConfigurationCmd.MarkFlagRequired("observability-configuration-name")
	})
	apprunnerCmd.AddCommand(apprunner_createObservabilityConfigurationCmd)
}
