package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_describeTagsCmd = &cobra.Command{
	Use:   "describe-tags",
	Short: "Describes the tags associated with the specified load balancers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_describeTagsCmd).Standalone()

	elb_describeTagsCmd.Flags().String("load-balancer-names", "", "The names of the load balancers.")
	elb_describeTagsCmd.MarkFlagRequired("load-balancer-names")
	elbCmd.AddCommand(elb_describeTagsCmd)
}
