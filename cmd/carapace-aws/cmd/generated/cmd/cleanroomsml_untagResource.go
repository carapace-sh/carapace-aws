package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes metadata tags from a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_untagResourceCmd).Standalone()

	cleanroomsml_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to remove tags from.")
	cleanroomsml_untagResourceCmd.Flags().String("tag-keys", "", "The key values of tags that you want to remove.")
	cleanroomsml_untagResourceCmd.MarkFlagRequired("resource-arn")
	cleanroomsml_untagResourceCmd.MarkFlagRequired("tag-keys")
	cleanroomsmlCmd.AddCommand(cleanroomsml_untagResourceCmd)
}
