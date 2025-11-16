package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_searchUsersByImageCmd = &cobra.Command{
	Use:   "search-users-by-image",
	Short: "Searches for UserIDs using a supplied image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_searchUsersByImageCmd).Standalone()

	rekognition_searchUsersByImageCmd.Flags().String("collection-id", "", "The ID of an existing collection containing the UserID.")
	rekognition_searchUsersByImageCmd.Flags().String("image", "", "")
	rekognition_searchUsersByImageCmd.Flags().String("max-users", "", "Maximum number of UserIDs to return.")
	rekognition_searchUsersByImageCmd.Flags().String("quality-filter", "", "A filter that specifies a quality bar for how much filtering is done to identify faces.")
	rekognition_searchUsersByImageCmd.Flags().String("user-match-threshold", "", "Specifies the minimum confidence in the UserID match to return.")
	rekognition_searchUsersByImageCmd.MarkFlagRequired("collection-id")
	rekognition_searchUsersByImageCmd.MarkFlagRequired("image")
	rekognitionCmd.AddCommand(rekognition_searchUsersByImageCmd)
}
