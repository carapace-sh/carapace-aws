package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_putRetentionConfigurationCmd = &cobra.Command{
	Use:   "put-retention-configuration",
	Short: "Creates and updates the retention configuration with details about retention period (number of days) that Config stores your historical information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_putRetentionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_putRetentionConfigurationCmd).Standalone()

		config_putRetentionConfigurationCmd.Flags().String("retention-period-in-days", "", "Number of days Config stores your historical information.")
		config_putRetentionConfigurationCmd.MarkFlagRequired("retention-period-in-days")
	})
	configCmd.AddCommand(config_putRetentionConfigurationCmd)
}
