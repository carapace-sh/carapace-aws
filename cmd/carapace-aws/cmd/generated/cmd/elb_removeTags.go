package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_removeTagsCmd = &cobra.Command{
	Use:   "remove-tags",
	Short: "Removes one or more tags from the specified load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_removeTagsCmd).Standalone()

	elb_removeTagsCmd.Flags().String("load-balancer-names", "", "The name of the load balancer.")
	elb_removeTagsCmd.Flags().String("tags", "", "The list of tag keys to remove.")
	elb_removeTagsCmd.MarkFlagRequired("load-balancer-names")
	elb_removeTagsCmd.MarkFlagRequired("tags")
	elbCmd.AddCommand(elb_removeTagsCmd)
}
