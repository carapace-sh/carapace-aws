package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates the specified tags to a resource with the specified `resourceArn`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_tagResourceCmd).Standalone()

	billingconductor_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to which to add tags.")
	billingconductor_tagResourceCmd.Flags().String("tags", "", "The tags to add to the resource as a list of key-value pairs.")
	billingconductor_tagResourceCmd.MarkFlagRequired("resource-arn")
	billingconductor_tagResourceCmd.MarkFlagRequired("tags")
	billingconductorCmd.AddCommand(billingconductor_tagResourceCmd)
}
