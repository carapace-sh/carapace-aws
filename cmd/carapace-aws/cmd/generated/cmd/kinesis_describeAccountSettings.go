package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_describeAccountSettingsCmd = &cobra.Command{
	Use:   "describe-account-settings",
	Short: "Describes the account-level settings for Amazon Kinesis Data Streams.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_describeAccountSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesis_describeAccountSettingsCmd).Standalone()

	})
	kinesisCmd.AddCommand(kinesis_describeAccountSettingsCmd)
}
