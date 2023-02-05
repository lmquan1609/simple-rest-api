package uploadprovider

import (
	"context"
	"simple-rest-api/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
