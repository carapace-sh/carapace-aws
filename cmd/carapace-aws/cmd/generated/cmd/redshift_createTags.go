package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_createTagsCmd = &cobra.Command{
	Use:   "create-tags",
	Short: "Adds tags to a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_createTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_createTagsCmd).Standalone()

		redshift_createTagsCmd.Flags().String("resource-name", "", "The Amazon Resource Name (ARN) to which you want to add the tag or tags.")
		redshift_createTagsCmd.Flags().String("tags", "", "One or more name/value pairs to add as tags to the specified resource.")
		redshift_createTagsCmd.MarkFlagRequired("resource-name")
		redshift_createTagsCmd.MarkFlagRequired("tags")
	})
	redshiftCmd.AddCommand(redshift_createTagsCmd)
}
