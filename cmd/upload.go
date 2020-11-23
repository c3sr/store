package cmd

import (
	"io"
	"os"
	"time"

	"github.com/Unknwon/com"
	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
	"github.com/c3sr/archive"
	"github.com/c3sr/aws"
	"github.com/c3sr/store"
	"github.com/c3sr/store/s3"
	"github.com/c3sr/uuid"
	"github.com/spf13/cobra"
)

var (
	uploadKey        string
	baseURL          string
	accessKey        string
	secretKey        string
	expirationMonths int
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use: "upload",
	Aliases: []string{
		"put",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("expected a file or directory to upload")
		}
		fileName := args[0]
		if !com.IsFile(fileName) && !com.IsDir(fileName) {
			return errors.Errorf("file or directory %v was not found", fileName)
		}

		var err error
		var reader io.ReadCloser

		if com.IsDir(fileName) {
			reader, err = archive.Zip(fileName)
			if err != nil {
				return errors.Wrapf(err, "unable to archive %v", fileName)
			}
		} else {
			reader, err = os.Open(fileName)
			if err != nil {
				return errors.Wrapf(err, "unable to open file %v", fileName)
			}
		}

		defer reader.Close()

		if uploadKey == "" {
			uploadKey = uuid.New(fileName)
		}

		awsSession, err := aws.NewSession(aws.AccessKey(accessKey), aws.SecretKey(secretKey))
		if err != nil {
			return err
		}
		str, err := s3.New(
			store.BaseURL(baseURL),
			s3.Session(awsSession),
		)
		if err != nil {
			return errors.Wrapf(err, "unable to create an s3 connection")
		}

		uploadOpts := []store.UploadOption{}

		if expirationMonths != 0 {
			uploadOpts = append(
				uploadOpts,
				s3.Expiration(time.Now().AddDate(0, expirationMonths, 0)),
			)
		}

		key, err := str.UploadFrom(reader,
			uploadKey,
			uploadOpts...,
		)
		if err != nil {
			return err
		}

		pp.Println(key)

		return nil
	},
}

func init() {
	uploadCmd.PersistentFlags().StringVarP(&uploadKey, "key", "k", "", "upload key")
	uploadCmd.PersistentFlags().StringVarP(&baseURL, "baseurl", "b", "", "base url")
	uploadCmd.PersistentFlags().StringVarP(&accessKey, "accesskey", "a", "", "aws access key")
	uploadCmd.PersistentFlags().StringVarP(&secretKey, "secretkey", "s", "", "aws secret key")
	uploadCmd.PersistentFlags().IntVarP(&expirationMonths, "expiration", "e", 0, "expiration of the upload in months. leave empty for no expiration")
}
