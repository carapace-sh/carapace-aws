package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_listStandardsControlAssociationsCmd = &cobra.Command{
	Use:   "list-standards-control-associations",
	Short: "Specifies whether a control is currently enabled or disabled in each enabled standard in the calling account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_listStandardsControlAssociationsCmd).Standalone()

	securityhub_listStandardsControlAssociationsCmd.Flags().String("max-results", "", "An optional parameter that limits the total results of the API response to the specified number.")
	securityhub_listStandardsControlAssociationsCmd.Flags().String("next-token", "", "Optional pagination parameter.")
	securityhub_listStandardsControlAssociationsCmd.Flags().String("security-control-id", "", "The identifier of the control (identified with `SecurityControlId`, `SecurityControlArn`, or a mix of both parameters) that you want to determine the enablement status of in each enabled standard.")
	securityhub_listStandardsControlAssociationsCmd.MarkFlagRequired("security-control-id")
	securityhubCmd.AddCommand(securityhub_listStandardsControlAssociationsCmd)
}
