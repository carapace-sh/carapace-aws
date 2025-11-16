package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_lookupPolicyCmd = &cobra.Command{
	Use:   "lookup-policy",
	Short: "Lists all policies from the root of the [Directory]() to the object specified.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_lookupPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_lookupPolicyCmd).Standalone()

		clouddirectory_lookupPolicyCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Directory]().")
		clouddirectory_lookupPolicyCmd.Flags().String("max-results", "", "The maximum number of items to be retrieved in a single call.")
		clouddirectory_lookupPolicyCmd.Flags().String("next-token", "", "The token to request the next page of results.")
		clouddirectory_lookupPolicyCmd.Flags().String("object-reference", "", "Reference that identifies the object whose policies will be looked up.")
		clouddirectory_lookupPolicyCmd.MarkFlagRequired("directory-arn")
		clouddirectory_lookupPolicyCmd.MarkFlagRequired("object-reference")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_lookupPolicyCmd)
}
