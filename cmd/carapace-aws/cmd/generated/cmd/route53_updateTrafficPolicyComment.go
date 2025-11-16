package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_updateTrafficPolicyCommentCmd = &cobra.Command{
	Use:   "update-traffic-policy-comment",
	Short: "Updates the comment for a specified traffic policy version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_updateTrafficPolicyCommentCmd).Standalone()

	route53_updateTrafficPolicyCommentCmd.Flags().String("comment", "", "The new comment for the specified traffic policy and version.")
	route53_updateTrafficPolicyCommentCmd.Flags().String("id", "", "The value of `Id` for the traffic policy that you want to update the comment for.")
	route53_updateTrafficPolicyCommentCmd.Flags().String("version", "", "The value of `Version` for the traffic policy that you want to update the comment for.")
	route53_updateTrafficPolicyCommentCmd.MarkFlagRequired("comment")
	route53_updateTrafficPolicyCommentCmd.MarkFlagRequired("id")
	route53_updateTrafficPolicyCommentCmd.MarkFlagRequired("version")
	route53Cmd.AddCommand(route53_updateTrafficPolicyCommentCmd)
}
