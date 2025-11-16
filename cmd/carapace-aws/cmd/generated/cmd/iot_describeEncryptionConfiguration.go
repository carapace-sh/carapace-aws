package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeEncryptionConfigurationCmd = &cobra.Command{
	Use:   "describe-encryption-configuration",
	Short: "Retrieves the encryption configuration for resources and data of your Amazon Web Services account in Amazon Web Services IoT Core.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeEncryptionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeEncryptionConfigurationCmd).Standalone()

	})
	iotCmd.AddCommand(iot_describeEncryptionConfigurationCmd)
}
