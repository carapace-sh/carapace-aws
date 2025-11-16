package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from the specified applied quota.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_untagResourceCmd).Standalone()

	serviceQuotas_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the applied quota that you want to untag.")
	serviceQuotas_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags that you want to remove from the resource.")
	serviceQuotas_untagResourceCmd.MarkFlagRequired("resource-arn")
	serviceQuotas_untagResourceCmd.MarkFlagRequired("tag-keys")
	serviceQuotasCmd.AddCommand(serviceQuotas_untagResourceCmd)
}
