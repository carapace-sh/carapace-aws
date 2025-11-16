package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_updateHostedZoneCommentCmd = &cobra.Command{
	Use:   "update-hosted-zone-comment",
	Short: "Updates the comment for a specified hosted zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_updateHostedZoneCommentCmd).Standalone()

	route53_updateHostedZoneCommentCmd.Flags().String("comment", "", "The new comment for the hosted zone.")
	route53_updateHostedZoneCommentCmd.Flags().String("id", "", "The ID for the hosted zone that you want to update the comment for.")
	route53_updateHostedZoneCommentCmd.MarkFlagRequired("id")
	route53Cmd.AddCommand(route53_updateHostedZoneCommentCmd)
}
