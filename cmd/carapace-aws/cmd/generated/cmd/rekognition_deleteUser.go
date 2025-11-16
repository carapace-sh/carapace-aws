package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "Deletes the specified UserID within the collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_deleteUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_deleteUserCmd).Standalone()

		rekognition_deleteUserCmd.Flags().String("client-request-token", "", "Idempotent token used to identify the request to `DeleteUser`.")
		rekognition_deleteUserCmd.Flags().String("collection-id", "", "The ID of an existing collection from which the UserID needs to be deleted.")
		rekognition_deleteUserCmd.Flags().String("user-id", "", "ID for the UserID to be deleted.")
		rekognition_deleteUserCmd.MarkFlagRequired("collection-id")
		rekognition_deleteUserCmd.MarkFlagRequired("user-id")
	})
	rekognitionCmd.AddCommand(rekognition_deleteUserCmd)
}
