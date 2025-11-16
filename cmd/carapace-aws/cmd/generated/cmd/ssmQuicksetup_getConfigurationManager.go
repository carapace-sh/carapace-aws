package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmQuicksetup_getConfigurationManagerCmd = &cobra.Command{
	Use:   "get-configuration-manager",
	Short: "Returns a configuration manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmQuicksetup_getConfigurationManagerCmd).Standalone()

	ssmQuicksetup_getConfigurationManagerCmd.Flags().String("manager-arn", "", "The ARN of the configuration manager.")
	ssmQuicksetup_getConfigurationManagerCmd.MarkFlagRequired("manager-arn")
	ssmQuicksetupCmd.AddCommand(ssmQuicksetup_getConfigurationManagerCmd)
}
