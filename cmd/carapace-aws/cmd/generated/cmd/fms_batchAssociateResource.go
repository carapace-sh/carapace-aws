package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_batchAssociateResourceCmd = &cobra.Command{
	Use:   "batch-associate-resource",
	Short: "Associate resources to a Firewall Manager resource set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_batchAssociateResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_batchAssociateResourceCmd).Standalone()

		fms_batchAssociateResourceCmd.Flags().String("items", "", "The uniform resource identifiers (URIs) of resources that should be associated to the resource set.")
		fms_batchAssociateResourceCmd.Flags().String("resource-set-identifier", "", "A unique identifier for the resource set, used in a request to refer to the resource set.")
		fms_batchAssociateResourceCmd.MarkFlagRequired("items")
		fms_batchAssociateResourceCmd.MarkFlagRequired("resource-set-identifier")
	})
	fmsCmd.AddCommand(fms_batchAssociateResourceCmd)
}
