package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmQuicksetup_createConfigurationManagerCmd = &cobra.Command{
	Use:   "create-configuration-manager",
	Short: "Creates a Quick Setup configuration manager resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmQuicksetup_createConfigurationManagerCmd).Standalone()

	ssmQuicksetup_createConfigurationManagerCmd.Flags().String("configuration-definitions", "", "The definition of the Quick Setup configuration that the configuration manager deploys.")
	ssmQuicksetup_createConfigurationManagerCmd.Flags().String("description", "", "A description of the configuration manager.")
	ssmQuicksetup_createConfigurationManagerCmd.Flags().String("name", "", "A name for the configuration manager.")
	ssmQuicksetup_createConfigurationManagerCmd.Flags().String("tags", "", "Key-value pairs of metadata to assign to the configuration manager.")
	ssmQuicksetup_createConfigurationManagerCmd.MarkFlagRequired("configuration-definitions")
	ssmQuicksetupCmd.AddCommand(ssmQuicksetup_createConfigurationManagerCmd)
}
