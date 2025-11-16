package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_deleteConfigurationRecorderCmd = &cobra.Command{
	Use:   "delete-configuration-recorder",
	Short: "Deletes the customer managed configuration recorder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteConfigurationRecorderCmd).Standalone()

	config_deleteConfigurationRecorderCmd.Flags().String("configuration-recorder-name", "", "The name of the customer managed configuration recorder that you want to delete.")
	config_deleteConfigurationRecorderCmd.MarkFlagRequired("configuration-recorder-name")
	configCmd.AddCommand(config_deleteConfigurationRecorderCmd)
}
