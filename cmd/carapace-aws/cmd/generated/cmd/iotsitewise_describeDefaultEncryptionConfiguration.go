package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeDefaultEncryptionConfigurationCmd = &cobra.Command{
	Use:   "describe-default-encryption-configuration",
	Short: "Retrieves information about the default encryption configuration for the Amazon Web Services account in the default or specified Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeDefaultEncryptionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_describeDefaultEncryptionConfigurationCmd).Standalone()

	})
	iotsitewiseCmd.AddCommand(iotsitewise_describeDefaultEncryptionConfigurationCmd)
}
