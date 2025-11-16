package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspace_getSpaceCmd = &cobra.Command{
	Use:   "get-space",
	Short: "Displays information about the AWS re:Post Private private re:Post.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspace_getSpaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(repostspace_getSpaceCmd).Standalone()

		repostspace_getSpaceCmd.Flags().String("space-id", "", "The ID of the private re:Post.")
		repostspace_getSpaceCmd.MarkFlagRequired("space-id")
	})
	repostspaceCmd.AddCommand(repostspace_getSpaceCmd)
}
