package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supportApp_deleteSlackChannelConfigurationCmd = &cobra.Command{
	Use:   "delete-slack-channel-configuration",
	Short: "Deletes a Slack channel configuration from your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supportApp_deleteSlackChannelConfigurationCmd).Standalone()

	supportApp_deleteSlackChannelConfigurationCmd.Flags().String("channel-id", "", "The channel ID in Slack.")
	supportApp_deleteSlackChannelConfigurationCmd.Flags().String("team-id", "", "The team ID in Slack.")
	supportApp_deleteSlackChannelConfigurationCmd.MarkFlagRequired("channel-id")
	supportApp_deleteSlackChannelConfigurationCmd.MarkFlagRequired("team-id")
	supportAppCmd.AddCommand(supportApp_deleteSlackChannelConfigurationCmd)
}
