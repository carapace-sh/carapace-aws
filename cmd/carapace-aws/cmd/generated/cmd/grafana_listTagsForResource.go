package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "The `ListTagsForResource` operation returns the tags that are associated with the Amazon Managed Service for Grafana resource specified by the `resourceArn`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(grafana_listTagsForResourceCmd).Standalone()

		grafana_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource the list of tags are associated with.")
		grafana_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	grafanaCmd.AddCommand(grafana_listTagsForResourceCmd)
}
