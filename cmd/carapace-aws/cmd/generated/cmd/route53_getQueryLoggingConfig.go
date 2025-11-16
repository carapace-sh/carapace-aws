package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_getQueryLoggingConfigCmd = &cobra.Command{
	Use:   "get-query-logging-config",
	Short: "Gets information about a specified configuration for DNS query logging.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_getQueryLoggingConfigCmd).Standalone()

	route53_getQueryLoggingConfigCmd.Flags().String("id", "", "The ID of the configuration for DNS query logging that you want to get information about.")
	route53_getQueryLoggingConfigCmd.MarkFlagRequired("id")
	route53Cmd.AddCommand(route53_getQueryLoggingConfigCmd)
}
