package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_disablePolicyTypeCmd = &cobra.Command{
	Use:   "disable-policy-type",
	Short: "Disables an organizational policy type in a root.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_disablePolicyTypeCmd).Standalone()

	organizations_disablePolicyTypeCmd.Flags().String("policy-type", "", "The policy type that you want to disable in this root.")
	organizations_disablePolicyTypeCmd.Flags().String("root-id", "", "The unique identifier (ID) of the root in which you want to disable a policy type.")
	organizations_disablePolicyTypeCmd.MarkFlagRequired("policy-type")
	organizations_disablePolicyTypeCmd.MarkFlagRequired("root-id")
	organizationsCmd.AddCommand(organizations_disablePolicyTypeCmd)
}
