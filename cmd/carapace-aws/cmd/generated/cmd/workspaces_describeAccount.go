package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeAccountCmd = &cobra.Command{
	Use:   "describe-account",
	Short: "Retrieves a list that describes the configuration of Bring Your Own License (BYOL) for the specified account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeAccountCmd).Standalone()

	workspacesCmd.AddCommand(workspaces_describeAccountCmd)
}
