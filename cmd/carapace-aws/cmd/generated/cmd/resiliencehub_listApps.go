package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listAppsCmd = &cobra.Command{
	Use:   "list-apps",
	Short: "Lists your Resilience Hub applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listAppsCmd).Standalone()

	resiliencehub_listAppsCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_listAppsCmd.Flags().String("aws-application-arn", "", "Amazon Resource Name (ARN) of Resource Groups group that is integrated with an AppRegistry application.")
	resiliencehub_listAppsCmd.Flags().String("from-last-assessment-time", "", "Lower limit of the range that is used to filter applications based on their last assessment times.")
	resiliencehub_listAppsCmd.Flags().String("max-results", "", "Maximum number of results to include in the response.")
	resiliencehub_listAppsCmd.Flags().String("name", "", "The name for the one of the listed applications.")
	resiliencehub_listAppsCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
	resiliencehub_listAppsCmd.Flags().String("reverse-order", "", "The application list is sorted based on the values of `lastAppComplianceEvaluationTime` field.")
	resiliencehub_listAppsCmd.Flags().String("to-last-assessment-time", "", "Upper limit of the range that is used to filter the applications based on their last assessment times.")
	resiliencehubCmd.AddCommand(resiliencehub_listAppsCmd)
}
