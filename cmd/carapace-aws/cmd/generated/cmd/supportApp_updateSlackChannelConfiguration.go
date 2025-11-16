package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supportApp_updateSlackChannelConfigurationCmd = &cobra.Command{
	Use:   "update-slack-channel-configuration",
	Short: "Updates the configuration for a Slack channel, such as case update notifications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supportApp_updateSlackChannelConfigurationCmd).Standalone()

	supportApp_updateSlackChannelConfigurationCmd.Flags().String("channel-id", "", "The channel ID in Slack.")
	supportApp_updateSlackChannelConfigurationCmd.Flags().String("channel-name", "", "The Slack channel name that you want to update.")
	supportApp_updateSlackChannelConfigurationCmd.Flags().String("channel-role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that you want to use to perform operations on Amazon Web Services.")
	supportApp_updateSlackChannelConfigurationCmd.Flags().String("notify-on-add-correspondence-to-case", "", "Whether you want to get notified when a support case has a new correspondence.")
	supportApp_updateSlackChannelConfigurationCmd.Flags().String("notify-on-case-severity", "", "The case severity for a support case that you want to receive notifications.")
	supportApp_updateSlackChannelConfigurationCmd.Flags().String("notify-on-create-or-reopen-case", "", "Whether you want to get notified when a support case is created or reopened.")
	supportApp_updateSlackChannelConfigurationCmd.Flags().String("notify-on-resolve-case", "", "Whether you want to get notified when a support case is resolved.")
	supportApp_updateSlackChannelConfigurationCmd.Flags().String("team-id", "", "The team ID in Slack.")
	supportApp_updateSlackChannelConfigurationCmd.MarkFlagRequired("channel-id")
	supportApp_updateSlackChannelConfigurationCmd.MarkFlagRequired("team-id")
	supportAppCmd.AddCommand(supportApp_updateSlackChannelConfigurationCmd)
}
