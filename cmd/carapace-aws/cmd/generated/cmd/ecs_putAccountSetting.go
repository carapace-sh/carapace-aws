package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_putAccountSettingCmd = &cobra.Command{
	Use:   "put-account-setting",
	Short: "Modifies an account setting.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_putAccountSettingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_putAccountSettingCmd).Standalone()

		ecs_putAccountSettingCmd.Flags().String("name", "", "The Amazon ECS account setting name to modify.")
		ecs_putAccountSettingCmd.Flags().String("principal-arn", "", "The ARN of the principal, which can be a user, role, or the root user.")
		ecs_putAccountSettingCmd.Flags().String("value", "", "The account setting value for the specified principal ARN.")
		ecs_putAccountSettingCmd.MarkFlagRequired("name")
		ecs_putAccountSettingCmd.MarkFlagRequired("value")
	})
	ecsCmd.AddCommand(ecs_putAccountSettingCmd)
}
