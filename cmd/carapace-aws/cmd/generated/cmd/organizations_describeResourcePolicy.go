package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_describeResourcePolicyCmd = &cobra.Command{
	Use:   "describe-resource-policy",
	Short: "Retrieves information about a resource policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_describeResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_describeResourcePolicyCmd).Standalone()

	})
	organizationsCmd.AddCommand(organizations_describeResourcePolicyCmd)
}
