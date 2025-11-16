package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_getEncryptionConfigurationCmd = &cobra.Command{
	Use:   "get-encryption-configuration",
	Short: "Retrieves the encryption configuration for resources and data in Amazon Web Services IoT FleetWise.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_getEncryptionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_getEncryptionConfigurationCmd).Standalone()

	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_getEncryptionConfigurationCmd)
}
