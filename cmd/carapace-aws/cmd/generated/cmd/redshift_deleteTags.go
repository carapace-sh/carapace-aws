package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deleteTagsCmd = &cobra.Command{
	Use:   "delete-tags",
	Short: "Deletes tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deleteTagsCmd).Standalone()

	redshift_deleteTagsCmd.Flags().String("resource-name", "", "The Amazon Resource Name (ARN) from which you want to remove the tag or tags.")
	redshift_deleteTagsCmd.Flags().String("tag-keys", "", "The tag key that you want to delete.")
	redshift_deleteTagsCmd.MarkFlagRequired("resource-name")
	redshift_deleteTagsCmd.MarkFlagRequired("tag-keys")
	redshiftCmd.AddCommand(redshift_deleteTagsCmd)
}
