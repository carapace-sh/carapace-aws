package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds tags to the specified applied quota.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(serviceQuotas_tagResourceCmd).Standalone()

		serviceQuotas_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the applied quota.")
		serviceQuotas_tagResourceCmd.Flags().String("tags", "", "The tags that you want to add to the resource.")
		serviceQuotas_tagResourceCmd.MarkFlagRequired("resource-arn")
		serviceQuotas_tagResourceCmd.MarkFlagRequired("tags")
	})
	serviceQuotasCmd.AddCommand(serviceQuotas_tagResourceCmd)
}
