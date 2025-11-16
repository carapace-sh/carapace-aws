package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supportApp_registerSlackWorkspaceForOrganizationCmd = &cobra.Command{
	Use:   "register-slack-workspace-for-organization",
	Short: "Registers a Slack workspace for your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supportApp_registerSlackWorkspaceForOrganizationCmd).Standalone()

	supportApp_registerSlackWorkspaceForOrganizationCmd.Flags().String("team-id", "", "The team ID in Slack.")
	supportApp_registerSlackWorkspaceForOrganizationCmd.MarkFlagRequired("team-id")
	supportAppCmd.AddCommand(supportApp_registerSlackWorkspaceForOrganizationCmd)
}
