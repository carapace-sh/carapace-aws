package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_describeUserStackAssociationsCmd = &cobra.Command{
	Use:   "describe-user-stack-associations",
	Short: "Retrieves a list that describes the UserStackAssociation objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_describeUserStackAssociationsCmd).Standalone()

	appstream_describeUserStackAssociationsCmd.Flags().String("authentication-type", "", "The authentication type for the user who is associated with the stack.")
	appstream_describeUserStackAssociationsCmd.Flags().String("max-results", "", "The maximum size of each page of results.")
	appstream_describeUserStackAssociationsCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	appstream_describeUserStackAssociationsCmd.Flags().String("stack-name", "", "The name of the stack that is associated with the user.")
	appstream_describeUserStackAssociationsCmd.Flags().String("user-name", "", "The email address of the user who is associated with the stack.")
	appstreamCmd.AddCommand(appstream_describeUserStackAssociationsCmd)
}
