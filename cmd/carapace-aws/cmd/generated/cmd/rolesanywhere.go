package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhereCmd = &cobra.Command{
	Use:   "rolesanywhere",
	Short: "Identity and Access Management Roles Anywhere provides a secure way for your workloads such as servers, containers, and applications that run outside of Amazon Web Services to obtain temporary Amazon Web Services credentials.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhereCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rolesanywhereCmd).Standalone()

	})
	rootCmd.AddCommand(rolesanywhereCmd)
}
