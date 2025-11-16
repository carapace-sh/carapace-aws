package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_deleteRetentionConfigurationCmd = &cobra.Command{
	Use:   "delete-retention-configuration",
	Short: "Deletes the retention configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteRetentionConfigurationCmd).Standalone()

	config_deleteRetentionConfigurationCmd.Flags().String("retention-configuration-name", "", "The name of the retention configuration to delete.")
	config_deleteRetentionConfigurationCmd.MarkFlagRequired("retention-configuration-name")
	configCmd.AddCommand(config_deleteRetentionConfigurationCmd)
}
