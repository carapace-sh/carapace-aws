package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_deleteQueryLoggingConfigCmd = &cobra.Command{
	Use:   "delete-query-logging-config",
	Short: "Deletes a configuration for DNS query logging.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_deleteQueryLoggingConfigCmd).Standalone()

	route53_deleteQueryLoggingConfigCmd.Flags().String("id", "", "The ID of the configuration that you want to delete.")
	route53_deleteQueryLoggingConfigCmd.MarkFlagRequired("id")
	route53Cmd.AddCommand(route53_deleteQueryLoggingConfigCmd)
}
