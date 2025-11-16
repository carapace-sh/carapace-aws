package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_listUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: "Returns metadata of the User such as `UserID` in the specified collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_listUsersCmd).Standalone()

	rekognition_listUsersCmd.Flags().String("collection-id", "", "The ID of an existing collection.")
	rekognition_listUsersCmd.Flags().String("max-results", "", "Maximum number of UsersID to return.")
	rekognition_listUsersCmd.Flags().String("next-token", "", "Pagingation token to receive the next set of UsersID.")
	rekognition_listUsersCmd.MarkFlagRequired("collection-id")
	rekognitionCmd.AddCommand(rekognition_listUsersCmd)
}
