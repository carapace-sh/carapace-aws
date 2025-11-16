package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_deleteAccessPointCmd = &cobra.Command{
	Use:   "delete-access-point",
	Short: "Deletes the specified access point.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_deleteAccessPointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_deleteAccessPointCmd).Standalone()

		s3control_deleteAccessPointCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the account that owns the specified access point.")
		s3control_deleteAccessPointCmd.Flags().String("name", "", "The name of the access point you want to delete.")
		s3control_deleteAccessPointCmd.MarkFlagRequired("account-id")
		s3control_deleteAccessPointCmd.MarkFlagRequired("name")
	})
	s3controlCmd.AddCommand(s3control_deleteAccessPointCmd)
}
