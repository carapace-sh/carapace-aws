package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_putAccountSettingCmd = &cobra.Command{
	Use:   "put-account-setting",
	Short: "Allows you to change the basic scan type version or registry policy scope.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_putAccountSettingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_putAccountSettingCmd).Standalone()

		ecr_putAccountSettingCmd.Flags().String("name", "", "The name of the account setting, such as `BASIC_SCAN_TYPE_VERSION` or `REGISTRY_POLICY_SCOPE`.")
		ecr_putAccountSettingCmd.Flags().String("value", "", "Setting value that is specified.")
		ecr_putAccountSettingCmd.MarkFlagRequired("name")
		ecr_putAccountSettingCmd.MarkFlagRequired("value")
	})
	ecrCmd.AddCommand(ecr_putAccountSettingCmd)
}
