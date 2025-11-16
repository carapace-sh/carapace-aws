package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Creates a new User within a collection specified by `CollectionId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_createUserCmd).Standalone()

	rekognition_createUserCmd.Flags().String("client-request-token", "", "Idempotent token used to identify the request to `CreateUser`.")
	rekognition_createUserCmd.Flags().String("collection-id", "", "The ID of an existing collection to which the new UserID needs to be created.")
	rekognition_createUserCmd.Flags().String("user-id", "", "ID for the UserID to be created.")
	rekognition_createUserCmd.MarkFlagRequired("collection-id")
	rekognition_createUserCmd.MarkFlagRequired("user-id")
	rekognitionCmd.AddCommand(rekognition_createUserCmd)
}
