package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmQuicksetup_getConfigurationCmd = &cobra.Command{
	Use:   "get-configuration",
	Short: "Returns details about the specified configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmQuicksetup_getConfigurationCmd).Standalone()

	ssmQuicksetup_getConfigurationCmd.Flags().String("configuration-id", "", "A service generated identifier for the configuration.")
	ssmQuicksetup_getConfigurationCmd.MarkFlagRequired("configuration-id")
	ssmQuicksetupCmd.AddCommand(ssmQuicksetup_getConfigurationCmd)
}
