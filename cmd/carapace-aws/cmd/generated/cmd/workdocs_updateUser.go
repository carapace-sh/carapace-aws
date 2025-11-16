package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_updateUserCmd = &cobra.Command{
	Use:   "update-user",
	Short: "Updates the specified attributes of the specified user, and grants or revokes administrative privileges to the Amazon WorkDocs site.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_updateUserCmd).Standalone()

	workdocs_updateUserCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_updateUserCmd.Flags().String("given-name", "", "The given name of the user.")
	workdocs_updateUserCmd.Flags().String("grant-poweruser-privileges", "", "Boolean value to determine whether the user is granted Power user privileges.")
	workdocs_updateUserCmd.Flags().String("locale", "", "The locale of the user.")
	workdocs_updateUserCmd.Flags().String("storage-rule", "", "The amount of storage for the user.")
	workdocs_updateUserCmd.Flags().String("surname", "", "The surname of the user.")
	workdocs_updateUserCmd.Flags().String("time-zone-id", "", "The time zone ID of the user.")
	workdocs_updateUserCmd.Flags().String("type", "", "The type of the user.")
	workdocs_updateUserCmd.Flags().String("user-id", "", "The ID of the user.")
	workdocs_updateUserCmd.MarkFlagRequired("user-id")
	workdocsCmd.AddCommand(workdocs_updateUserCmd)
}
