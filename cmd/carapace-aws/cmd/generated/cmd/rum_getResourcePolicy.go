package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Use this operation to retrieve information about a resource-based policy that is attached to an app monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_getResourcePolicyCmd).Standalone()

	rum_getResourcePolicyCmd.Flags().String("name", "", "The name of the app monitor that is associated with the resource-based policy that you want to view.")
	rum_getResourcePolicyCmd.MarkFlagRequired("name")
	rumCmd.AddCommand(rum_getResourcePolicyCmd)
}
