package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_listAutoScalingConfigurationsCmd = &cobra.Command{
	Use:   "list-auto-scaling-configurations",
	Short: "Returns a list of active App Runner automatic scaling configurations in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_listAutoScalingConfigurationsCmd).Standalone()

	apprunner_listAutoScalingConfigurationsCmd.Flags().String("auto-scaling-configuration-name", "", "The name of the App Runner auto scaling configuration that you want to list.")
	apprunner_listAutoScalingConfigurationsCmd.Flags().Bool("latest-only", false, "Set to `true` to list only the latest revision for each requested configuration name.")
	apprunner_listAutoScalingConfigurationsCmd.Flags().String("max-results", "", "The maximum number of results to include in each response (result page).")
	apprunner_listAutoScalingConfigurationsCmd.Flags().String("next-token", "", "A token from a previous result page.")
	apprunner_listAutoScalingConfigurationsCmd.Flags().Bool("no-latest-only", false, "Set to `true` to list only the latest revision for each requested configuration name.")
	apprunner_listAutoScalingConfigurationsCmd.Flag("no-latest-only").Hidden = true
	apprunnerCmd.AddCommand(apprunner_listAutoScalingConfigurationsCmd)
}
