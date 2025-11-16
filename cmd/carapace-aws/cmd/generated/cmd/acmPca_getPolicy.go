package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_getPolicyCmd = &cobra.Command{
	Use:   "get-policy",
	Short: "Retrieves the resource-based policy attached to a private CA.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_getPolicyCmd).Standalone()

	acmPca_getPolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Number (ARN) of the private CA that will have its policy retrieved.")
	acmPca_getPolicyCmd.MarkFlagRequired("resource-arn")
	acmPcaCmd.AddCommand(acmPca_getPolicyCmd)
}
