package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_updateResolverTypeCmd = &cobra.Command{
	Use:   "update-resolver-type",
	Short: "Updates the resolver type for a case.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_updateResolverTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityIr_updateResolverTypeCmd).Standalone()

		securityIr_updateResolverTypeCmd.Flags().String("case-id", "", "Required element for UpdateResolverType to identify the case to update.")
		securityIr_updateResolverTypeCmd.Flags().String("resolver-type", "", "Required element for UpdateResolverType to identify the new resolver.")
		securityIr_updateResolverTypeCmd.MarkFlagRequired("case-id")
		securityIr_updateResolverTypeCmd.MarkFlagRequired("resolver-type")
	})
	securityIrCmd.AddCommand(securityIr_updateResolverTypeCmd)
}
