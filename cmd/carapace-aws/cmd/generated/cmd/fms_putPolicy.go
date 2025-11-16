package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_putPolicyCmd = &cobra.Command{
	Use:   "put-policy",
	Short: "Creates an Firewall Manager policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_putPolicyCmd).Standalone()

	fms_putPolicyCmd.Flags().String("policy", "", "The details of the Firewall Manager policy to be created.")
	fms_putPolicyCmd.Flags().String("tag-list", "", "The tags to add to the Amazon Web Services resource.")
	fms_putPolicyCmd.MarkFlagRequired("policy")
	fmsCmd.AddCommand(fms_putPolicyCmd)
}
