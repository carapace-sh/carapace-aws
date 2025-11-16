package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmv2_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or overwrites one or more tags for the specified CloudHSM cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmv2_tagResourceCmd).Standalone()

	cloudhsmv2_tagResourceCmd.Flags().String("resource-id", "", "The cluster identifier (ID) for the cluster that you are tagging.")
	cloudhsmv2_tagResourceCmd.Flags().String("tag-list", "", "A list of one or more tags.")
	cloudhsmv2_tagResourceCmd.MarkFlagRequired("resource-id")
	cloudhsmv2_tagResourceCmd.MarkFlagRequired("tag-list")
	cloudhsmv2Cmd.AddCommand(cloudhsmv2_tagResourceCmd)
}
