package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspace_updateSpaceCmd = &cobra.Command{
	Use:   "update-space",
	Short: "Modifies an existing AWS re:Post Private private re:Post.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspace_updateSpaceCmd).Standalone()

	repostspace_updateSpaceCmd.Flags().String("description", "", "A description for the private re:Post.")
	repostspace_updateSpaceCmd.Flags().String("role-arn", "", "The IAM role that grants permissions to the private re:Post to convert unanswered questions into AWS support tickets.")
	repostspace_updateSpaceCmd.Flags().String("space-id", "", "The unique ID of this private re:Post.")
	repostspace_updateSpaceCmd.Flags().String("supported-email-domains", "", "")
	repostspace_updateSpaceCmd.Flags().String("tier", "", "The pricing tier of this private re:Post.")
	repostspace_updateSpaceCmd.MarkFlagRequired("space-id")
	repostspaceCmd.AddCommand(repostspace_updateSpaceCmd)
}
