package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWebCmd = &cobra.Command{
	Use:   "workspaces-web",
	Short: "Amazon WorkSpaces Secure Browser is a low cost, fully managed WorkSpace built specifically to facilitate secure, web-based workloads.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWebCmd).Standalone()

	rootCmd.AddCommand(workspacesWebCmd)
}
