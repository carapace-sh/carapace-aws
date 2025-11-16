package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_deleteConfigurationSetCmd = &cobra.Command{
	Use:   "delete-configuration-set",
	Short: "Deletes a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_deleteConfigurationSetCmd).Standalone()

	ses_deleteConfigurationSetCmd.Flags().String("configuration-set-name", "", "The name of the configuration set to delete.")
	ses_deleteConfigurationSetCmd.MarkFlagRequired("configuration-set-name")
	sesCmd.AddCommand(ses_deleteConfigurationSetCmd)
}
