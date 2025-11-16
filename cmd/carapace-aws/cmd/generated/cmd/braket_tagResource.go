package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var braket_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Add a tag to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(braket_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(braket_tagResourceCmd).Standalone()

		braket_tagResourceCmd.Flags().String("resource-arn", "", "Specify the `resourceArn` of the resource to which a tag will be added.")
		braket_tagResourceCmd.Flags().String("tags", "", "Specify the tags to add to the resource.")
		braket_tagResourceCmd.MarkFlagRequired("resource-arn")
		braket_tagResourceCmd.MarkFlagRequired("tags")
	})
	braketCmd.AddCommand(braket_tagResourceCmd)
}
