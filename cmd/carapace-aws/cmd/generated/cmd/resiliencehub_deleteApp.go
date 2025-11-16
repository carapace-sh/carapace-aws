package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_deleteAppCmd = &cobra.Command{
	Use:   "delete-app",
	Short: "Deletes an Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_deleteAppCmd).Standalone()

	resiliencehub_deleteAppCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_deleteAppCmd.Flags().String("client-token", "", "Used for an idempotency token.")
	resiliencehub_deleteAppCmd.Flags().String("force-delete", "", "A boolean option to force the deletion of an Resilience Hub application.")
	resiliencehub_deleteAppCmd.MarkFlagRequired("app-arn")
	resiliencehubCmd.AddCommand(resiliencehub_deleteAppCmd)
}
