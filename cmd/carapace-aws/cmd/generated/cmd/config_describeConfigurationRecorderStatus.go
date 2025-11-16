package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeConfigurationRecorderStatusCmd = &cobra.Command{
	Use:   "describe-configuration-recorder-status",
	Short: "Returns the current status of the configuration recorder you specify as well as the status of the last recording event for the configuration recorders.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeConfigurationRecorderStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_describeConfigurationRecorderStatusCmd).Standalone()

		config_describeConfigurationRecorderStatusCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the configuration recorder that you want to specify.")
		config_describeConfigurationRecorderStatusCmd.Flags().String("configuration-recorder-names", "", "The name of the configuration recorder.")
		config_describeConfigurationRecorderStatusCmd.Flags().String("service-principal", "", "For service-linked configuration recorders, you can use the service principal of the linked Amazon Web Services service to specify the configuration recorder.")
	})
	configCmd.AddCommand(config_describeConfigurationRecorderStatusCmd)
}
