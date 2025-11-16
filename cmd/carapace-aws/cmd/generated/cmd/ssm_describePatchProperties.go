package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describePatchPropertiesCmd = &cobra.Command{
	Use:   "describe-patch-properties",
	Short: "Lists the properties of available patches organized by product, product family, classification, severity, and other properties of available patches.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describePatchPropertiesCmd).Standalone()

	ssm_describePatchPropertiesCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_describePatchPropertiesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	ssm_describePatchPropertiesCmd.Flags().String("operating-system", "", "The operating system type for which to list patches.")
	ssm_describePatchPropertiesCmd.Flags().String("patch-set", "", "Indicates whether to list patches for the Windows operating system or for applications released by Microsoft.")
	ssm_describePatchPropertiesCmd.Flags().String("property", "", "The patch property for which you want to view patch details.")
	ssm_describePatchPropertiesCmd.MarkFlagRequired("operating-system")
	ssm_describePatchPropertiesCmd.MarkFlagRequired("property")
	ssmCmd.AddCommand(ssm_describePatchPropertiesCmd)
}
