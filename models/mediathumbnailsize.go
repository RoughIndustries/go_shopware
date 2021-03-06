package models

type MediaThumbnailSize struct {
	Entity
	Width                     int                         `json:"width"`
	Height                    int                         `json:"height"`
	MediaFolderConfigurations []*MediaFolderConfiguration `json:"mediaFolderConfigurations"`
}

// GetWidth returns the width of thumbnail
func (m *MediaThumbnailSize) GetWidth() int {
	return m.Width
}

// GetHeight returns the height of thumbnail
func (m *MediaThumbnailSize) GetHeight() int {
	return m.Height
}

// GetAPIAlias returns the ApiAlias of thumbnail
func (m *MediaThumbnailSize) GetAPIAlias() string {
	return "media_thumbnail_size"
}
