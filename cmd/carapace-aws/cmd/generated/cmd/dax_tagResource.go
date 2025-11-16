package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates a set of tags with a DAX resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_tagResourceCmd).Standalone()

	dax_tagResourceCmd.Flags().String("resource-name", "", "The name of the DAX resource to which tags should be added.")
	dax_tagResourceCmd.Flags().String("tags", "", "The tags to be assigned to the DAX resource.")
	dax_tagResourceCmd.MarkFlagRequired("resource-name")
	dax_tagResourceCmd.MarkFlagRequired("tags")
	daxCmd.AddCommand(dax_tagResourceCmd)
}
