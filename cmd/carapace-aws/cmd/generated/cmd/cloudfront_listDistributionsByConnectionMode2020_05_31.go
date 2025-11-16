package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listDistributionsByConnectionMode2020_05_31Cmd = &cobra.Command{
	Use:   "list-distributions-by-connection-mode2020_05_31",
	Short: "Lists the distributions by the connection mode that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listDistributionsByConnectionMode2020_05_31Cmd).Standalone()

	cloudfront_listDistributionsByConnectionMode2020_05_31Cmd.Flags().String("connection-mode", "", "This field specifies whether the connection mode is through a standard distribution (direct) or a multi-tenant distribution with distribution tenants (tenant-only).")
	cloudfront_listDistributionsByConnectionMode2020_05_31Cmd.Flags().String("marker", "", "The marker for the next set of distributions to retrieve.")
	cloudfront_listDistributionsByConnectionMode2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of distributions to return.")
	cloudfront_listDistributionsByConnectionMode2020_05_31Cmd.MarkFlagRequired("connection-mode")
	cloudfrontCmd.AddCommand(cloudfront_listDistributionsByConnectionMode2020_05_31Cmd)
}
