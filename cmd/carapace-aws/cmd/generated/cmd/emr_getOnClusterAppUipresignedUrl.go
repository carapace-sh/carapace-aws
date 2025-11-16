package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_getOnClusterAppUipresignedUrlCmd = &cobra.Command{
	Use:   "get-on-cluster-app-uipresigned-url",
	Short: "The presigned URL properties for the cluster's application user interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_getOnClusterAppUipresignedUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_getOnClusterAppUipresignedUrlCmd).Standalone()

		emr_getOnClusterAppUipresignedUrlCmd.Flags().String("application-id", "", "The application ID associated with the cluster's application user interface presigned URL.")
		emr_getOnClusterAppUipresignedUrlCmd.Flags().String("cluster-id", "", "The cluster ID associated with the cluster's application user interface presigned URL.")
		emr_getOnClusterAppUipresignedUrlCmd.Flags().String("dry-run", "", "Determines if the user interface presigned URL is for a dry run.")
		emr_getOnClusterAppUipresignedUrlCmd.Flags().String("execution-role-arn", "", "The execution role ARN associated with the cluster's application user interface presigned URL.")
		emr_getOnClusterAppUipresignedUrlCmd.Flags().String("on-cluster-app-uitype", "", "The application UI type associated with the cluster's application user interface presigned URL.")
		emr_getOnClusterAppUipresignedUrlCmd.MarkFlagRequired("cluster-id")
	})
	emrCmd.AddCommand(emr_getOnClusterAppUipresignedUrlCmd)
}
