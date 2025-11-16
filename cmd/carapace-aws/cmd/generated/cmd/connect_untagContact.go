package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_untagContactCmd = &cobra.Command{
	Use:   "untag-contact",
	Short: "Removes the specified tags from the contact resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_untagContactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_untagContactCmd).Standalone()

		connect_untagContactCmd.Flags().String("contact-id", "", "The identifier of the contact in this instance of Amazon Connect.")
		connect_untagContactCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_untagContactCmd.Flags().String("tag-keys", "", "A list of tag keys.")
		connect_untagContactCmd.MarkFlagRequired("contact-id")
		connect_untagContactCmd.MarkFlagRequired("instance-id")
		connect_untagContactCmd.MarkFlagRequired("tag-keys")
	})
	connectCmd.AddCommand(connect_untagContactCmd)
}
