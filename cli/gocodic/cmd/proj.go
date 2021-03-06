package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/gocodic"
	"github.com/spiegel-im-spiegel/gocodic/options"
	"github.com/spiegel-im-spiegel/gocodic/response"
)

// projCmd represents the proj command
var projCmd = &cobra.Command{
	Use:   "proj",
	Short: "Refer projects API for codic.jp",
	Long:  "Refer projects API for codic.jp",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return os.ErrInvalid
		}
		pid, err := cmd.Flags().GetInt("projectid")
		if err != nil {
			return os.ErrInvalid
		}
		var opts *options.Options
		if pid > 0 {
			opts, err = options.NewOptions(options.CmdProj, viper.GetString("token"))
		} else {
			opts, err = options.NewOptions(options.CmdProjLst, viper.GetString("token"))
		}
		if err != nil {
			return err
		}
		jsonFlag := viper.GetBool("json")

		res, err := gocodic.ReferProjects(opts, options.ProjectID(pid))
		if err != nil {
			return err
		}
		if !res.IsSuccess() {
			errFlag = true
			if jsonFlag {
				cui.WriteErrFrom(res.Body())
				return nil
			}
			ed, err2 := response.DecodeError(res.Body())
			if err != nil {
				return err2
			}
			for _, d := range ed.Errors {
				cui.OutputErrln(d.Message)
			}
			return nil
		}
		if jsonFlag {
			cui.WriteFrom(res.Body())
		}
		if opts.Cmd() == options.CmdProj {
			sd, err := response.DecodeSuccessProject(res.Body())
			if err != nil {
				return err
			}
			cui.Outputln(fmt.Sprintf("%d:%s, %s (Owner: %d:%s)", sd.ID, sd.Name, sd.Description, sd.Owner.ID, sd.Owner.Name))
		} else {
			sd, err := response.DecodeSuccessProjects(res.Body())
			if err != nil {
				return err
			}
			for _, d := range sd {
				cui.Outputln(fmt.Sprintf("%d:%s, %s (Owner: %d:%s)", d.ID, d.Name, d.Description, d.Owner.ID, d.Owner.Name))
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(projCmd)

	projCmd.Flags().IntP("projectid", "p", 0, "project_id parameter")
}
