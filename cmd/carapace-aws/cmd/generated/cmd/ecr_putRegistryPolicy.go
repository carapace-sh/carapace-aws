package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_putRegistryPolicyCmd = &cobra.Command{
	Use:   "put-registry-policy",
	Short: "Creates or updates the permissions policy for your registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_putRegistryPolicyCmd).Standalone()

	ecr_putRegistryPolicyCmd.Flags().String("policy-text", "", "The JSON policy text to apply to your registry.")
	ecr_putRegistryPolicyCmd.MarkFlagRequired("policy-text")
	ecrCmd.AddCommand(ecr_putRegistryPolicyCmd)
}
