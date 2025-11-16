package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getConfigurationSetCmd = &cobra.Command{
	Use:   "get-configuration-set",
	Short: "Get information about an existing configuration set, including the dedicated IP pool that it's associated with, whether or not it's enabled for sending email, and more.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getConfigurationSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_getConfigurationSetCmd).Standalone()

		sesv2_getConfigurationSetCmd.Flags().String("configuration-set-name", "", "The name of the configuration set.")
		sesv2_getConfigurationSetCmd.MarkFlagRequired("configuration-set-name")
	})
	sesv2Cmd.AddCommand(sesv2_getConfigurationSetCmd)
}
