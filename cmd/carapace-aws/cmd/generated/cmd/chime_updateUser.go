package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_updateUserCmd = &cobra.Command{
	Use:   "update-user",
	Short: "Updates user details for a specified user ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_updateUserCmd).Standalone()

	chime_updateUserCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_updateUserCmd.Flags().String("alexa-for-business-metadata", "", "The Alexa for Business metadata.")
	chime_updateUserCmd.Flags().String("license-type", "", "The user license type to update.")
	chime_updateUserCmd.Flags().String("user-id", "", "The user ID.")
	chime_updateUserCmd.Flags().String("user-type", "", "The user type.")
	chime_updateUserCmd.MarkFlagRequired("account-id")
	chime_updateUserCmd.MarkFlagRequired("user-id")
	chimeCmd.AddCommand(chime_updateUserCmd)
}
