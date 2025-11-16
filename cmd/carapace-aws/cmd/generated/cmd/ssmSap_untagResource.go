package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Delete the tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmSap_untagResourceCmd).Standalone()

		ssmSap_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		ssmSap_untagResourceCmd.Flags().String("tag-keys", "", "Adds/updates or removes credentials for applications registered with AWS Systems Manager for SAP.")
		ssmSap_untagResourceCmd.MarkFlagRequired("resource-arn")
		ssmSap_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	ssmSapCmd.AddCommand(ssmSap_untagResourceCmd)
}
