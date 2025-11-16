package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmQuicksetup_updateConfigurationManagerCmd = &cobra.Command{
	Use:   "update-configuration-manager",
	Short: "Updates a Quick Setup configuration manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmQuicksetup_updateConfigurationManagerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmQuicksetup_updateConfigurationManagerCmd).Standalone()

		ssmQuicksetup_updateConfigurationManagerCmd.Flags().String("description", "", "A description of the configuration manager.")
		ssmQuicksetup_updateConfigurationManagerCmd.Flags().String("manager-arn", "", "The ARN of the configuration manager.")
		ssmQuicksetup_updateConfigurationManagerCmd.Flags().String("name", "", "A name for the configuration manager.")
		ssmQuicksetup_updateConfigurationManagerCmd.MarkFlagRequired("manager-arn")
	})
	ssmQuicksetupCmd.AddCommand(ssmQuicksetup_updateConfigurationManagerCmd)
}
