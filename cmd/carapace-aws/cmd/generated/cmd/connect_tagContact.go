package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_tagContactCmd = &cobra.Command{
	Use:   "tag-contact",
	Short: "Adds the specified tags to the contact resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_tagContactCmd).Standalone()

	connect_tagContactCmd.Flags().String("contact-id", "", "The identifier of the contact in this instance of Amazon Connect.")
	connect_tagContactCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_tagContactCmd.Flags().String("tags", "", "The tags to be assigned to the contact resource.")
	connect_tagContactCmd.MarkFlagRequired("contact-id")
	connect_tagContactCmd.MarkFlagRequired("instance-id")
	connect_tagContactCmd.MarkFlagRequired("tags")
	connectCmd.AddCommand(connect_tagContactCmd)
}
