package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_deleteConfigurationSetCmd = &cobra.Command{
	Use:   "delete-configuration-set",
	Short: "Delete an existing configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_deleteConfigurationSetCmd).Standalone()

	sesv2_deleteConfigurationSetCmd.Flags().String("configuration-set-name", "", "The name of the configuration set.")
	sesv2_deleteConfigurationSetCmd.MarkFlagRequired("configuration-set-name")
	sesv2Cmd.AddCommand(sesv2_deleteConfigurationSetCmd)
}
