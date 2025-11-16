package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_describePolicyCmd = &cobra.Command{
	Use:   "describe-policy",
	Short: "Retrieves information about a policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_describePolicyCmd).Standalone()

	organizations_describePolicyCmd.Flags().String("policy-id", "", "The unique identifier (ID) of the policy that you want details about.")
	organizations_describePolicyCmd.MarkFlagRequired("policy-id")
	organizationsCmd.AddCommand(organizations_describePolicyCmd)
}
