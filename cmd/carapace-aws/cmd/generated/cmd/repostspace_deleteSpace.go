package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspace_deleteSpaceCmd = &cobra.Command{
	Use:   "delete-space",
	Short: "Deletes an AWS re:Post Private private re:Post.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspace_deleteSpaceCmd).Standalone()

	repostspace_deleteSpaceCmd.Flags().String("space-id", "", "The unique ID of the private re:Post.")
	repostspace_deleteSpaceCmd.MarkFlagRequired("space-id")
	repostspaceCmd.AddCommand(repostspace_deleteSpaceCmd)
}
