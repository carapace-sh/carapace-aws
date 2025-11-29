package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed s3/aws.s3.cp.yaml
var s3_cp []byte

//go:embed s3/aws.s3.ls.yaml
var s3_ls []byte

//go:embed s3/aws.s3.mb.yaml
var s3_mb []byte

//go:embed s3/aws.s3.mv.yaml
var s3_mv []byte

//go:embed s3/aws.s3.presign.yaml
var s3_presign []byte

//go:embed s3/aws.s3.rb.yaml
var s3_rb []byte

//go:embed s3/aws.s3.rm.yaml
var s3_rm []byte

//go:embed s3/aws.s3.sync.yaml
var s3_sync []byte

//go:embed s3/aws.s3.website.yaml
var s3_website []byte

func init() {
	customizations["s3"] = func(cmd *command.Command) error {
		cmd.Commands = make([]command.Command, 0)
		for _, spec := range [][]byte{
			s3_cp,
			s3_ls,
			s3_mb,
			s3_mv,
			s3_presign,
			s3_rb,
			s3_rm,
			s3_sync,
			s3_website,
		} {
			var specCommand command.Command
			if err := yaml.Unmarshal(spec, &specCommand); err != nil {
				return err
			}
			cmd.Commands = append(cmd.Commands, specCommand)
		}
		return nil
	}
}
