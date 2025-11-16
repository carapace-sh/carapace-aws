package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds the specified tags to the specified Direct Connect resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_tagResourceCmd).Standalone()

	directconnect_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	directconnect_tagResourceCmd.Flags().String("tags", "", "The tags to add.")
	directconnect_tagResourceCmd.MarkFlagRequired("resource-arn")
	directconnect_tagResourceCmd.MarkFlagRequired("tags")
	directconnectCmd.AddCommand(directconnect_tagResourceCmd)
}
