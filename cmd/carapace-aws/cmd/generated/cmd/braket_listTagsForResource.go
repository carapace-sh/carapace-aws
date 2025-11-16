package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var braket_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Shows the tags associated with this resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(braket_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(braket_listTagsForResourceCmd).Standalone()

		braket_listTagsForResourceCmd.Flags().String("resource-arn", "", "Specify the `resourceArn` for the resource whose tags to display.")
		braket_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	braketCmd.AddCommand(braket_listTagsForResourceCmd)
}
