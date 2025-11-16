package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Untags a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_untagResourceCmd).Standalone()

	connectcases_untagResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN)")
	connectcases_untagResourceCmd.Flags().String("tag-keys", "", "List of tag keys.")
	connectcases_untagResourceCmd.MarkFlagRequired("arn")
	connectcases_untagResourceCmd.MarkFlagRequired("tag-keys")
	connectcasesCmd.AddCommand(connectcases_untagResourceCmd)
}
