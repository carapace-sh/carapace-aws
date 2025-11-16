package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_searchUsersCmd = &cobra.Command{
	Use:   "search-users",
	Short: "Searches for UserIDs within a collection based on a `FaceId` or `UserId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_searchUsersCmd).Standalone()

	rekognition_searchUsersCmd.Flags().String("collection-id", "", "The ID of an existing collection containing the UserID, used with a UserId or FaceId.")
	rekognition_searchUsersCmd.Flags().String("face-id", "", "ID for the existing face.")
	rekognition_searchUsersCmd.Flags().String("max-users", "", "Maximum number of identities to return.")
	rekognition_searchUsersCmd.Flags().String("user-id", "", "ID for the existing User.")
	rekognition_searchUsersCmd.Flags().String("user-match-threshold", "", "Optional value that specifies the minimum confidence in the matched UserID to return.")
	rekognition_searchUsersCmd.MarkFlagRequired("collection-id")
	rekognitionCmd.AddCommand(rekognition_searchUsersCmd)
}
