package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supportApp_createSlackChannelConfigurationCmd = &cobra.Command{
	Use:   "create-slack-channel-configuration",
	Short: "Creates a Slack channel configuration for your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supportApp_createSlackChannelConfigurationCmd).Standalone()

	supportApp_createSlackChannelConfigurationCmd.Flags().String("channel-id", "", "The channel ID in Slack.")
	supportApp_createSlackChannelConfigurationCmd.Flags().String("channel-name", "", "The name of the Slack channel that you configure for the Amazon Web Services Support App.")
	supportApp_createSlackChannelConfigurationCmd.Flags().String("channel-role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that you want to use to perform operations on Amazon Web Services.")
	supportApp_createSlackChannelConfigurationCmd.Flags().String("notify-on-add-correspondence-to-case", "", "Whether you want to get notified when a support case has a new correspondence.")
	supportApp_createSlackChannelConfigurationCmd.Flags().String("notify-on-case-severity", "", "The case severity for a support case that you want to receive notifications.")
	supportApp_createSlackChannelConfigurationCmd.Flags().String("notify-on-create-or-reopen-case", "", "Whether you want to get notified when a support case is created or reopened.")
	supportApp_createSlackChannelConfigurationCmd.Flags().String("notify-on-resolve-case", "", "Whether you want to get notified when a support case is resolved.")
	supportApp_createSlackChannelConfigurationCmd.Flags().String("team-id", "", "The team ID in Slack.")
	supportApp_createSlackChannelConfigurationCmd.MarkFlagRequired("channel-id")
	supportApp_createSlackChannelConfigurationCmd.MarkFlagRequired("channel-role-arn")
	supportApp_createSlackChannelConfigurationCmd.MarkFlagRequired("notify-on-case-severity")
	supportApp_createSlackChannelConfigurationCmd.MarkFlagRequired("team-id")
	supportAppCmd.AddCommand(supportApp_createSlackChannelConfigurationCmd)
}
