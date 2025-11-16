package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getAccessPointScopeCmd = &cobra.Command{
	Use:   "get-access-point-scope",
	Short: "Returns the access point scope for a directory bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getAccessPointScopeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_getAccessPointScopeCmd).Standalone()

		s3control_getAccessPointScopeCmd.Flags().String("account-id", "", "The Amazon Web Services account ID that owns the access point with the scope that you want to retrieve.")
		s3control_getAccessPointScopeCmd.Flags().String("name", "", "The name of the access point with the scope you want to retrieve.")
		s3control_getAccessPointScopeCmd.MarkFlagRequired("account-id")
		s3control_getAccessPointScopeCmd.MarkFlagRequired("name")
	})
	s3controlCmd.AddCommand(s3control_getAccessPointScopeCmd)
}
