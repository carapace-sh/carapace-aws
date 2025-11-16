package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_listTagsCmd = &cobra.Command{
	Use:   "list-tags",
	Short: "List all of the tags for a DAX cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_listTagsCmd).Standalone()

	dax_listTagsCmd.Flags().String("next-token", "", "An optional token returned from a prior request.")
	dax_listTagsCmd.Flags().String("resource-name", "", "The name of the DAX resource to which the tags belong.")
	dax_listTagsCmd.MarkFlagRequired("resource-name")
	daxCmd.AddCommand(dax_listTagsCmd)
}
