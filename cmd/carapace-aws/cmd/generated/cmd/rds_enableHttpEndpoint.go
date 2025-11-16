package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_enableHttpEndpointCmd = &cobra.Command{
	Use:   "enable-http-endpoint",
	Short: "Enables the HTTP endpoint for the DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_enableHttpEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_enableHttpEndpointCmd).Standalone()

		rds_enableHttpEndpointCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the DB cluster.")
		rds_enableHttpEndpointCmd.MarkFlagRequired("resource-arn")
	})
	rdsCmd.AddCommand(rds_enableHttpEndpointCmd)
}
