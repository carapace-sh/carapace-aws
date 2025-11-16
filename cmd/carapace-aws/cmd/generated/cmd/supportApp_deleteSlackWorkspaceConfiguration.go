package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supportApp_deleteSlackWorkspaceConfigurationCmd = &cobra.Command{
	Use:   "delete-slack-workspace-configuration",
	Short: "Deletes a Slack workspace configuration from your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supportApp_deleteSlackWorkspaceConfigurationCmd).Standalone()

	supportApp_deleteSlackWorkspaceConfigurationCmd.Flags().String("team-id", "", "The team ID in Slack.")
	supportApp_deleteSlackWorkspaceConfigurationCmd.MarkFlagRequired("team-id")
	supportAppCmd.AddCommand(supportApp_deleteSlackWorkspaceConfigurationCmd)
}
