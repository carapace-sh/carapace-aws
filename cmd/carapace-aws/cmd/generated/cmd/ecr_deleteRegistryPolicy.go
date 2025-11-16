package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_deleteRegistryPolicyCmd = &cobra.Command{
	Use:   "delete-registry-policy",
	Short: "Deletes the registry permissions policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_deleteRegistryPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_deleteRegistryPolicyCmd).Standalone()

	})
	ecrCmd.AddCommand(ecr_deleteRegistryPolicyCmd)
}
