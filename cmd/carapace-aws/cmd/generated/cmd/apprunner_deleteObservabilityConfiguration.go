package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_deleteObservabilityConfigurationCmd = &cobra.Command{
	Use:   "delete-observability-configuration",
	Short: "Delete an App Runner observability configuration resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_deleteObservabilityConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_deleteObservabilityConfigurationCmd).Standalone()

		apprunner_deleteObservabilityConfigurationCmd.Flags().String("observability-configuration-arn", "", "The Amazon Resource Name (ARN) of the App Runner observability configuration that you want to delete.")
		apprunner_deleteObservabilityConfigurationCmd.MarkFlagRequired("observability-configuration-arn")
	})
	apprunnerCmd.AddCommand(apprunner_deleteObservabilityConfigurationCmd)
}
