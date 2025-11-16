package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags on an Amazon Neptune resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_listTagsForResourceCmd).Standalone()

	neptune_listTagsForResourceCmd.Flags().String("filters", "", "This parameter is not currently supported.")
	neptune_listTagsForResourceCmd.Flags().String("resource-name", "", "The Amazon Neptune resource with tags to be listed.")
	neptune_listTagsForResourceCmd.MarkFlagRequired("resource-name")
	neptuneCmd.AddCommand(neptune_listTagsForResourceCmd)
}
