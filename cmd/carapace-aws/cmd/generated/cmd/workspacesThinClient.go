package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesThinClientCmd = &cobra.Command{
	Use:   "workspaces-thin-client",
	Short: "Amazon WorkSpaces Thin Client is an affordable device built to work with Amazon Web Services End User Computing (EUC) virtual desktops to provide users with a complete cloud desktop solution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesThinClientCmd).Standalone()

	rootCmd.AddCommand(workspacesThinClientCmd)
}
