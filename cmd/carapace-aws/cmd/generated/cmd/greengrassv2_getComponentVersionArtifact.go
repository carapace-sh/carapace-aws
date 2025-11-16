package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_getComponentVersionArtifactCmd = &cobra.Command{
	Use:   "get-component-version-artifact",
	Short: "Gets the pre-signed URL to download a public or a Lambda component artifact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_getComponentVersionArtifactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrassv2_getComponentVersionArtifactCmd).Standalone()

		greengrassv2_getComponentVersionArtifactCmd.Flags().String("arn", "", "The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the component version.")
		greengrassv2_getComponentVersionArtifactCmd.Flags().String("artifact-name", "", "The name of the artifact.")
		greengrassv2_getComponentVersionArtifactCmd.Flags().String("iot-endpoint-type", "", "Determines if the Amazon S3 URL returned is a FIPS pre-signed URL endpoint.")
		greengrassv2_getComponentVersionArtifactCmd.Flags().String("s3-endpoint-type", "", "Specifies the endpoint to use when getting Amazon S3 pre-signed URLs.")
		greengrassv2_getComponentVersionArtifactCmd.MarkFlagRequired("arn")
		greengrassv2_getComponentVersionArtifactCmd.MarkFlagRequired("artifact-name")
	})
	greengrassv2Cmd.AddCommand(greengrassv2_getComponentVersionArtifactCmd)
}
