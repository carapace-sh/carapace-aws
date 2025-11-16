package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspace_createSpaceCmd = &cobra.Command{
	Use:   "create-space",
	Short: "Creates an AWS re:Post Private private re:Post.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspace_createSpaceCmd).Standalone()

	repostspace_createSpaceCmd.Flags().String("description", "", "A description for the private re:Post.")
	repostspace_createSpaceCmd.Flags().String("name", "", "The name for the private re:Post.")
	repostspace_createSpaceCmd.Flags().String("role-arn", "", "The IAM role that grants permissions to the private re:Post to convert unanswered questions into AWS support tickets.")
	repostspace_createSpaceCmd.Flags().String("subdomain", "", "The subdomain that you use to access your AWS re:Post Private private re:Post.")
	repostspace_createSpaceCmd.Flags().String("supported-email-domains", "", "")
	repostspace_createSpaceCmd.Flags().String("tags", "", "The list of tags associated with the private re:Post.")
	repostspace_createSpaceCmd.Flags().String("tier", "", "The pricing tier for the private re:Post.")
	repostspace_createSpaceCmd.Flags().String("user-kmskey", "", "The AWS KMS key ARN that’s used for the AWS KMS encryption.")
	repostspace_createSpaceCmd.MarkFlagRequired("name")
	repostspace_createSpaceCmd.MarkFlagRequired("subdomain")
	repostspace_createSpaceCmd.MarkFlagRequired("tier")
	repostspaceCmd.AddCommand(repostspace_createSpaceCmd)
}
