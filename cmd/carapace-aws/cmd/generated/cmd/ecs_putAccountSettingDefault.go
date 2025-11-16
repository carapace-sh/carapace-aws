package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_putAccountSettingDefaultCmd = &cobra.Command{
	Use:   "put-account-setting-default",
	Short: "Modifies an account setting for all users on an account for whom no individual account setting has been specified.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_putAccountSettingDefaultCmd).Standalone()

	ecs_putAccountSettingDefaultCmd.Flags().String("name", "", "The resource name for which to modify the account setting.")
	ecs_putAccountSettingDefaultCmd.Flags().String("value", "", "The account setting value for the specified principal ARN.")
	ecs_putAccountSettingDefaultCmd.MarkFlagRequired("name")
	ecs_putAccountSettingDefaultCmd.MarkFlagRequired("value")
	ecsCmd.AddCommand(ecs_putAccountSettingDefaultCmd)
}
