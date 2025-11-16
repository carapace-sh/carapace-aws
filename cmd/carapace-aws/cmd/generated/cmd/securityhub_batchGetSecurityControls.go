package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_batchGetSecurityControlsCmd = &cobra.Command{
	Use:   "batch-get-security-controls",
	Short: "Provides details about a batch of security controls for the current Amazon Web Services account and Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_batchGetSecurityControlsCmd).Standalone()

	securityhub_batchGetSecurityControlsCmd.Flags().String("security-control-ids", "", "A list of security controls (identified with `SecurityControlId`, `SecurityControlArn`, or a mix of both parameters).")
	securityhub_batchGetSecurityControlsCmd.MarkFlagRequired("security-control-ids")
	securityhubCmd.AddCommand(securityhub_batchGetSecurityControlsCmd)
}
