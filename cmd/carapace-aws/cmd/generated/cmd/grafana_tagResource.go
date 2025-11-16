package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "The `TagResource` operation associates tags with an Amazon Managed Grafana resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(grafana_tagResourceCmd).Standalone()

		grafana_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource the tag is associated with.")
		grafana_tagResourceCmd.Flags().String("tags", "", "The list of tag keys and values to associate with the resource.")
		grafana_tagResourceCmd.MarkFlagRequired("resource-arn")
		grafana_tagResourceCmd.MarkFlagRequired("tags")
	})
	grafanaCmd.AddCommand(grafana_tagResourceCmd)
}
