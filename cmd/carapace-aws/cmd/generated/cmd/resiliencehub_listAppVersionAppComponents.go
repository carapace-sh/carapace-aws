package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listAppVersionAppComponentsCmd = &cobra.Command{
	Use:   "list-app-version-app-components",
	Short: "Lists all the Application Components in the Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listAppVersionAppComponentsCmd).Standalone()

	resiliencehub_listAppVersionAppComponentsCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_listAppVersionAppComponentsCmd.Flags().String("app-version", "", "Version of the Application Component.")
	resiliencehub_listAppVersionAppComponentsCmd.Flags().String("max-results", "", "Maximum number of Application Components to be displayed per Resilience Hub application version.")
	resiliencehub_listAppVersionAppComponentsCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
	resiliencehub_listAppVersionAppComponentsCmd.MarkFlagRequired("app-arn")
	resiliencehub_listAppVersionAppComponentsCmd.MarkFlagRequired("app-version")
	resiliencehubCmd.AddCommand(resiliencehub_listAppVersionAppComponentsCmd)
}
