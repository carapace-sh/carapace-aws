package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivschat_deleteLoggingConfigurationCmd = &cobra.Command{
	Use:   "delete-logging-configuration",
	Short: "Deletes the specified logging configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivschat_deleteLoggingConfigurationCmd).Standalone()

	ivschat_deleteLoggingConfigurationCmd.Flags().String("identifier", "", "Identifier of the logging configuration to be deleted.")
	ivschat_deleteLoggingConfigurationCmd.MarkFlagRequired("identifier")
	ivschatCmd.AddCommand(ivschat_deleteLoggingConfigurationCmd)
}
