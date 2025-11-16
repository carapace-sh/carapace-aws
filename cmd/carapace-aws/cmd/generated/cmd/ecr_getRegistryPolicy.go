package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_getRegistryPolicyCmd = &cobra.Command{
	Use:   "get-registry-policy",
	Short: "Retrieves the permissions policy for a registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_getRegistryPolicyCmd).Standalone()

	ecrCmd.AddCommand(ecr_getRegistryPolicyCmd)
}
