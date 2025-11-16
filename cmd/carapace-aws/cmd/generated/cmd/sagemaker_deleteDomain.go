package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteDomainCmd = &cobra.Command{
	Use:   "delete-domain",
	Short: "Used to delete a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteDomainCmd).Standalone()

	sagemaker_deleteDomainCmd.Flags().String("domain-id", "", "The domain ID.")
	sagemaker_deleteDomainCmd.Flags().String("retention-policy", "", "The retention policy for this domain, which specifies whether resources will be retained after the Domain is deleted.")
	sagemaker_deleteDomainCmd.MarkFlagRequired("domain-id")
	sagemakerCmd.AddCommand(sagemaker_deleteDomainCmd)
}
