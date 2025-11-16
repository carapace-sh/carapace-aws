package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeTagsCmd = &cobra.Command{
	Use:   "describe-tags",
	Short: "Returns a list of tags.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeTagsCmd).Standalone()

	redshift_describeTagsCmd.Flags().String("marker", "", "A value that indicates the starting point for the next set of response records in a subsequent request.")
	redshift_describeTagsCmd.Flags().String("max-records", "", "The maximum number or response records to return in each call.")
	redshift_describeTagsCmd.Flags().String("resource-name", "", "The Amazon Resource Name (ARN) for which you want to describe the tag or tags.")
	redshift_describeTagsCmd.Flags().String("resource-type", "", "The type of resource with which you want to view tags.")
	redshift_describeTagsCmd.Flags().String("tag-keys", "", "A tag key or keys for which you want to return all matching resources that are associated with the specified key or keys.")
	redshift_describeTagsCmd.Flags().String("tag-values", "", "A tag value or values for which you want to return all matching resources that are associated with the specified value or values.")
	redshiftCmd.AddCommand(redshift_describeTagsCmd)
}
