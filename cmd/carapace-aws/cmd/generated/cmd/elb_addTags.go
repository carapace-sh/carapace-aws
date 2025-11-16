package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_addTagsCmd = &cobra.Command{
	Use:   "add-tags",
	Short: "Adds the specified tags to the specified load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_addTagsCmd).Standalone()

	elb_addTagsCmd.Flags().String("load-balancer-names", "", "The name of the load balancer.")
	elb_addTagsCmd.Flags().String("tags", "", "The tags.")
	elb_addTagsCmd.MarkFlagRequired("load-balancer-names")
	elb_addTagsCmd.MarkFlagRequired("tags")
	elbCmd.AddCommand(elb_addTagsCmd)
}
