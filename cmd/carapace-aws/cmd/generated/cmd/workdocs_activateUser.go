package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_activateUserCmd = &cobra.Command{
	Use:   "activate-user",
	Short: "Activates the specified user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_activateUserCmd).Standalone()

	workdocs_activateUserCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_activateUserCmd.Flags().String("user-id", "", "The ID of the user.")
	workdocs_activateUserCmd.MarkFlagRequired("user-id")
	workdocsCmd.AddCommand(workdocs_activateUserCmd)
}
