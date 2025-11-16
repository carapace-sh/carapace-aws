package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_describeAppCmd = &cobra.Command{
	Use:   "describe-app",
	Short: "Describes an Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_describeAppCmd).Standalone()

	resiliencehub_describeAppCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_describeAppCmd.MarkFlagRequired("app-arn")
	resiliencehubCmd.AddCommand(resiliencehub_describeAppCmd)
}
