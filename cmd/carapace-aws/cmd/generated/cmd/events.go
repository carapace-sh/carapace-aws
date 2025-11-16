package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Amazon EventBridge helps you to respond to state changes in your Amazon Web Services resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eventsCmd).Standalone()

	rootCmd.AddCommand(eventsCmd)
}
