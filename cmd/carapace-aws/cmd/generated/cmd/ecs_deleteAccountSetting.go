package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_deleteAccountSettingCmd = &cobra.Command{
	Use:   "delete-account-setting",
	Short: "Disables an account setting for a specified user, role, or the root user for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_deleteAccountSettingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_deleteAccountSettingCmd).Standalone()

		ecs_deleteAccountSettingCmd.Flags().String("name", "", "The resource name to disable the account setting for.")
		ecs_deleteAccountSettingCmd.Flags().String("principal-arn", "", "The Amazon Resource Name (ARN) of the principal.")
		ecs_deleteAccountSettingCmd.MarkFlagRequired("name")
	})
	ecsCmd.AddCommand(ecs_deleteAccountSettingCmd)
}
