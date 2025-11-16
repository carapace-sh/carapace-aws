package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_listLoggingConfigurationsCmd = &cobra.Command{
	Use:   "list-logging-configurations",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_listLoggingConfigurationsCmd).Standalone()

	wafRegional_listLoggingConfigurationsCmd.Flags().String("limit", "", "Specifies the number of `LoggingConfigurations` that you want AWS WAF to return for this request.")
	wafRegional_listLoggingConfigurationsCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more `LoggingConfigurations` than the value of `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `LoggingConfigurations`.")
	wafRegionalCmd.AddCommand(wafRegional_listLoggingConfigurationsCmd)
}
