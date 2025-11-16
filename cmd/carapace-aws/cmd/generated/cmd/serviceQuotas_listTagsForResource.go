package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of the tags assigned to the specified applied quota.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_listTagsForResourceCmd).Standalone()

	serviceQuotas_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the applied quota for which you want to list tags.")
	serviceQuotas_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	serviceQuotasCmd.AddCommand(serviceQuotas_listTagsForResourceCmd)
}
