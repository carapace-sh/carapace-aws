package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_disassociateTrustStoreCmd = &cobra.Command{
	Use:   "disassociate-trust-store",
	Short: "Disassociates a trust store from a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_disassociateTrustStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_disassociateTrustStoreCmd).Standalone()

		workspacesWeb_disassociateTrustStoreCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
		workspacesWeb_disassociateTrustStoreCmd.MarkFlagRequired("portal-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_disassociateTrustStoreCmd)
}
