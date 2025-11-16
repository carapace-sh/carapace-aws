package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmQuicksetup_deleteConfigurationManagerCmd = &cobra.Command{
	Use:   "delete-configuration-manager",
	Short: "Deletes a configuration manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmQuicksetup_deleteConfigurationManagerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmQuicksetup_deleteConfigurationManagerCmd).Standalone()

		ssmQuicksetup_deleteConfigurationManagerCmd.Flags().String("manager-arn", "", "The ID of the configuration manager.")
		ssmQuicksetup_deleteConfigurationManagerCmd.MarkFlagRequired("manager-arn")
	})
	ssmQuicksetupCmd.AddCommand(ssmQuicksetup_deleteConfigurationManagerCmd)
}
