package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_listLoggingConfigurationsCmd = &cobra.Command{
	Use:   "list-logging-configurations",
	Short: "Retrieves an array of your [LoggingConfiguration]() objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_listLoggingConfigurationsCmd).Standalone()

	wafv2_listLoggingConfigurationsCmd.Flags().String("limit", "", "The maximum number of objects that you want WAF to return for this request.")
	wafv2_listLoggingConfigurationsCmd.Flags().String("log-scope", "", "The owner of the logging configuration, which must be set to `CUSTOMER` for the configurations that you manage.")
	wafv2_listLoggingConfigurationsCmd.Flags().String("next-marker", "", "When you request a list of objects with a `Limit` setting, if the number of objects that are still available for retrieval exceeds the limit, WAF returns a `NextMarker` value in the response.")
	wafv2_listLoggingConfigurationsCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
	wafv2_listLoggingConfigurationsCmd.MarkFlagRequired("scope")
	wafv2Cmd.AddCommand(wafv2_listLoggingConfigurationsCmd)
}
