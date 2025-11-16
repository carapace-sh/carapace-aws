package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_listAccountSettingsCmd = &cobra.Command{
	Use:   "list-account-settings",
	Short: "Lists the account settings for a specified principal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_listAccountSettingsCmd).Standalone()

	ecs_listAccountSettingsCmd.Flags().Bool("effective-settings", false, "Determines whether to return the effective settings.")
	ecs_listAccountSettingsCmd.Flags().String("max-results", "", "The maximum number of account setting results returned by `ListAccountSettings` in paginated output.")
	ecs_listAccountSettingsCmd.Flags().String("name", "", "The name of the account setting you want to list the settings for.")
	ecs_listAccountSettingsCmd.Flags().String("next-token", "", "The `nextToken` value returned from a `ListAccountSettings` request indicating that more results are available to fulfill the request and further calls will be needed.")
	ecs_listAccountSettingsCmd.Flags().Bool("no-effective-settings", false, "Determines whether to return the effective settings.")
	ecs_listAccountSettingsCmd.Flags().String("principal-arn", "", "The ARN of the principal, which can be a user, role, or the root user.")
	ecs_listAccountSettingsCmd.Flags().String("value", "", "The value of the account settings to filter results with.")
	ecs_listAccountSettingsCmd.Flag("no-effective-settings").Hidden = true
	ecsCmd.AddCommand(ecs_listAccountSettingsCmd)
}
