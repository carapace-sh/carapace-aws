package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Creates tag for a resource by specifying the ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_tagResourceCmd).Standalone()

	ssmSap_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	ssmSap_tagResourceCmd.Flags().String("tags", "", "The tags on a resource.")
	ssmSap_tagResourceCmd.MarkFlagRequired("resource-arn")
	ssmSap_tagResourceCmd.MarkFlagRequired("tags")
	ssmSapCmd.AddCommand(ssmSap_tagResourceCmd)
}
