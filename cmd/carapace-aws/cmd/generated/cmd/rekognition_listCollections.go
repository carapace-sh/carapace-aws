package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_listCollectionsCmd = &cobra.Command{
	Use:   "list-collections",
	Short: "Returns list of collection IDs in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_listCollectionsCmd).Standalone()

	rekognition_listCollectionsCmd.Flags().String("max-results", "", "Maximum number of collection IDs to return.")
	rekognition_listCollectionsCmd.Flags().String("next-token", "", "Pagination token from the previous response.")
	rekognitionCmd.AddCommand(rekognition_listCollectionsCmd)
}
