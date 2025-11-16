package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_deactivateUserCmd = &cobra.Command{
	Use:   "deactivate-user",
	Short: "Deactivates the specified user, which revokes the user's access to Amazon WorkDocs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_deactivateUserCmd).Standalone()

	workdocs_deactivateUserCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_deactivateUserCmd.Flags().String("user-id", "", "The ID of the user.")
	workdocs_deactivateUserCmd.MarkFlagRequired("user-id")
	workdocsCmd.AddCommand(workdocs_deactivateUserCmd)
}
