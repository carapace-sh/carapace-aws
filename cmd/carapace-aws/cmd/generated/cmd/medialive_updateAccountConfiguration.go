package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_updateAccountConfigurationCmd = &cobra.Command{
	Use:   "update-account-configuration",
	Short: "Update account configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_updateAccountConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_updateAccountConfigurationCmd).Standalone()

		medialive_updateAccountConfigurationCmd.Flags().String("account-configuration", "", "")
	})
	medialiveCmd.AddCommand(medialive_updateAccountConfigurationCmd)
}
