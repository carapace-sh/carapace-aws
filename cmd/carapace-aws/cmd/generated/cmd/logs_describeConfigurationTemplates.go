package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_describeConfigurationTemplatesCmd = &cobra.Command{
	Use:   "describe-configuration-templates",
	Short: "Use this operation to return the valid and default values that are used when creating delivery sources, delivery destinations, and deliveries.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_describeConfigurationTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_describeConfigurationTemplatesCmd).Standalone()

		logs_describeConfigurationTemplatesCmd.Flags().String("delivery-destination-types", "", "Use this parameter to filter the response to include only the configuration templates that apply to the delivery destination types that you specify here.")
		logs_describeConfigurationTemplatesCmd.Flags().String("limit", "", "Use this parameter to limit the number of configuration templates that are returned in the response.")
		logs_describeConfigurationTemplatesCmd.Flags().String("log-types", "", "Use this parameter to filter the response to include only the configuration templates that apply to the log types that you specify here.")
		logs_describeConfigurationTemplatesCmd.Flags().String("next-token", "", "")
		logs_describeConfigurationTemplatesCmd.Flags().String("resource-types", "", "Use this parameter to filter the response to include only the configuration templates that apply to the resource types that you specify here.")
		logs_describeConfigurationTemplatesCmd.Flags().String("service", "", "Use this parameter to filter the response to include only the configuration templates that apply to the Amazon Web Services service that you specify here.")
	})
	logsCmd.AddCommand(logs_describeConfigurationTemplatesCmd)
}
