package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supportAppCmd = &cobra.Command{
	Use:   "support-app",
	Short: "Amazon Web Services Support App in Slack\n\nYou can use the Amazon Web Services Support App in Slack API to manage your support cases in Slack for your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supportAppCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(supportAppCmd).Standalone()

	})
	rootCmd.AddCommand(supportAppCmd)
}
