package models

type MediaEntity struct {
	Entity
	UserID               string                 `json:"userId"`
	MimeType             string                 `json:"mimeType"`
	FileExtension        string                 `json:"fileExtension"`
	FileSize             int                    `json:"fileSize"`
	Title                string                 `json:"title"`
	MetaDataRaw          string                 `json:"metaDataRaw"`
	MediaTypeRaw         string                 `json:"mediaTypeRaw"`
	MetaData             MediaMetaData          `json:"metaData"`
	MediaType            MediaType              `json:"mediaType"`
	UploadedAt           string                 `json:"uploadedAt"`
	Alt                  string                 `json:"alt"`
	URL                  string                 `json:"url"`
	FileName             string                 `json:"fileName"`
	Translations         []*MediaTranslation    `json:"translations"`
	Categories           []*Category            `json:"categories"`
	ProductManufacturers []*ProductManufacturer `json:"productManufacturers"`
	ProductMedia         []*ProductMediaEntity  `json:"productMedia"`
	Thumbnails           []*MediaThumbnail      `json:"thumbnails"`
	MediaFolderID        string                 `json:"mediaFolderId"`
	MediaFolder          MediaFolder            `json:"mediaFolder"`
	HasFile              bool                   `json:"hasFile"`
	Private              bool                   `json:"private"`
	Tags                 []*Tag                 `json:"tags"`
	ThumbnailsRo         string                 `json:"thumbnailsRo"`
}

type MediaMetaData struct {
	Type   int     `json:"type"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type MediaType struct {
	Name  string   `json:"name"`
	Flags []string `json:"flags"`
}

type ImageType struct {
	MediaType
}
