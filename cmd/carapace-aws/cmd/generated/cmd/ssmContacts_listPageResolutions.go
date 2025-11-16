package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_listPageResolutionsCmd = &cobra.Command{
	Use:   "list-page-resolutions",
	Short: "Returns the resolution path of an engagement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_listPageResolutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_listPageResolutionsCmd).Standalone()

		ssmContacts_listPageResolutionsCmd.Flags().String("next-token", "", "A token to start the list.")
		ssmContacts_listPageResolutionsCmd.Flags().String("page-id", "", "The Amazon Resource Name (ARN) of the contact engaged for the incident.")
		ssmContacts_listPageResolutionsCmd.MarkFlagRequired("page-id")
	})
	ssmContactsCmd.AddCommand(ssmContacts_listPageResolutionsCmd)
}
