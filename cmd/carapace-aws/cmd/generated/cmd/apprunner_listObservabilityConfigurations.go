package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_listObservabilityConfigurationsCmd = &cobra.Command{
	Use:   "list-observability-configurations",
	Short: "Returns a list of active App Runner observability configurations in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_listObservabilityConfigurationsCmd).Standalone()

	apprunner_listObservabilityConfigurationsCmd.Flags().Bool("latest-only", false, "Set to `true` to list only the latest revision for each requested configuration name.")
	apprunner_listObservabilityConfigurationsCmd.Flags().String("max-results", "", "The maximum number of results to include in each response (result page).")
	apprunner_listObservabilityConfigurationsCmd.Flags().String("next-token", "", "A token from a previous result page.")
	apprunner_listObservabilityConfigurationsCmd.Flags().Bool("no-latest-only", false, "Set to `true` to list only the latest revision for each requested configuration name.")
	apprunner_listObservabilityConfigurationsCmd.Flags().String("observability-configuration-name", "", "The name of the App Runner observability configuration that you want to list.")
	apprunner_listObservabilityConfigurationsCmd.Flag("no-latest-only").Hidden = true
	apprunnerCmd.AddCommand(apprunner_listObservabilityConfigurationsCmd)
}
