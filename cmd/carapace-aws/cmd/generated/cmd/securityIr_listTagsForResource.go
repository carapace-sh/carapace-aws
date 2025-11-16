package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns currently configured tags on a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_listTagsForResourceCmd).Standalone()

	securityIr_listTagsForResourceCmd.Flags().String("resource-arn", "", "Required element for ListTagsForResource to provide the ARN to identify a specific resource.")
	securityIr_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	securityIrCmd.AddCommand(securityIr_listTagsForResourceCmd)
}
