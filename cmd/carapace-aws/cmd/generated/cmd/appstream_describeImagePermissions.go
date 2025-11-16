package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_describeImagePermissionsCmd = &cobra.Command{
	Use:   "describe-image-permissions",
	Short: "Retrieves a list that describes the permissions for shared AWS account IDs on a private image that you own.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_describeImagePermissionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_describeImagePermissionsCmd).Standalone()

		appstream_describeImagePermissionsCmd.Flags().String("max-results", "", "The maximum size of each page of results.")
		appstream_describeImagePermissionsCmd.Flags().String("name", "", "The name of the private image for which to describe permissions.")
		appstream_describeImagePermissionsCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
		appstream_describeImagePermissionsCmd.Flags().String("shared-aws-account-ids", "", "The 12-digit identifier of one or more AWS accounts with which the image is shared.")
		appstream_describeImagePermissionsCmd.MarkFlagRequired("name")
	})
	appstreamCmd.AddCommand(appstream_describeImagePermissionsCmd)
}
