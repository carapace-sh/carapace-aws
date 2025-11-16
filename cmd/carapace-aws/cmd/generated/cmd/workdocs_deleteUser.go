package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "Deletes the specified user from a Simple AD or Microsoft AD directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_deleteUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_deleteUserCmd).Standalone()

		workdocs_deleteUserCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
		workdocs_deleteUserCmd.Flags().String("user-id", "", "The ID of the user.")
		workdocs_deleteUserCmd.MarkFlagRequired("user-id")
	})
	workdocsCmd.AddCommand(workdocs_deleteUserCmd)
}
