package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstreamCmd = &cobra.Command{
	Use:   "appstream",
	Short: "Amazon AppStream 2.0\n\nThis is the *Amazon AppStream 2.0 API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstreamCmd).Standalone()

	})
	rootCmd.AddCommand(appstreamCmd)
}
