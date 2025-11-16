package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_deleteAppVersionResourceCmd = &cobra.Command{
	Use:   "delete-app-version-resource",
	Short: "Deletes a resource from the Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_deleteAppVersionResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_deleteAppVersionResourceCmd).Standalone()

		resiliencehub_deleteAppVersionResourceCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
		resiliencehub_deleteAppVersionResourceCmd.Flags().String("aws-account-id", "", "Amazon Web Services account that owns the physical resource.")
		resiliencehub_deleteAppVersionResourceCmd.Flags().String("aws-region", "", "Amazon Web Services region that owns the physical resource.")
		resiliencehub_deleteAppVersionResourceCmd.Flags().String("client-token", "", "Used for an idempotency token.")
		resiliencehub_deleteAppVersionResourceCmd.Flags().String("logical-resource-id", "", "Logical identifier of the resource.")
		resiliencehub_deleteAppVersionResourceCmd.Flags().String("physical-resource-id", "", "Physical identifier of the resource.")
		resiliencehub_deleteAppVersionResourceCmd.Flags().String("resource-name", "", "Name of the resource.")
		resiliencehub_deleteAppVersionResourceCmd.MarkFlagRequired("app-arn")
	})
	resiliencehubCmd.AddCommand(resiliencehub_deleteAppVersionResourceCmd)
}
