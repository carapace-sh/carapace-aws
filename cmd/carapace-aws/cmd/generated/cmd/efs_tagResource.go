package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Creates a tag for an EFS resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_tagResourceCmd).Standalone()

	efs_tagResourceCmd.Flags().String("resource-id", "", "The ID specifying the EFS resource that you want to create a tag for.")
	efs_tagResourceCmd.Flags().String("tags", "", "An array of `Tag` objects to add.")
	efs_tagResourceCmd.MarkFlagRequired("resource-id")
	efs_tagResourceCmd.MarkFlagRequired("tags")
	efsCmd.AddCommand(efs_tagResourceCmd)
}
