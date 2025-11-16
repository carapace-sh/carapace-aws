package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Attaches a key-value pair to a resource, as identified by its Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_tagResourceCmd).Standalone()

	transfer_tagResourceCmd.Flags().String("arn", "", "An Amazon Resource Name (ARN) for a specific Amazon Web Services resource, such as a server, user, or role.")
	transfer_tagResourceCmd.Flags().String("tags", "", "Key-value pairs assigned to ARNs that you can use to group and search for resources by type.")
	transfer_tagResourceCmd.MarkFlagRequired("arn")
	transfer_tagResourceCmd.MarkFlagRequired("tags")
	transferCmd.AddCommand(transfer_tagResourceCmd)
}
