package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesCmd = &cobra.Command{
	Use:   "workspaces",
	Short: "Amazon WorkSpaces Service\n\nAmazon WorkSpaces enables you to provision virtual, cloud-based Microsoft Windows or Amazon Linux desktops for your users, known as *WorkSpaces*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesCmd).Standalone()

	})
	rootCmd.AddCommand(workspacesCmd)
}
