package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_listTagsCmd = &cobra.Command{
	Use:   "list-tags",
	Short: "Lists all tags currently on a named resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_listTagsCmd).Standalone()

	memorydb_listTagsCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource for which you want the list of tags.")
	memorydb_listTagsCmd.MarkFlagRequired("resource-arn")
	memorydbCmd.AddCommand(memorydb_listTagsCmd)
}
