package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_importResourcesToDraftAppVersionCmd = &cobra.Command{
	Use:   "import-resources-to-draft-app-version",
	Short: "Imports resources to Resilience Hub application draft version from different input sources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_importResourcesToDraftAppVersionCmd).Standalone()

	resiliencehub_importResourcesToDraftAppVersionCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_importResourcesToDraftAppVersionCmd.Flags().String("eks-sources", "", "The input sources of the Amazon Elastic Kubernetes Service resources you need to import.")
	resiliencehub_importResourcesToDraftAppVersionCmd.Flags().String("import-strategy", "", "The import strategy you would like to set to import resources into Resilience Hub application.")
	resiliencehub_importResourcesToDraftAppVersionCmd.Flags().String("source-arns", "", "The Amazon Resource Names (ARNs) for the resources.")
	resiliencehub_importResourcesToDraftAppVersionCmd.Flags().String("terraform-sources", "", "A list of terraform file s3 URLs you need to import.")
	resiliencehub_importResourcesToDraftAppVersionCmd.MarkFlagRequired("app-arn")
	resiliencehubCmd.AddCommand(resiliencehub_importResourcesToDraftAppVersionCmd)
}
