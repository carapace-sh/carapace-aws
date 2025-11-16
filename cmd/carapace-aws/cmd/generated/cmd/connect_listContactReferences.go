package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listContactReferencesCmd = &cobra.Command{
	Use:   "list-contact-references",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listContactReferencesCmd).Standalone()

	connect_listContactReferencesCmd.Flags().String("contact-id", "", "The identifier of the initial contact.")
	connect_listContactReferencesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_listContactReferencesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_listContactReferencesCmd.Flags().String("reference-types", "", "The type of reference.")
	connect_listContactReferencesCmd.MarkFlagRequired("contact-id")
	connect_listContactReferencesCmd.MarkFlagRequired("instance-id")
	connect_listContactReferencesCmd.MarkFlagRequired("reference-types")
	connectCmd.AddCommand(connect_listContactReferencesCmd)
}
