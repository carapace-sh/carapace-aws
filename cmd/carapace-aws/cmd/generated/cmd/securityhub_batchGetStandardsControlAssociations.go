package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_batchGetStandardsControlAssociationsCmd = &cobra.Command{
	Use:   "batch-get-standards-control-associations",
	Short: "For a batch of security controls and standards, identifies whether each control is currently enabled or disabled in a standard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_batchGetStandardsControlAssociationsCmd).Standalone()

	securityhub_batchGetStandardsControlAssociationsCmd.Flags().String("standards-control-association-ids", "", "An array with one or more objects that includes a security control (identified with `SecurityControlId`, `SecurityControlArn`, or a mix of both parameters) and the Amazon Resource Name (ARN) of a standard.")
	securityhub_batchGetStandardsControlAssociationsCmd.MarkFlagRequired("standards-control-association-ids")
	securityhubCmd.AddCommand(securityhub_batchGetStandardsControlAssociationsCmd)
}
