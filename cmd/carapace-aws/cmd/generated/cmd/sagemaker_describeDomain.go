package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeDomainCmd = &cobra.Command{
	Use:   "describe-domain",
	Short: "The description of the domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeDomainCmd).Standalone()

	sagemaker_describeDomainCmd.Flags().String("domain-id", "", "The domain ID.")
	sagemaker_describeDomainCmd.MarkFlagRequired("domain-id")
	sagemakerCmd.AddCommand(sagemaker_describeDomainCmd)
}
