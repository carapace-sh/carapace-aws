package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_describePageCmd = &cobra.Command{
	Use:   "describe-page",
	Short: "Lists details of the engagement to a contact channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_describePageCmd).Standalone()

	ssmContacts_describePageCmd.Flags().String("page-id", "", "The ID of the engagement to a contact channel.")
	ssmContacts_describePageCmd.MarkFlagRequired("page-id")
	ssmContactsCmd.AddCommand(ssmContacts_describePageCmd)
}
