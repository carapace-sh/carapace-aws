package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_describeEffectivePolicyCmd = &cobra.Command{
	Use:   "describe-effective-policy",
	Short: "Returns the contents of the effective policy for specified policy type and account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_describeEffectivePolicyCmd).Standalone()

	organizations_describeEffectivePolicyCmd.Flags().String("policy-type", "", "The type of policy that you want information about.")
	organizations_describeEffectivePolicyCmd.Flags().String("target-id", "", "When you're signed in as the management account, specify the ID of the account that you want details about.")
	organizations_describeEffectivePolicyCmd.MarkFlagRequired("policy-type")
	organizationsCmd.AddCommand(organizations_describeEffectivePolicyCmd)
}
