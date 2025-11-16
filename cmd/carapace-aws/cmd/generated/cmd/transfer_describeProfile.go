package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_describeProfileCmd = &cobra.Command{
	Use:   "describe-profile",
	Short: "Returns the details of the profile that's specified by the `ProfileId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_describeProfileCmd).Standalone()

	transfer_describeProfileCmd.Flags().String("profile-id", "", "The identifier of the profile that you want described.")
	transfer_describeProfileCmd.MarkFlagRequired("profile-id")
	transferCmd.AddCommand(transfer_describeProfileCmd)
}
