package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_exportSourceNetworkCfnTemplateCmd = &cobra.Command{
	Use:   "export-source-network-cfn-template",
	Short: "Export the Source Network CloudFormation template to an S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_exportSourceNetworkCfnTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_exportSourceNetworkCfnTemplateCmd).Standalone()

		drs_exportSourceNetworkCfnTemplateCmd.Flags().String("source-network-id", "", "The Source Network ID to export its CloudFormation template to an S3 bucket.")
		drs_exportSourceNetworkCfnTemplateCmd.MarkFlagRequired("source-network-id")
	})
	drsCmd.AddCommand(drs_exportSourceNetworkCfnTemplateCmd)
}
