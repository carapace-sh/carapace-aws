package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_updateResourceShareCmd = &cobra.Command{
	Use:   "update-resource-share",
	Short: "Modifies some of the properties of the specified resource share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_updateResourceShareCmd).Standalone()

	ram_updateResourceShareCmd.Flags().Bool("allow-external-principals", false, "Specifies whether principals outside your organization in Organizations can be associated with a resource share.")
	ram_updateResourceShareCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ram_updateResourceShareCmd.Flags().String("name", "", "If specified, the new name that you want to attach to the resource share.")
	ram_updateResourceShareCmd.Flags().Bool("no-allow-external-principals", false, "Specifies whether principals outside your organization in Organizations can be associated with a resource share.")
	ram_updateResourceShareCmd.Flags().String("resource-share-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share that you want to modify.")
	ram_updateResourceShareCmd.Flag("no-allow-external-principals").Hidden = true
	ram_updateResourceShareCmd.MarkFlagRequired("resource-share-arn")
	ramCmd.AddCommand(ram_updateResourceShareCmd)
}
