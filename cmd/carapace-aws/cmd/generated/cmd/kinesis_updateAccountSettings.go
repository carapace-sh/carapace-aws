package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_updateAccountSettingsCmd = &cobra.Command{
	Use:   "update-account-settings",
	Short: "Updates the account-level settings for Amazon Kinesis Data Streams.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_updateAccountSettingsCmd).Standalone()

	kinesis_updateAccountSettingsCmd.Flags().String("minimum-throughput-billing-commitment", "", "Specifies the minimum throughput billing commitment configuration for your account.")
	kinesis_updateAccountSettingsCmd.MarkFlagRequired("minimum-throughput-billing-commitment")
	kinesisCmd.AddCommand(kinesis_updateAccountSettingsCmd)
}
