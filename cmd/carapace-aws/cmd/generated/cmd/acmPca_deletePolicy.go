package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_deletePolicyCmd = &cobra.Command{
	Use:   "delete-policy",
	Short: "Deletes the resource-based policy attached to a private CA.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_deletePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(acmPca_deletePolicyCmd).Standalone()

		acmPca_deletePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Number (ARN) of the private CA that will have its policy deleted.")
		acmPca_deletePolicyCmd.MarkFlagRequired("resource-arn")
	})
	acmPcaCmd.AddCommand(acmPca_deletePolicyCmd)
}
