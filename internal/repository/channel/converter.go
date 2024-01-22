package channel

import (
	"github.com/api-sekejap/internal/entity"
	"github.com/api-sekejap/internal/repository/base"
)

func parserParams(params entity.Channel) Channels {
	return Channels{
		PackageID:   params.PackageID,
		Name:        params.Name,
		Link:        params.Link,
		AssetUrl:    params.AssetUrl,
		Description: params.AssetUrl,
		Meta:        base.Metadata(params.Meta),
		Extra:       base.ExtraAttribute(params.Extra),
	}
}
