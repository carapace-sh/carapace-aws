package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_deleteObjectCmd = &cobra.Command{
	Use:   "delete-object",
	Short: "Deletes an object and its associated attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_deleteObjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_deleteObjectCmd).Standalone()

		clouddirectory_deleteObjectCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Directory]() where the object resides.")
		clouddirectory_deleteObjectCmd.Flags().String("object-reference", "", "A reference that identifies the object.")
		clouddirectory_deleteObjectCmd.MarkFlagRequired("directory-arn")
		clouddirectory_deleteObjectCmd.MarkFlagRequired("object-reference")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_deleteObjectCmd)
}
