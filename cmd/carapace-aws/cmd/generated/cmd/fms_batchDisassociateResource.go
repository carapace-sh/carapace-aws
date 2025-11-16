package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_batchDisassociateResourceCmd = &cobra.Command{
	Use:   "batch-disassociate-resource",
	Short: "Disassociates resources from a Firewall Manager resource set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_batchDisassociateResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_batchDisassociateResourceCmd).Standalone()

		fms_batchDisassociateResourceCmd.Flags().String("items", "", "The uniform resource identifiers (URI) of resources that should be disassociated from the resource set.")
		fms_batchDisassociateResourceCmd.Flags().String("resource-set-identifier", "", "A unique identifier for the resource set, used in a request to refer to the resource set.")
		fms_batchDisassociateResourceCmd.MarkFlagRequired("items")
		fms_batchDisassociateResourceCmd.MarkFlagRequired("resource-set-identifier")
	})
	fmsCmd.AddCommand(fms_batchDisassociateResourceCmd)
}
