package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_disassociateApiCmd = &cobra.Command{
	Use:   "disassociate-api",
	Short: "Removes an `ApiAssociation` object from a custom domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_disassociateApiCmd).Standalone()

	appsync_disassociateApiCmd.Flags().String("domain-name", "", "The domain name.")
	appsync_disassociateApiCmd.MarkFlagRequired("domain-name")
	appsyncCmd.AddCommand(appsync_disassociateApiCmd)
}
