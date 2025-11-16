package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_putAccessPointScopeCmd = &cobra.Command{
	Use:   "put-access-point-scope",
	Short: "Creates or replaces the access point scope for a directory bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_putAccessPointScopeCmd).Standalone()

	s3control_putAccessPointScopeCmd.Flags().String("account-id", "", "The Amazon Web Services account ID that owns the access point with scope that you want to create or replace.")
	s3control_putAccessPointScopeCmd.Flags().String("name", "", "The name of the access point with the scope that you want to create or replace.")
	s3control_putAccessPointScopeCmd.Flags().String("scope", "", "Object prefixes, API operations, or a combination of both.")
	s3control_putAccessPointScopeCmd.MarkFlagRequired("account-id")
	s3control_putAccessPointScopeCmd.MarkFlagRequired("name")
	s3control_putAccessPointScopeCmd.MarkFlagRequired("scope")
	s3controlCmd.AddCommand(s3control_putAccessPointScopeCmd)
}
