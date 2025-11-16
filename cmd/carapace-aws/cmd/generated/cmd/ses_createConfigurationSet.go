package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_createConfigurationSetCmd = &cobra.Command{
	Use:   "create-configuration-set",
	Short: "Creates a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_createConfigurationSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_createConfigurationSetCmd).Standalone()

		ses_createConfigurationSetCmd.Flags().String("configuration-set", "", "A data structure that contains the name of the configuration set.")
		ses_createConfigurationSetCmd.MarkFlagRequired("configuration-set")
	})
	sesCmd.AddCommand(ses_createConfigurationSetCmd)
}
