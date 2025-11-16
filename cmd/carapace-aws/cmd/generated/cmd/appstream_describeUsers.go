package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_describeUsersCmd = &cobra.Command{
	Use:   "describe-users",
	Short: "Retrieves a list that describes one or more specified users in the user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_describeUsersCmd).Standalone()

	appstream_describeUsersCmd.Flags().String("authentication-type", "", "The authentication type for the users in the user pool to describe.")
	appstream_describeUsersCmd.Flags().String("max-results", "", "The maximum size of each page of results.")
	appstream_describeUsersCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	appstream_describeUsersCmd.MarkFlagRequired("authentication-type")
	appstreamCmd.AddCommand(appstream_describeUsersCmd)
}
