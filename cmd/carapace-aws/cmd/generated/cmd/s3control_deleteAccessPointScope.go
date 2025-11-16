package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_deleteAccessPointScopeCmd = &cobra.Command{
	Use:   "delete-access-point-scope",
	Short: "Deletes an existing access point scope for a directory bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_deleteAccessPointScopeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_deleteAccessPointScopeCmd).Standalone()

		s3control_deleteAccessPointScopeCmd.Flags().String("account-id", "", "The Amazon Web Services account ID that owns the access point with the scope that you want to delete.")
		s3control_deleteAccessPointScopeCmd.Flags().String("name", "", "The name of the access point with the scope that you want to delete.")
		s3control_deleteAccessPointScopeCmd.MarkFlagRequired("account-id")
		s3control_deleteAccessPointScopeCmd.MarkFlagRequired("name")
	})
	s3controlCmd.AddCommand(s3control_deleteAccessPointScopeCmd)
}
