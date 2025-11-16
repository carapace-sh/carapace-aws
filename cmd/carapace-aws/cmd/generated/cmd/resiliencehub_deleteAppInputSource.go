package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_deleteAppInputSourceCmd = &cobra.Command{
	Use:   "delete-app-input-source",
	Short: "Deletes the input source and all of its imported resources from the Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_deleteAppInputSourceCmd).Standalone()

	resiliencehub_deleteAppInputSourceCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_deleteAppInputSourceCmd.Flags().String("client-token", "", "Used for an idempotency token.")
	resiliencehub_deleteAppInputSourceCmd.Flags().String("eks-source-cluster-namespace", "", "The namespace on your Amazon Elastic Kubernetes Service cluster that you want to delete from the Resilience Hub application.")
	resiliencehub_deleteAppInputSourceCmd.Flags().String("source-arn", "", "The Amazon Resource Name (ARN) of the imported resource you want to remove from the Resilience Hub application.")
	resiliencehub_deleteAppInputSourceCmd.Flags().String("terraform-source", "", "The imported Terraform s3 state ﬁle you want to remove from the Resilience Hub application.")
	resiliencehub_deleteAppInputSourceCmd.MarkFlagRequired("app-arn")
	resiliencehubCmd.AddCommand(resiliencehub_deleteAppInputSourceCmd)
}
