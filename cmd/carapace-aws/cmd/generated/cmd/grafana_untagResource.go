package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "The `UntagResource` operation removes the association of the tag with the Amazon Managed Grafana resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(grafana_untagResourceCmd).Standalone()

		grafana_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource the tag association is removed from.")
		grafana_untagResourceCmd.Flags().String("tag-keys", "", "The key values of the tag to be removed from the resource.")
		grafana_untagResourceCmd.MarkFlagRequired("resource-arn")
		grafana_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	grafanaCmd.AddCommand(grafana_untagResourceCmd)
}
