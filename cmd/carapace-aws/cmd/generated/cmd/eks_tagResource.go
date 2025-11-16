package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates the specified tags to an Amazon EKS resource with the specified `resourceArn`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_tagResourceCmd).Standalone()

		eks_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to add tags to.")
		eks_tagResourceCmd.Flags().String("tags", "", "Metadata that assists with categorization and organization.")
		eks_tagResourceCmd.MarkFlagRequired("resource-arn")
		eks_tagResourceCmd.MarkFlagRequired("tags")
	})
	eksCmd.AddCommand(eks_tagResourceCmd)
}
