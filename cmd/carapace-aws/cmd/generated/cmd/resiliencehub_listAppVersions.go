package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listAppVersionsCmd = &cobra.Command{
	Use:   "list-app-versions",
	Short: "Lists the different versions for the Resilience Hub applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listAppVersionsCmd).Standalone()

	resiliencehub_listAppVersionsCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_listAppVersionsCmd.Flags().String("end-time", "", "Upper limit of the time range to filter the application versions.")
	resiliencehub_listAppVersionsCmd.Flags().String("max-results", "", "Maximum number of results to include in the response.")
	resiliencehub_listAppVersionsCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
	resiliencehub_listAppVersionsCmd.Flags().String("start-time", "", "Lower limit of the time range to filter the application versions.")
	resiliencehub_listAppVersionsCmd.MarkFlagRequired("app-arn")
	resiliencehubCmd.AddCommand(resiliencehub_listAppVersionsCmd)
}
