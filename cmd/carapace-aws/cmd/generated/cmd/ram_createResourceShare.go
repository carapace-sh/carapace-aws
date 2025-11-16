package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_createResourceShareCmd = &cobra.Command{
	Use:   "create-resource-share",
	Short: "Creates a resource share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_createResourceShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_createResourceShareCmd).Standalone()

		ram_createResourceShareCmd.Flags().Bool("allow-external-principals", false, "Specifies whether principals outside your organization in Organizations can be associated with a resource share.")
		ram_createResourceShareCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ram_createResourceShareCmd.Flags().String("name", "", "Specifies the name of the resource share.")
		ram_createResourceShareCmd.Flags().Bool("no-allow-external-principals", false, "Specifies whether principals outside your organization in Organizations can be associated with a resource share.")
		ram_createResourceShareCmd.Flags().String("permission-arns", "", "Specifies the [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the RAM permission to associate with the resource share.")
		ram_createResourceShareCmd.Flags().String("principals", "", "Specifies a list of one or more principals to associate with the resource share.")
		ram_createResourceShareCmd.Flags().String("resource-arns", "", "Specifies a list of one or more ARNs of the resources to associate with the resource share.")
		ram_createResourceShareCmd.Flags().String("sources", "", "Specifies from which source accounts the service principal has access to the resources in this resource share.")
		ram_createResourceShareCmd.Flags().String("tags", "", "Specifies one or more tags to attach to the resource share itself.")
		ram_createResourceShareCmd.MarkFlagRequired("name")
		ram_createResourceShareCmd.Flag("no-allow-external-principals").Hidden = true
	})
	ramCmd.AddCommand(ram_createResourceShareCmd)
}
