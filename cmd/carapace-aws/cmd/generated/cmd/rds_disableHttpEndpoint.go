package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_disableHttpEndpointCmd = &cobra.Command{
	Use:   "disable-http-endpoint",
	Short: "Disables the HTTP endpoint for the specified DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_disableHttpEndpointCmd).Standalone()

	rds_disableHttpEndpointCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the DB cluster.")
	rds_disableHttpEndpointCmd.MarkFlagRequired("resource-arn")
	rdsCmd.AddCommand(rds_disableHttpEndpointCmd)
}
