package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeConfigurationRecordersCmd = &cobra.Command{
	Use:   "describe-configuration-recorders",
	Short: "Returns details for the configuration recorder you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeConfigurationRecordersCmd).Standalone()

	config_describeConfigurationRecordersCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the configuration recorder that you want to specify.")
	config_describeConfigurationRecordersCmd.Flags().String("configuration-recorder-names", "", "A list of names of the configuration recorders that you want to specify.")
	config_describeConfigurationRecordersCmd.Flags().String("service-principal", "", "For service-linked configuration recorders, you can use the service principal of the linked Amazon Web Services service to specify the configuration recorder.")
	configCmd.AddCommand(config_describeConfigurationRecordersCmd)
}
