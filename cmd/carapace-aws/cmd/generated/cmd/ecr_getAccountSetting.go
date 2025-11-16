package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_getAccountSettingCmd = &cobra.Command{
	Use:   "get-account-setting",
	Short: "Retrieves the account setting value for the specified setting name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_getAccountSettingCmd).Standalone()

	ecr_getAccountSettingCmd.Flags().String("name", "", "The name of the account setting, such as `BASIC_SCAN_TYPE_VERSION` or `REGISTRY_POLICY_SCOPE`.")
	ecr_getAccountSettingCmd.MarkFlagRequired("name")
	ecrCmd.AddCommand(ecr_getAccountSettingCmd)
}
